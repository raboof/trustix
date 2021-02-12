# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from trustix_schema import sth_pb2 as schema_dot_sth__pb2
from trustix_schema import logleaf_pb2 as schema_dot_logleaf__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
    name="api/api.proto",
    package="trustix",
    syntax="proto2",
    serialized_options=b"Z\034github.com/tweag/trustix/api",
    create_key=_descriptor._internal_create_key,
    serialized_pb=b'\n\rapi/api.proto\x12\x07trustix\x1a\x10schema/sth.proto\x1a\x14schema/logleaf.proto"\x0c\n\nSTHRequest"F\n\x1dGetLogConsistencyProofRequest\x12\x11\n\tFirstSize\x18\x01 \x02(\x04\x12\x12\n\nSecondSize\x18\x02 \x02(\x04"\x1e\n\rProofResponse\x12\r\n\x05Proof\x18\x01 \x03(\x0c":\n\x17GetLogAuditProofRequest\x12\r\n\x05Index\x18\x01 \x02(\x04\x12\x10\n\x08TreeSize\x18\x02 \x02(\x04"5\n\x14GetLogEntriesRequest\x12\r\n\x05Start\x18\x01 \x02(\x04\x12\x0e\n\x06\x46inish\x18\x02 \x02(\x04"2\n\x12GetMapValueRequest\x12\x0b\n\x03Key\x18\x01 \x02(\x0c\x12\x0f\n\x07MapRoot\x18\x02 \x02(\x0c"s\n\x18SparseCompactMerkleProof\x12\x11\n\tSideNodes\x18\x01 \x03(\x0c\x12\x1d\n\x15NonMembershipLeafData\x18\x02 \x02(\x0c\x12\x0f\n\x07\x42itMask\x18\x03 \x02(\x0c\x12\x14\n\x0cNumSideNodes\x18\x04 \x02(\x04"S\n\x10MapValueResponse\x12\r\n\x05Value\x18\x01 \x02(\x0c\x12\x30\n\x05Proof\x18\x02 \x02(\x0b\x32!.trustix.SparseCompactMerkleProof".\n\x12LogEntriesResponse\x12\x18\n\x06Leaves\x18\x01 \x03(\x0b\x32\x08.LogLeaf"*\n\x0cKeyValuePair\x12\x0b\n\x03Key\x18\x01 \x02(\x0c\x12\r\n\x05Value\x18\x02 \x02(\x0c"5\n\rSubmitRequest\x12$\n\x05Items\x18\x01 \x03(\x0b\x32\x15.trustix.KeyValuePair"R\n\x0eSubmitResponse\x12.\n\x06status\x18\x01 \x02(\x0e\x32\x1e.trustix.SubmitResponse.Status"\x10\n\x06Status\x12\x06\n\x02OK\x10\x00"\x0e\n\x0c\x46lushRequest"\x0f\n\rFlushResponse"\x1e\n\x0cValueRequest\x12\x0e\n\x06\x44igest\x18\x01 \x02(\x0c"\x1e\n\rValueResponse\x12\r\n\x05Value\x18\x01 \x02(\x0c\x32\xaf\x06\n\rTrustixLogAPI\x12%\n\x06GetSTH\x12\x13.trustix.STHRequest\x1a\x04.STH"\x00\x12;\n\x06Submit\x12\x16.trustix.SubmitRequest\x1a\x17.trustix.SubmitResponse"\x00\x12\x38\n\x05\x46lush\x12\x15.trustix.FlushRequest\x1a\x16.trustix.FlushResponse"\x00\x12;\n\x08GetValue\x12\x15.trustix.ValueRequest\x1a\x16.trustix.ValueResponse"\x00\x12Z\n\x16GetLogConsistencyProof\x12&.trustix.GetLogConsistencyProofRequest\x1a\x16.trustix.ProofResponse"\x00\x12N\n\x10GetLogAuditProof\x12 .trustix.GetLogAuditProofRequest\x1a\x16.trustix.ProofResponse"\x00\x12M\n\rGetLogEntries\x12\x1d.trustix.GetLogEntriesRequest\x1a\x1b.trustix.LogEntriesResponse"\x00\x12G\n\x0bGetMapValue\x12\x1b.trustix.GetMapValueRequest\x1a\x19.trustix.MapValueResponse"\x00\x12\\\n\x18GetMHLogConsistencyProof\x12&.trustix.GetLogConsistencyProofRequest\x1a\x16.trustix.ProofResponse"\x00\x12P\n\x12GetMHLogAuditProof\x12 .trustix.GetLogAuditProofRequest\x1a\x16.trustix.ProofResponse"\x00\x12O\n\x0fGetMHLogEntries\x12\x1d.trustix.GetLogEntriesRequest\x1a\x1b.trustix.LogEntriesResponse"\x00\x42\x1eZ\x1cgithub.com/tweag/trustix/api',
    dependencies=[
        schema_dot_sth__pb2.DESCRIPTOR,
        schema_dot_logleaf__pb2.DESCRIPTOR,
    ],
)


