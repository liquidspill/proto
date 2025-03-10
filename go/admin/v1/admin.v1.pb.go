// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: admin.v1.proto

package adminv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DatabaseGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      uint64                 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	EnvironmentId uint64                 `protobuf:"varint,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Table         []*Variant             `protobuf:"bytes,3,rep,name=table,proto3" json:"table,omitempty"`
	Keys          []*Variant             `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DatabaseGetRequest) Reset() {
	*x = DatabaseGetRequest{}
	mi := &file_admin_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatabaseGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseGetRequest) ProtoMessage() {}

func (x *DatabaseGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseGetRequest.ProtoReflect.Descriptor instead.
func (*DatabaseGetRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseGetRequest) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *DatabaseGetRequest) GetEnvironmentId() uint64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *DatabaseGetRequest) GetTable() []*Variant {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *DatabaseGetRequest) GetKeys() []*Variant {
	if x != nil {
		return x.Keys
	}
	return nil
}

type DatabaseGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      uint64                 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	EnvironmentId uint64                 `protobuf:"varint,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Values        []*Variant             `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DatabaseGetResponse) Reset() {
	*x = DatabaseGetResponse{}
	mi := &file_admin_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatabaseGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseGetResponse) ProtoMessage() {}

func (x *DatabaseGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseGetResponse.ProtoReflect.Descriptor instead.
func (*DatabaseGetResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseGetResponse) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *DatabaseGetResponse) GetEnvironmentId() uint64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *DatabaseGetResponse) GetValues() []*Variant {
	if x != nil {
		return x.Values
	}
	return nil
}

type DatabaseScanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      uint64                 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	EnvironmentId uint64                 `protobuf:"varint,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Table         []*Variant             `protobuf:"bytes,3,rep,name=table,proto3" json:"table,omitempty"`
	Start         *Variant               `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End           *Variant               `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DatabaseScanRequest) Reset() {
	*x = DatabaseScanRequest{}
	mi := &file_admin_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatabaseScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseScanRequest) ProtoMessage() {}

func (x *DatabaseScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseScanRequest.ProtoReflect.Descriptor instead.
func (*DatabaseScanRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DatabaseScanRequest) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *DatabaseScanRequest) GetEnvironmentId() uint64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *DatabaseScanRequest) GetTable() []*Variant {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *DatabaseScanRequest) GetStart() *Variant {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *DatabaseScanRequest) GetEnd() *Variant {
	if x != nil {
		return x.End
	}
	return nil
}

type DatabaseScanResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      uint64                 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	EnvironmentId uint64                 `protobuf:"varint,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Values        []*Variant             `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DatabaseScanResponse) Reset() {
	*x = DatabaseScanResponse{}
	mi := &file_admin_v1_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatabaseScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseScanResponse) ProtoMessage() {}

func (x *DatabaseScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseScanResponse.ProtoReflect.Descriptor instead.
func (*DatabaseScanResponse) Descriptor() ([]byte, []int) {
	return file_admin_v1_proto_rawDescGZIP(), []int{3}
}

func (x *DatabaseScanResponse) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *DatabaseScanResponse) GetEnvironmentId() uint64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *DatabaseScanResponse) GetValues() []*Variant {
	if x != nil {
		return x.Values
	}
	return nil
}

type Variant struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Variant:
	//
	//	*Variant_Int32Variant
	//	*Variant_Int64Variant
	//	*Variant_Uint32Variant
	//	*Variant_Uint64Variant
	//	*Variant_BoolVariant
	//	*Variant_StringVariant
	//	*Variant_BytesVariant
	//	*Variant_TimestampVariant
	Variant       isVariant_Variant `protobuf_oneof:"variant"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Variant) Reset() {
	*x = Variant{}
	mi := &file_admin_v1_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_admin_v1_proto_rawDescGZIP(), []int{4}
}

func (x *Variant) GetVariant() isVariant_Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *Variant) GetInt32Variant() int32 {
	if x != nil {
		if x, ok := x.Variant.(*Variant_Int32Variant); ok {
			return x.Int32Variant
		}
	}
	return 0
}

func (x *Variant) GetInt64Variant() int64 {
	if x != nil {
		if x, ok := x.Variant.(*Variant_Int64Variant); ok {
			return x.Int64Variant
		}
	}
	return 0
}

func (x *Variant) GetUint32Variant() uint32 {
	if x != nil {
		if x, ok := x.Variant.(*Variant_Uint32Variant); ok {
			return x.Uint32Variant
		}
	}
	return 0
}

func (x *Variant) GetUint64Variant() uint64 {
	if x != nil {
		if x, ok := x.Variant.(*Variant_Uint64Variant); ok {
			return x.Uint64Variant
		}
	}
	return 0
}

func (x *Variant) GetBoolVariant() bool {
	if x != nil {
		if x, ok := x.Variant.(*Variant_BoolVariant); ok {
			return x.BoolVariant
		}
	}
	return false
}

func (x *Variant) GetStringVariant() string {
	if x != nil {
		if x, ok := x.Variant.(*Variant_StringVariant); ok {
			return x.StringVariant
		}
	}
	return ""
}

func (x *Variant) GetBytesVariant() []byte {
	if x != nil {
		if x, ok := x.Variant.(*Variant_BytesVariant); ok {
			return x.BytesVariant
		}
	}
	return nil
}

func (x *Variant) GetTimestampVariant() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.Variant.(*Variant_TimestampVariant); ok {
			return x.TimestampVariant
		}
	}
	return nil
}

type isVariant_Variant interface {
	isVariant_Variant()
}

type Variant_Int32Variant struct {
	Int32Variant int32 `protobuf:"varint,1,opt,name=int32_variant,json=int32Variant,proto3,oneof"`
}

type Variant_Int64Variant struct {
	Int64Variant int64 `protobuf:"varint,2,opt,name=int64_variant,json=int64Variant,proto3,oneof"`
}

type Variant_Uint32Variant struct {
	Uint32Variant uint32 `protobuf:"varint,3,opt,name=uint32_variant,json=uint32Variant,proto3,oneof"`
}

type Variant_Uint64Variant struct {
	Uint64Variant uint64 `protobuf:"varint,4,opt,name=uint64_variant,json=uint64Variant,proto3,oneof"`
}

type Variant_BoolVariant struct {
	BoolVariant bool `protobuf:"varint,5,opt,name=bool_variant,json=boolVariant,proto3,oneof"`
}

type Variant_StringVariant struct {
	StringVariant string `protobuf:"bytes,6,opt,name=string_variant,json=stringVariant,proto3,oneof"`
}

type Variant_BytesVariant struct {
	BytesVariant []byte `protobuf:"bytes,7,opt,name=bytes_variant,json=bytesVariant,proto3,oneof"`
}

type Variant_TimestampVariant struct {
	TimestampVariant *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp_variant,json=timestampVariant,proto3,oneof"`
}

