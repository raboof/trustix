package core

import (
	"crypto/sha256"
	"fmt"
	"github.com/lazyledger/smt"
	"github.com/tweag/trustix/config"
	"github.com/tweag/trustix/correlator"
	vlog "github.com/tweag/trustix/log"
	"github.com/tweag/trustix/signer"
	"github.com/tweag/trustix/sth"
	"github.com/tweag/trustix/storage"
	"github.com/tweag/trustix/transport"
	"hash"
	"time"
)

type FlagConfig struct {
	StateDirectory string
}

type TrustixCore struct {
	store      storage.TrustixStorage
	hasher     hash.Hash
	signer     signer.TrustixSigner
	correlator correlator.LogCorrelator

	mapRoot  []byte
	logRoot  []byte
	treeSize int
}

func (s *TrustixCore) Query(key []byte) ([]byte, error) {
	var buf []byte

	err := s.store.View(func(txn storage.Transaction) error {
		hasher := sha256.New()
		tree := smt.ImportSparseMerkleTree(newMapStore(txn), hasher, s.mapRoot)

		// TODO: Log verification (but optional?)

		v, err := tree.Get(key)
		if err != nil {
			return err
		}
		buf = v

		// Check inclusion proof
		proof, err := tree.Prove(key)
		if err != nil {
			return err
		}

		if !smt.VerifyProof(proof, tree.Root(), key, v, s.hasher) {
			return fmt.Errorf("Proof verification failed")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TrustixCore) Get(bucket []byte, key []byte) ([]byte, error) {
	var buf []byte

	err := s.store.View(func(txn storage.Transaction) error {
		v, err := txn.Get(bucket, key)
		if err != nil {
			return err
		}
		buf = v

		return nil
	})
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TrustixCore) Submit(key []byte, value []byte) error {
	return s.store.Update(func(txn storage.Transaction) error {

		// The append-only log
		vLog, err := vlog.NewVerifiableLog(txn, s.treeSize)
		if err != nil {
			return err
		}

		// The sparse merkle tree
		smTree := smt.ImportSparseMerkleTree(newMapStore(txn), s.hasher, s.mapRoot)

		// Append value to both verifiable log & sparse indexed tree
		vLog.Append(value)
		smTree.Update(key, value)

		sth, err := sth.SignHead(smTree, vLog, s.signer)
		if err != nil {
			return err
		}

		smhBytes, err := sth.Marshal()
		if err != nil {
			return err
		}

		mapRoot, err := sth.UnmarshalSMHRoot()
		if err != nil {
			return err
		}

		logRoot, err := sth.UnmarshalSTHRoot()
		if err != nil {
			return err
		}

		err = txn.Set([]byte("META"), []byte("HEAD"), smhBytes)
		if err != nil {
			return err
		}

		s.mapRoot = mapRoot
		s.logRoot = logRoot
		s.treeSize = vLog.Size()

		return nil
	})
}

func (s *TrustixCore) updateRoot() error {
	return s.store.View(func(txn storage.Transaction) error {

		oldHead, err := txn.Get([]byte("META"), []byte("HEAD"))
		if err != nil {
			return err
		} else {
			oldSMH, err := sth.NewSMHFromJSON(oldHead)
			if err != nil {
				return err
			}
			sthRootBytes, err := oldSMH.UnmarshalSTHRoot()
			if err != nil {
				return err
			}

			// Verify signed map head
			smhRootBytes, err := oldSMH.UnmarshalSMHRoot()
			if err != nil {
				return err
			}

			err = oldSMH.Verify(s.signer)
			if err != nil {
				return err
			}

			s.mapRoot = smhRootBytes
			s.logRoot = sthRootBytes
			s.treeSize = oldSMH.LogSth.Size

		}

		return nil
	})
}

func CoreFromConfig(conf *config.LogConfig, flags *FlagConfig) (*TrustixCore, error) {

	hasher := sha256.New()

	sig, err := signer.FromConfig(conf.Signer)
	if err != nil {
		return nil, err
	}

	if conf.Mode == "trustix-log" {
		if !sig.CanSign() {
			return nil, fmt.Errorf("Cannot sign using the current configuration, aborting.")
		}
	}

	var store storage.TrustixStorage
	switch conf.Mode {
	case "trustix-log":
		store, err = storage.FromConfig(conf.Name, flags.StateDirectory, conf.Storage)
		if err != nil {
			return nil, err
		}
	case "trustix-follower":
		store, err = transport.NewGRPCTransport(conf.Transport.GRPC)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Mode '%s' unhandled", conf.Mode)
	}

	corr, err := correlator.NewMinimumPercentCorrelator(100)
	if err != nil {
		return nil, err
	}

	core := &TrustixCore{
		store:      store,
		hasher:     hasher,
		signer:     sig,
		correlator: corr,
		// TODO: Log root
	}

	err = store.View(func(txn storage.Transaction) error {
		_, err := txn.Get([]byte("META"), []byte("HEAD"))
		if err != nil {

			// No STH yet, new tree
			// TODO: Create a completely separate command for new tree, no magic should happen at startup
			if err == storage.ObjectNotFoundError {
				tree := smt.NewSparseMerkleTree(newMapStore(txn), hasher)
				core.mapRoot = tree.Root()
				return nil
			}
			return err
		} else {
			// TODO: Implement local cache and set to old values so we can verify consistency between last known good HEAD
			// and the newest HEAD
			return core.updateRoot()
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	switch conf.Mode {
	case "trustix-follower":
		go func() {
			for {
				// TODO: This is just an arbitrary interval for testing, not particularly intelligent
				time.Sleep(10 * time.Second)
				err := core.updateRoot()
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	default:
	}

	return core, nil
}