_SUBMITRESPONSE_STATUS = _descriptor.EnumDescriptor(
    name="Status",
    full_name="trustix.SubmitResponse.Status",
    filename=None,
    file=DESCRIPTOR,
    create_key=_descriptor._internal_create_key,
    values=[
        _descriptor.EnumValueDescriptor(
            name="OK",
            index=0,
            number=0,
            serialized_options=None,
            type=None,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    containing_type=None,
    serialized_options=None,
    serialized_start=766,
    serialized_end=782,
)
_sym_db.RegisterEnumDescriptor(_SUBMITRESPONSE_STATUS)


_STHREQUEST = _descriptor.Descriptor(
    name="STHRequest",
    full_name="trustix.STHRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=66,
    serialized_end=78,
)


_GETLOGCONSISTENCYPROOFREQUEST = _descriptor.Descriptor(
    name="GetLogConsistencyProofRequest",
    full_name="trustix.GetLogConsistencyProofRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="FirstSize",
            full_name="trustix.GetLogConsistencyProofRequest.FirstSize",
            index=0,
            number=1,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="SecondSize",
            full_name="trustix.GetLogConsistencyProofRequest.SecondSize",
            index=1,
            number=2,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=80,
    serialized_end=150,
)


_PROOFRESPONSE = _descriptor.Descriptor(
    name="ProofResponse",
    full_name="trustix.ProofResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Proof",
            full_name="trustix.ProofResponse.Proof",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=152,
    serialized_end=182,
)


_GETLOGAUDITPROOFREQUEST = _descriptor.Descriptor(
    name="GetLogAuditProofRequest",
    full_name="trustix.GetLogAuditProofRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Index",
            full_name="trustix.GetLogAuditProofRequest.Index",
            index=0,
            number=1,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="TreeSize",
            full_name="trustix.GetLogAuditProofRequest.TreeSize",
            index=1,
            number=2,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=184,
    serialized_end=242,
)


_GETLOGENTRIESREQUEST = _descriptor.Descriptor(
    name="GetLogEntriesRequest",
    full_name="trustix.GetLogEntriesRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Start",
            full_name="trustix.GetLogEntriesRequest.Start",
            index=0,
            number=1,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="Finish",
            full_name="trustix.GetLogEntriesRequest.Finish",
            index=1,
            number=2,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=244,
    serialized_end=297,
)


_GETMAPVALUEREQUEST = _descriptor.Descriptor(
    name="GetMapValueRequest",
    full_name="trustix.GetMapValueRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Key",
            full_name="trustix.GetMapValueRequest.Key",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="MapRoot",
            full_name="trustix.GetMapValueRequest.MapRoot",
            index=1,
            number=2,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=299,
    serialized_end=349,
)


_SPARSECOMPACTMERKLEPROOF = _descriptor.Descriptor(
    name="SparseCompactMerkleProof",
    full_name="trustix.SparseCompactMerkleProof",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="SideNodes",
            full_name="trustix.SparseCompactMerkleProof.SideNodes",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="NonMembershipLeafData",
            full_name="trustix.SparseCompactMerkleProof.NonMembershipLeafData",
            index=1,
            number=2,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="BitMask",
            full_name="trustix.SparseCompactMerkleProof.BitMask",
            index=2,
            number=3,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="NumSideNodes",
            full_name="trustix.SparseCompactMerkleProof.NumSideNodes",
            index=3,
            number=4,
            type=4,
            cpp_type=4,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=351,
    serialized_end=466,
)


_MAPVALUERESPONSE = _descriptor.Descriptor(
    name="MapValueResponse",
    full_name="trustix.MapValueResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Value",
            full_name="trustix.MapValueResponse.Value",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="Proof",
            full_name="trustix.MapValueResponse.Proof",
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=2,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=468,
    serialized_end=551,
)


_LOGENTRIESRESPONSE = _descriptor.Descriptor(
    name="LogEntriesResponse",
    full_name="trustix.LogEntriesResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Leaves",
            full_name="trustix.LogEntriesResponse.Leaves",
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=553,
    serialized_end=599,
)


_KEYVALUEPAIR = _descriptor.Descriptor(
    name="KeyValuePair",
    full_name="trustix.KeyValuePair",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Key",
            full_name="trustix.KeyValuePair.Key",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="Value",
            full_name="trustix.KeyValuePair.Value",
            index=1,
            number=2,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=601,
    serialized_end=643,
)


_SUBMITREQUEST = _descriptor.Descriptor(
    name="SubmitRequest",
    full_name="trustix.SubmitRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Items",
            full_name="trustix.SubmitRequest.Items",
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=645,
    serialized_end=698,
)


_SUBMITRESPONSE = _descriptor.Descriptor(
    name="SubmitResponse",
    full_name="trustix.SubmitResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="status",
            full_name="trustix.SubmitResponse.status",
            index=0,
            number=1,
            type=14,
            cpp_type=8,
            label=2,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[
        _SUBMITRESPONSE_STATUS,
    ],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=700,
    serialized_end=782,
)


_FLUSHREQUEST = _descriptor.Descriptor(
    name="FlushRequest",
    full_name="trustix.FlushRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=784,
    serialized_end=798,
)


_FLUSHRESPONSE = _descriptor.Descriptor(
    name="FlushResponse",
    full_name="trustix.FlushResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=800,
    serialized_end=815,
)


_VALUEREQUEST = _descriptor.Descriptor(
    name="ValueRequest",
    full_name="trustix.ValueRequest",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Digest",
            full_name="trustix.ValueRequest.Digest",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=817,
    serialized_end=847,
)


_VALUERESPONSE = _descriptor.Descriptor(
    name="ValueResponse",
    full_name="trustix.ValueResponse",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="Value",
            full_name="trustix.ValueResponse.Value",
            index=0,
            number=1,
            type=12,
            cpp_type=9,
            label=2,
            has_default_value=False,
            default_value=b"",
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto2",
    extension_ranges=[],
    oneofs=[],
    serialized_start=849,
    serialized_end=879,
)

_MAPVALUERESPONSE.fields_by_name["Proof"].message_type = _SPARSECOMPACTMERKLEPROOF
_LOGENTRIESRESPONSE.fields_by_name[
    "Leaves"
].message_type = schema_dot_logleaf__pb2._LOGLEAF
_SUBMITREQUEST.fields_by_name["Items"].message_type = _KEYVALUEPAIR
_SUBMITRESPONSE.fields_by_name["status"].enum_type = _SUBMITRESPONSE_STATUS
_SUBMITRESPONSE_STATUS.containing_type = _SUBMITRESPONSE
DESCRIPTOR.message_types_by_name["STHRequest"] = _STHREQUEST
DESCRIPTOR.message_types_by_name[
    "GetLogConsistencyProofRequest"
] = _GETLOGCONSISTENCYPROOFREQUEST
DESCRIPTOR.message_types_by_name["ProofResponse"] = _PROOFRESPONSE
DESCRIPTOR.message_types_by_name["GetLogAuditProofRequest"] = _GETLOGAUDITPROOFREQUEST
DESCRIPTOR.message_types_by_name["GetLogEntriesRequest"] = _GETLOGENTRIESREQUEST
DESCRIPTOR.message_types_by_name["GetMapValueRequest"] = _GETMAPVALUEREQUEST
DESCRIPTOR.message_types_by_name["SparseCompactMerkleProof"] = _SPARSECOMPACTMERKLEPROOF
DESCRIPTOR.message_types_by_name["MapValueResponse"] = _MAPVALUERESPONSE
DESCRIPTOR.message_types_by_name["LogEntriesResponse"] = _LOGENTRIESRESPONSE
DESCRIPTOR.message_types_by_name["KeyValuePair"] = _KEYVALUEPAIR
DESCRIPTOR.message_types_by_name["SubmitRequest"] = _SUBMITREQUEST
DESCRIPTOR.message_types_by_name["SubmitResponse"] = _SUBMITRESPONSE
DESCRIPTOR.message_types_by_name["FlushRequest"] = _FLUSHREQUEST
DESCRIPTOR.message_types_by_name["FlushResponse"] = _FLUSHRESPONSE
DESCRIPTOR.message_types_by_name["ValueRequest"] = _VALUEREQUEST
DESCRIPTOR.message_types_by_name["ValueResponse"] = _VALUERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

STHRequest = _reflection.GeneratedProtocolMessageType(
    "STHRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _STHREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.STHRequest)
    },
)
_sym_db.RegisterMessage(STHRequest)

GetLogConsistencyProofRequest = _reflection.GeneratedProtocolMessageType(
    "GetLogConsistencyProofRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETLOGCONSISTENCYPROOFREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.GetLogConsistencyProofRequest)
    },
)
_sym_db.RegisterMessage(GetLogConsistencyProofRequest)