func (*Variant_Int32Variant) isVariant_Variant() {}

func (*Variant_Int64Variant) isVariant_Variant() {}

func (*Variant_Uint32Variant) isVariant_Variant() {}

func (*Variant_Uint64Variant) isVariant_Variant() {}

func (*Variant_BoolVariant) isVariant_Variant() {}

func (*Variant_StringVariant) isVariant_Variant() {}

func (*Variant_BytesVariant) isVariant_Variant() {}

func (*Variant_TimestampVariant) isVariant_Variant() {}

var File_admin_v1_proto protoreflect.FileDescriptor

var file_admin_v1_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xd0, 0x01,
	0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x22, 0x85, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xf4, 0x02, 0x0a, 0x07, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0d, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x00, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x32,
	0xad, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4c, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x65, 0x74, 0x12,
	0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x1d,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_admin_v1_proto_rawDescOnce sync.Once
	file_admin_v1_proto_rawDescData []byte
)

func file_admin_v1_proto_rawDescGZIP() []byte {
	file_admin_v1_proto_rawDescOnce.Do(func() {
		file_admin_v1_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_admin_v1_proto_rawDesc), len(file_admin_v1_proto_rawDesc)))
	})
	return file_admin_v1_proto_rawDescData
}

var file_admin_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_admin_v1_proto_goTypes = []any{
	(*DatabaseGetRequest)(nil),    // 0: admin.v1.DatabaseGetRequest
	(*DatabaseGetResponse)(nil),   // 1: admin.v1.DatabaseGetResponse
	(*DatabaseScanRequest)(nil),   // 2: admin.v1.DatabaseScanRequest
	(*DatabaseScanResponse)(nil),  // 3: admin.v1.DatabaseScanResponse
	(*Variant)(nil),               // 4: admin.v1.Variant
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_admin_v1_proto_depIdxs = []int32{
	4,  // 0: admin.v1.DatabaseGetRequest.table:type_name -> admin.v1.Variant
	4,  // 1: admin.v1.DatabaseGetRequest.keys:type_name -> admin.v1.Variant
	4,  // 2: admin.v1.DatabaseGetResponse.values:type_name -> admin.v1.Variant
	4,  // 3: admin.v1.DatabaseScanRequest.table:type_name -> admin.v1.Variant
	4,  // 4: admin.v1.DatabaseScanRequest.start:type_name -> admin.v1.Variant
	4,  // 5: admin.v1.DatabaseScanRequest.end:type_name -> admin.v1.Variant
	4,  // 6: admin.v1.DatabaseScanResponse.values:type_name -> admin.v1.Variant
	5,  // 7: admin.v1.Variant.timestamp_variant:type_name -> google.protobuf.Timestamp
	0,  // 8: admin.v1.AdminService.DatabaseGet:input_type -> admin.v1.DatabaseGetRequest
	2,  // 9: admin.v1.AdminService.DatabaseScan:input_type -> admin.v1.DatabaseScanRequest
	1,  // 10: admin.v1.AdminService.DatabaseGet:output_type -> admin.v1.DatabaseGetResponse
	3,  // 11: admin.v1.AdminService.DatabaseScan:output_type -> admin.v1.DatabaseScanResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_admin_v1_proto_init() }
func file_admin_v1_proto_init() {
	if File_admin_v1_proto != nil {
		return
	}
	file_admin_v1_proto_msgTypes[4].OneofWrappers = []any{
		(*Variant_Int32Variant)(nil),
		(*Variant_Int64Variant)(nil),
		(*Variant_Uint32Variant)(nil),
		(*Variant_Uint64Variant)(nil),
		(*Variant_BoolVariant)(nil),
		(*Variant_StringVariant)(nil),
		(*Variant_BytesVariant)(nil),
		(*Variant_TimestampVariant)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_v1_proto_rawDesc), len(file_admin_v1_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_proto_goTypes,
		DependencyIndexes: file_admin_v1_proto_depIdxs,
		MessageInfos:      file_admin_v1_proto_msgTypes,
	}.Build()
	File_admin_v1_proto = out.File
	file_admin_v1_proto_goTypes = nil
	file_admin_v1_proto_depIdxs = nil
}