ProofResponse = _reflection.GeneratedProtocolMessageType(
    "ProofResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _PROOFRESPONSE,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.ProofResponse)
    },
)
_sym_db.RegisterMessage(ProofResponse)

GetLogAuditProofRequest = _reflection.GeneratedProtocolMessageType(
    "GetLogAuditProofRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETLOGAUDITPROOFREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.GetLogAuditProofRequest)
    },
)
_sym_db.RegisterMessage(GetLogAuditProofRequest)

GetLogEntriesRequest = _reflection.GeneratedProtocolMessageType(
    "GetLogEntriesRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETLOGENTRIESREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.GetLogEntriesRequest)
    },
)
_sym_db.RegisterMessage(GetLogEntriesRequest)

GetMapValueRequest = _reflection.GeneratedProtocolMessageType(
    "GetMapValueRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETMAPVALUEREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.GetMapValueRequest)
    },
)
_sym_db.RegisterMessage(GetMapValueRequest)

SparseCompactMerkleProof = _reflection.GeneratedProtocolMessageType(
    "SparseCompactMerkleProof",
    (_message.Message,),
    {
        "DESCRIPTOR": _SPARSECOMPACTMERKLEPROOF,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.SparseCompactMerkleProof)
    },
)
_sym_db.RegisterMessage(SparseCompactMerkleProof)

MapValueResponse = _reflection.GeneratedProtocolMessageType(
    "MapValueResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _MAPVALUERESPONSE,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.MapValueResponse)
    },
)
_sym_db.RegisterMessage(MapValueResponse)

LogEntriesResponse = _reflection.GeneratedProtocolMessageType(
    "LogEntriesResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _LOGENTRIESRESPONSE,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.LogEntriesResponse)
    },
)
_sym_db.RegisterMessage(LogEntriesResponse)

KeyValuePair = _reflection.GeneratedProtocolMessageType(
    "KeyValuePair",
    (_message.Message,),
    {
        "DESCRIPTOR": _KEYVALUEPAIR,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.KeyValuePair)
    },
)
_sym_db.RegisterMessage(KeyValuePair)

SubmitRequest = _reflection.GeneratedProtocolMessageType(
    "SubmitRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _SUBMITREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.SubmitRequest)
    },
)
_sym_db.RegisterMessage(SubmitRequest)

SubmitResponse = _reflection.GeneratedProtocolMessageType(
    "SubmitResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _SUBMITRESPONSE,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.SubmitResponse)
    },
)
_sym_db.RegisterMessage(SubmitResponse)

FlushRequest = _reflection.GeneratedProtocolMessageType(
    "FlushRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _FLUSHREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.FlushRequest)
    },
)
_sym_db.RegisterMessage(FlushRequest)

FlushResponse = _reflection.GeneratedProtocolMessageType(
    "FlushResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _FLUSHRESPONSE,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.FlushResponse)
    },
)
_sym_db.RegisterMessage(FlushResponse)

ValueRequest = _reflection.GeneratedProtocolMessageType(
    "ValueRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _VALUEREQUEST,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.ValueRequest)
    },
)
_sym_db.RegisterMessage(ValueRequest)

ValueResponse = _reflection.GeneratedProtocolMessageType(
    "ValueResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _VALUERESPONSE,
        "__module__": "api.api_pb2"
        # @@protoc_insertion_point(class_scope:trustix.ValueResponse)
    },
)
_sym_db.RegisterMessage(ValueResponse)


DESCRIPTOR._options = None

_TRUSTIXLOGAPI = _descriptor.ServiceDescriptor(
    name="TrustixLogAPI",
    full_name="trustix.TrustixLogAPI",
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
    serialized_start=882,
    serialized_end=1697,
    methods=[
        _descriptor.MethodDescriptor(
            name="GetSTH",
            full_name="trustix.TrustixLogAPI.GetSTH",
            index=0,
            containing_service=None,
            input_type=_STHREQUEST,
            output_type=schema_dot_sth__pb2._STH,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="Submit",
            full_name="trustix.TrustixLogAPI.Submit",
            index=1,
            containing_service=None,
            input_type=_SUBMITREQUEST,
            output_type=_SUBMITRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="Flush",
            full_name="trustix.TrustixLogAPI.Flush",
            index=2,
            containing_service=None,
            input_type=_FLUSHREQUEST,
            output_type=_FLUSHRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetValue",
            full_name="trustix.TrustixLogAPI.GetValue",
            index=3,
            containing_service=None,
            input_type=_VALUEREQUEST,
            output_type=_VALUERESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetLogConsistencyProof",
            full_name="trustix.TrustixLogAPI.GetLogConsistencyProof",
            index=4,
            containing_service=None,
            input_type=_GETLOGCONSISTENCYPROOFREQUEST,
            output_type=_PROOFRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetLogAuditProof",
            full_name="trustix.TrustixLogAPI.GetLogAuditProof",
            index=5,
            containing_service=None,
            input_type=_GETLOGAUDITPROOFREQUEST,
            output_type=_PROOFRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetLogEntries",
            full_name="trustix.TrustixLogAPI.GetLogEntries",
            index=6,
            containing_service=None,
            input_type=_GETLOGENTRIESREQUEST,
            output_type=_LOGENTRIESRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetMapValue",
            full_name="trustix.TrustixLogAPI.GetMapValue",
            index=7,
            containing_service=None,
            input_type=_GETMAPVALUEREQUEST,
            output_type=_MAPVALUERESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetMHLogConsistencyProof",
            full_name="trustix.TrustixLogAPI.GetMHLogConsistencyProof",
            index=8,
            containing_service=None,
            input_type=_GETLOGCONSISTENCYPROOFREQUEST,
            output_type=_PROOFRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetMHLogAuditProof",
            full_name="trustix.TrustixLogAPI.GetMHLogAuditProof",
            index=9,
            containing_service=None,
            input_type=_GETLOGAUDITPROOFREQUEST,
            output_type=_PROOFRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.MethodDescriptor(
            name="GetMHLogEntries",
            full_name="trustix.TrustixLogAPI.GetMHLogEntries",
            index=10,
            containing_service=None,
            input_type=_GETLOGENTRIESREQUEST,
            output_type=_LOGENTRIESRESPONSE,
            serialized_options=None,
            create_key=_descriptor._internal_create_key,
        ),
    ],
)
_sym_db.RegisterServiceDescriptor(_TRUSTIXLOGAPI)

DESCRIPTOR.services_by_name["TrustixLogAPI"] = _TRUSTIXLOGAPI

# @@protoc_insertion_point(module_scope)
