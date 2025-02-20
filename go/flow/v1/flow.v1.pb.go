// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: flow.v1.proto

package flowv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Flow_FlowType int32

const (
	Flow_FLOWUNKNOWN Flow_FlowType = 0
	Flow_SFLOW_5     Flow_FlowType = 1
	Flow_NETFLOW_V5  Flow_FlowType = 2
	Flow_NETFLOW_V9  Flow_FlowType = 3
	Flow_IPFIX       Flow_FlowType = 4
)

// Enum value maps for Flow_FlowType.
var (
	Flow_FlowType_name = map[int32]string{
		0: "FLOWUNKNOWN",
		1: "SFLOW_5",
		2: "NETFLOW_V5",
		3: "NETFLOW_V9",
		4: "IPFIX",
	}
	Flow_FlowType_value = map[string]int32{
		"FLOWUNKNOWN": 0,
		"SFLOW_5":     1,
		"NETFLOW_V5":  2,
		"NETFLOW_V9":  3,
		"IPFIX":       4,
	}
)

func (x Flow_FlowType) Enum() *Flow_FlowType {
	p := new(Flow_FlowType)
	*p = x
	return p
}

func (x Flow_FlowType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Flow_FlowType) Descriptor() protoreflect.EnumDescriptor {
	return file_flow_v1_proto_enumTypes[0].Descriptor()
}

func (Flow_FlowType) Type() protoreflect.EnumType {
	return &file_flow_v1_proto_enumTypes[0]
}

func (x Flow_FlowType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Flow_FlowType.Descriptor instead.
func (Flow_FlowType) EnumDescriptor() ([]byte, []int) {
	return file_flow_v1_proto_rawDescGZIP(), []int{0, 0}
}

type Flow struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	FluidVersion string                 `protobuf:"bytes,1,opt,name=fluid_version,json=fluidVersion,proto3" json:"fluid_version,omitempty"`
	Type         Flow_FlowType          `protobuf:"varint,2,opt,name=type,proto3,enum=flow.v1.Flow_FlowType" json:"type,omitempty"`
	TimeReceived int64                  `protobuf:"varint,3,opt,name=time_received,json=timeReceived,proto3" json:"time_received,omitempty"`
	Sequence     uint32                 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	SamplingRate uint32                 `protobuf:"varint,5,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`
	Direction    uint32                 `protobuf:"varint,6,opt,name=direction,proto3" json:"direction,omitempty"`
	// Sampler information
	SamplerAddress []byte `protobuf:"bytes,7,opt,name=sampler_address,json=samplerAddress,proto3" json:"sampler_address,omitempty"` // Fixed 16 byte array
	// Flow timing
	FlowStartTime int64 `protobuf:"varint,8,opt,name=flow_start_time,json=flowStartTime,proto3" json:"flow_start_time,omitempty"`
	FlowEndTime   int64 `protobuf:"varint,9,opt,name=flow_end_time,json=flowEndTime,proto3" json:"flow_end_time,omitempty"`
	// Size metrics
	Bytes   uint32 `protobuf:"varint,10,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Packets uint32 `protobuf:"varint,11,opt,name=packets,proto3" json:"packets,omitempty"`
	// Network addresses
	SrcAddr   []byte `protobuf:"bytes,12,opt,name=src_addr,json=srcAddr,proto3" json:"src_addr,omitempty"` // Fixed 16 byte array
	DstAddr   []byte `protobuf:"bytes,13,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"` // Fixed 16 byte array
	Ethertype string `protobuf:"bytes,14,opt,name=ethertype,proto3" json:"ethertype,omitempty"`
	Protocol  uint32 `protobuf:"varint,15,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Ports
	SrcPort uint32 `protobuf:"varint,16,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort uint32 `protobuf:"varint,17,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	// Interfaces
	InInterface  uint32 `protobuf:"varint,18,opt,name=in_interface,json=inInterface,proto3" json:"in_interface,omitempty"`
	OutInterface uint32 `protobuf:"varint,19,opt,name=out_interface,json=outInterface,proto3" json:"out_interface,omitempty"`
	// Ethernet information
	SrcMac uint64 `protobuf:"varint,20,opt,name=src_mac,json=srcMac,proto3" json:"src_mac,omitempty"`
	DstMac uint64 `protobuf:"varint,21,opt,name=dst_mac,json=dstMac,proto3" json:"dst_mac,omitempty"`
	// VLAN information
	SrcVlan uint32 `protobuf:"varint,22,opt,name=src_vlan,json=srcVlan,proto3" json:"src_vlan,omitempty"`
	DstVlan uint32 `protobuf:"varint,23,opt,name=dst_vlan,json=dstVlan,proto3" json:"dst_vlan,omitempty"`
	VlanId  uint32 `protobuf:"varint,24,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	// VRF information
	IngressVrfId uint32 `protobuf:"varint,25,opt,name=ingress_vrf_id,json=ingressVrfId,proto3" json:"ingress_vrf_id,omitempty"`
	EgressVrfId  uint32 `protobuf:"varint,26,opt,name=egress_vrf_id,json=egressVrfId,proto3" json:"egress_vrf_id,omitempty"`
	// IP and TCP flags
	IpTos            uint32 `protobuf:"varint,27,opt,name=ip_tos,json=ipTos,proto3" json:"ip_tos,omitempty"`
	IpTtl            uint32 `protobuf:"varint,28,opt,name=ip_ttl,json=ipTtl,proto3" json:"ip_ttl,omitempty"`
	TcpFlags         uint32 `protobuf:"varint,29,opt,name=tcp_flags,json=tcpFlags,proto3" json:"tcp_flags,omitempty"`
	IcmpType         uint32 `protobuf:"varint,30,opt,name=icmp_type,json=icmpType,proto3" json:"icmp_type,omitempty"`
	IcmpCode         uint32 `protobuf:"varint,31,opt,name=icmp_code,json=icmpCode,proto3" json:"icmp_code,omitempty"`
	Ipv6FlowLabel    uint32 `protobuf:"varint,32,opt,name=ipv6_flow_label,json=ipv6FlowLabel,proto3" json:"ipv6_flow_label,omitempty"`
	ForwardingStatus uint32 `protobuf:"varint,33,opt,name=forwarding_status,json=forwardingStatus,proto3" json:"forwarding_status,omitempty"`
	// Fragment information
	FragmentId      uint32 `protobuf:"varint,34,opt,name=fragment_id,json=fragmentId,proto3" json:"fragment_id,omitempty"`
	FragmentOffset  uint32 `protobuf:"varint,35,opt,name=fragment_offset,json=fragmentOffset,proto3" json:"fragment_offset,omitempty"`
	BiFlowDirection uint32 `protobuf:"varint,36,opt,name=bi_flow_direction,json=biFlowDirection,proto3" json:"bi_flow_direction,omitempty"`
	// AS information
	SrcAs     string `protobuf:"bytes,37,opt,name=src_as,json=srcAs,proto3" json:"src_as,omitempty"`
	SrcAsn    uint32 `protobuf:"varint,38,opt,name=src_asn,json=srcAsn,proto3" json:"src_asn,omitempty"`
	DstAs     string `protobuf:"bytes,39,opt,name=dst_as,json=dstAs,proto3" json:"dst_as,omitempty"`
	DstAsn    uint32 `protobuf:"varint,40,opt,name=dst_asn,json=dstAsn,proto3" json:"dst_asn,omitempty"`
	NexthopAs uint32 `protobuf:"varint,41,opt,name=nexthop_as,json=nexthopAs,proto3" json:"nexthop_as,omitempty"`
	Nexthop   []byte `protobuf:"bytes,42,opt,name=nexthop,proto3" json:"nexthop,omitempty"` // Fixed 16 byte array
	// Geo information
	SrcCity    string `protobuf:"bytes,43,opt,name=src_city,json=srcCity,proto3" json:"src_city,omitempty"`
	SrcCountry string `protobuf:"bytes,44,opt,name=src_country,json=srcCountry,proto3" json:"src_country,omitempty"`
	DstCity    string `protobuf:"bytes,45,opt,name=dst_city,json=dstCity,proto3" json:"dst_city,omitempty"`
	DstCountry string `protobuf:"bytes,46,opt,name=dst_country,json=dstCountry,proto3" json:"dst_country,omitempty"`
	// Network prefix
	SrcNet uint32 `protobuf:"varint,47,opt,name=src_net,json=srcNet,proto3" json:"src_net,omitempty"`
	DstNet uint32 `protobuf:"varint,48,opt,name=dst_net,json=dstNet,proto3" json:"dst_net,omitempty"`
	// Encapsulation information
	HasEncap            bool   `protobuf:"varint,49,opt,name=has_encap,json=hasEncap,proto3" json:"has_encap,omitempty"`
	SrcAddrEncap        []byte `protobuf:"bytes,50,opt,name=src_addr_encap,json=srcAddrEncap,proto3" json:"src_addr_encap,omitempty"` // Fixed 16 byte array
	DstAddrEncap        []byte `protobuf:"bytes,51,opt,name=dst_addr_encap,json=dstAddrEncap,proto3" json:"dst_addr_encap,omitempty"` // Fixed 16 byte array
	ProtoEncap          uint32 `protobuf:"varint,52,opt,name=proto_encap,json=protoEncap,proto3" json:"proto_encap,omitempty"`
	EthertypeEncap      uint32 `protobuf:"varint,53,opt,name=ethertype_encap,json=ethertypeEncap,proto3" json:"ethertype_encap,omitempty"`
	IpTosEncap          uint32 `protobuf:"varint,54,opt,name=ip_tos_encap,json=ipTosEncap,proto3" json:"ip_tos_encap,omitempty"`
	IpTtlEncap          uint32 `protobuf:"varint,55,opt,name=ip_ttl_encap,json=ipTtlEncap,proto3" json:"ip_ttl_encap,omitempty"`
	Ipv6FlowLabelEncap  uint32 `protobuf:"varint,56,opt,name=ipv6_flow_label_encap,json=ipv6FlowLabelEncap,proto3" json:"ipv6_flow_label_encap,omitempty"`
	FragmentIdEncap     uint32 `protobuf:"varint,57,opt,name=fragment_id_encap,json=fragmentIdEncap,proto3" json:"fragment_id_encap,omitempty"`
	FragmentOffsetEncap uint32 `protobuf:"varint,58,opt,name=fragment_offset_encap,json=fragmentOffsetEncap,proto3" json:"fragment_offset_encap,omitempty"`
	// MPLS information
	MplsTtl   []uint32 `protobuf:"varint,80,rep,packed,name=mpls_ttl,json=mplsTtl,proto3" json:"mpls_ttl,omitempty"`
	MplsLabel []uint32 `protobuf:"varint,81,rep,packed,name=mpls_label,json=mplsLabel,proto3" json:"mpls_label,omitempty"`
	MplsIp    [][]byte `protobuf:"bytes,82,rep,name=mpls_ip,json=mplsIp,proto3" json:"mpls_ip,omitempty"`
	// PPP information
	HasPpp            bool   `protobuf:"varint,71,opt,name=has_ppp,json=hasPpp,proto3" json:"has_ppp,omitempty"`
	PppAddressControl uint32 `protobuf:"varint,72,opt,name=ppp_address_control,json=pppAddressControl,proto3" json:"ppp_address_control,omitempty"`
	// BGP information
	BgpNextHop     []byte   `protobuf:"bytes,73,opt,name=bgp_next_hop,json=bgpNextHop,proto3" json:"bgp_next_hop,omitempty"`
	BgpCommunities []uint32 `protobuf:"varint,74,rep,packed,name=bgp_communities,json=bgpCommunities,proto3" json:"bgp_communities,omitempty"`
	AsPath         []uint32 `protobuf:"varint,75,rep,packed,name=as_path,json=asPath,proto3" json:"as_path,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Flow) Reset() {
	*x = Flow{}
	mi := &file_flow_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Flow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flow) ProtoMessage() {}

func (x *Flow) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flow.ProtoReflect.Descriptor instead.
func (*Flow) Descriptor() ([]byte, []int) {
	return file_flow_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Flow) GetFluidVersion() string {
	if x != nil {
		return x.FluidVersion
	}
	return ""
}

func (x *Flow) GetType() Flow_FlowType {
	if x != nil {
		return x.Type
	}
	return Flow_FLOWUNKNOWN
}

func (x *Flow) GetTimeReceived() int64 {
	if x != nil {
		return x.TimeReceived
	}
	return 0
}

func (x *Flow) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Flow) GetSamplingRate() uint32 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

func (x *Flow) GetDirection() uint32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *Flow) GetSamplerAddress() []byte {
	if x != nil {
		return x.SamplerAddress
	}
	return nil
}

func (x *Flow) GetFlowStartTime() int64 {
	if x != nil {
		return x.FlowStartTime
	}
	return 0
}

func (x *Flow) GetFlowEndTime() int64 {
	if x != nil {
		return x.FlowEndTime
	}
	return 0
}

func (x *Flow) GetBytes() uint32 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *Flow) GetPackets() uint32 {
	if x != nil {
		return x.Packets
	}
	return 0
}

func (x *Flow) GetSrcAddr() []byte {
	if x != nil {
		return x.SrcAddr
	}
	return nil
}

func (x *Flow) GetDstAddr() []byte {
	if x != nil {
		return x.DstAddr
	}
	return nil
}

func (x *Flow) GetEthertype() string {
	if x != nil {
		return x.Ethertype
	}
	return ""
}

func (x *Flow) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *Flow) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *Flow) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *Flow) GetInInterface() uint32 {
	if x != nil {
		return x.InInterface
	}
	return 0
}

func (x *Flow) GetOutInterface() uint32 {
	if x != nil {
		return x.OutInterface
	}
	return 0
}

func (x *Flow) GetSrcMac() uint64 {
	if x != nil {
		return x.SrcMac
	}
	return 0
}

func (x *Flow) GetDstMac() uint64 {
	if x != nil {
		return x.DstMac
	}
	return 0
}

func (x *Flow) GetSrcVlan() uint32 {
	if x != nil {
		return x.SrcVlan
	}
	return 0
}

func (x *Flow) GetDstVlan() uint32 {
	if x != nil {
		return x.DstVlan
	}
	return 0
}

func (x *Flow) GetVlanId() uint32 {
	if x != nil {
		return x.VlanId
	}
	return 0
}

func (x *Flow) GetIngressVrfId() uint32 {
	if x != nil {
		return x.IngressVrfId
	}
	return 0
}

func (x *Flow) GetEgressVrfId() uint32 {
	if x != nil {
		return x.EgressVrfId
	}
	return 0
}

func (x *Flow) GetIpTos() uint32 {
	if x != nil {
		return x.IpTos
	}
	return 0
}

func (x *Flow) GetIpTtl() uint32 {
	if x != nil {
		return x.IpTtl
	}
	return 0
}

func (x *Flow) GetTcpFlags() uint32 {
	if x != nil {
		return x.TcpFlags
	}
	return 0
}

func (x *Flow) GetIcmpType() uint32 {
	if x != nil {
		return x.IcmpType
	}
	return 0
}

func (x *Flow) GetIcmpCode() uint32 {
	if x != nil {
		return x.IcmpCode
	}
	return 0
}

func (x *Flow) GetIpv6FlowLabel() uint32 {
	if x != nil {
		return x.Ipv6FlowLabel
	}
	return 0
}

func (x *Flow) GetForwardingStatus() uint32 {
	if x != nil {
		return x.ForwardingStatus
	}
	return 0
}

func (x *Flow) GetFragmentId() uint32 {
	if x != nil {
		return x.FragmentId
	}
	return 0
}

func (x *Flow) GetFragmentOffset() uint32 {
	if x != nil {
		return x.FragmentOffset
	}
	return 0
}

func (x *Flow) GetBiFlowDirection() uint32 {
	if x != nil {
		return x.BiFlowDirection
	}
	return 0
}

func (x *Flow) GetSrcAs() string {
	if x != nil {
		return x.SrcAs
	}
	return ""
}

func (x *Flow) GetSrcAsn() uint32 {
	if x != nil {
		return x.SrcAsn
	}
	return 0
}

func (x *Flow) GetDstAs() string {
	if x != nil {
		return x.DstAs
	}
	return ""
}

func (x *Flow) GetDstAsn() uint32 {
	if x != nil {
		return x.DstAsn
	}
	return 0
}

func (x *Flow) GetNexthopAs() uint32 {
	if x != nil {
		return x.NexthopAs
	}
	return 0
}

func (x *Flow) GetNexthop() []byte {
	if x != nil {
		return x.Nexthop
	}
	return nil
}

func (x *Flow) GetSrcCity() string {
	if x != nil {
		return x.SrcCity
	}
	return ""
}

func (x *Flow) GetSrcCountry() string {
	if x != nil {
		return x.SrcCountry
	}
	return ""
}

func (x *Flow) GetDstCity() string {
	if x != nil {
		return x.DstCity
	}
	return ""
}

func (x *Flow) GetDstCountry() string {
	if x != nil {
		return x.DstCountry
	}
	return ""
}

func (x *Flow) GetSrcNet() uint32 {
	if x != nil {
		return x.SrcNet
	}
	return 0
}

func (x *Flow) GetDstNet() uint32 {
	if x != nil {
		return x.DstNet
	}
	return 0
}

func (x *Flow) GetHasEncap() bool {
	if x != nil {
		return x.HasEncap
	}
	return false
}

func (x *Flow) GetSrcAddrEncap() []byte {
	if x != nil {
		return x.SrcAddrEncap
	}
	return nil
}

func (x *Flow) GetDstAddrEncap() []byte {
	if x != nil {
		return x.DstAddrEncap
	}
	return nil
}

func (x *Flow) GetProtoEncap() uint32 {
	if x != nil {
		return x.ProtoEncap
	}
	return 0
}

func (x *Flow) GetEthertypeEncap() uint32 {
	if x != nil {
		return x.EthertypeEncap
	}
	return 0
}

func (x *Flow) GetIpTosEncap() uint32 {
	if x != nil {
		return x.IpTosEncap
	}
	return 0
}

func (x *Flow) GetIpTtlEncap() uint32 {
	if x != nil {
		return x.IpTtlEncap
	}
	return 0
}

func (x *Flow) GetIpv6FlowLabelEncap() uint32 {
	if x != nil {
		return x.Ipv6FlowLabelEncap
	}
	return 0
}

func (x *Flow) GetFragmentIdEncap() uint32 {
	if x != nil {
		return x.FragmentIdEncap
	}
	return 0
}

func (x *Flow) GetFragmentOffsetEncap() uint32 {
	if x != nil {
		return x.FragmentOffsetEncap
	}
	return 0
}

func (x *Flow) GetMplsTtl() []uint32 {
	if x != nil {
		return x.MplsTtl
	}
	return nil
}

func (x *Flow) GetMplsLabel() []uint32 {
	if x != nil {
		return x.MplsLabel
	}
	return nil
}

func (x *Flow) GetMplsIp() [][]byte {
	if x != nil {
		return x.MplsIp
	}
	return nil
}

func (x *Flow) GetHasPpp() bool {
	if x != nil {
		return x.HasPpp
	}
	return false
}

func (x *Flow) GetPppAddressControl() uint32 {
	if x != nil {
		return x.PppAddressControl
	}
	return 0
}

func (x *Flow) GetBgpNextHop() []byte {
	if x != nil {
		return x.BgpNextHop
	}
	return nil
}

func (x *Flow) GetBgpCommunities() []uint32 {
	if x != nil {
		return x.BgpCommunities
	}
	return nil
}

func (x *Flow) GetAsPath() []uint32 {
	if x != nil {
		return x.AsPath
	}
	return nil
}

var File_flow_v1_proto protoreflect.FileDescriptor

var file_flow_v1_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x22, 0xab, 0x11, 0x0a, 0x04, 0x46, 0x6c, 0x6f,
	0x77, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6c, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6c, 0x75, 0x69, 0x64, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x66, 0x6c, 0x6f, 0x77, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73,
	0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x74, 0x68, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x74, 0x68, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6f, 0x75, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x72, 0x63,
	0x5f, 0x6d, 0x61, 0x63, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x72, 0x63, 0x4d,
	0x61, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x63, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x64, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x72, 0x63, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x72, 0x63, 0x56, 0x6c, 0x61, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x76, 0x6c,
	0x61, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x73, 0x74, 0x56, 0x6c, 0x61,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x72, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x72, 0x66, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x72, 0x66, 0x5f, 0x69,
	0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56,
	0x72, 0x66, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x70, 0x5f, 0x74, 0x6f, 0x73, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x70, 0x54, 0x6f, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x70, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x70, 0x54,
	0x74, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x63, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x69, 0x63, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x63, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x69, 0x63, 0x6d, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x70, 0x76,
	0x36, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x69, 0x5f, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x24, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x73, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63, 0x41, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x72, 0x63, 0x5f, 0x61, 0x73, 0x6e, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x72,
	0x63, 0x41, 0x73, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x73, 0x18, 0x27,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x73, 0x74, 0x41, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x73, 0x74, 0x5f, 0x61, 0x73, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x73,
	0x74, 0x41, 0x73, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x5f,
	0x61, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f,
	0x70, 0x41, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x18, 0x2a,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x72, 0x63, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x72, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74,
	0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x5f, 0x6e, 0x65, 0x74,
	0x18, 0x2f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x72, 0x63, 0x4e, 0x65, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x18, 0x30, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x64, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x65,
	0x6e, 0x63, 0x61, 0x70, 0x18, 0x31, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x45,
	0x6e, 0x63, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x72,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x73,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x63, 0x61, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x6e, 0x63, 0x61,
	0x70, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65,
	0x6e, 0x63, 0x61, 0x70, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x74, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x70,
	0x5f, 0x74, 0x6f, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x69, 0x70, 0x54, 0x6f, 0x73, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0c,
	0x69, 0x70, 0x5f, 0x74, 0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x37, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x69, 0x70, 0x54, 0x74, 0x6c, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x31,
	0x0a, 0x15, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x69,
	0x70, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x63, 0x61,
	0x70, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x39, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x12, 0x32, 0x0a,
	0x15, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x5f, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x61,
	0x70, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x50, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x70, 0x6c, 0x73, 0x54, 0x74, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x51, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x70, 0x6c, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x70, 0x6c, 0x73, 0x5f, 0x69, 0x70, 0x18, 0x52, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x70,
	0x6c, 0x73, 0x49, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x70, 0x70, 0x18,
	0x47, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x61, 0x73, 0x50, 0x70, 0x70, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x70, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x18, 0x48, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x70, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x20, 0x0a,
	0x0c, 0x62, 0x67, 0x70, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x18, 0x49, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x67, 0x70, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x67, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x4a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x67, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x73, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x4b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x73, 0x50, 0x61, 0x74,
	0x68, 0x22, 0x53, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x4c, 0x4f, 0x57, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x35, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4e,
	0x45, 0x54, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x35, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4e,
	0x45, 0x54, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x39, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x49,
	0x50, 0x46, 0x49, 0x58, 0x10, 0x04, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x73, 0x70, 0x69, 0x6c, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x6c, 0x6f, 0x77, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_flow_v1_proto_rawDescOnce sync.Once
	file_flow_v1_proto_rawDescData []byte
)

func file_flow_v1_proto_rawDescGZIP() []byte {
	file_flow_v1_proto_rawDescOnce.Do(func() {
		file_flow_v1_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_flow_v1_proto_rawDesc), len(file_flow_v1_proto_rawDesc)))
	})
	return file_flow_v1_proto_rawDescData
}

var file_flow_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flow_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flow_v1_proto_goTypes = []any{
	(Flow_FlowType)(0), // 0: flow.v1.Flow.FlowType
	(*Flow)(nil),       // 1: flow.v1.Flow
}
var file_flow_v1_proto_depIdxs = []int32{
	0, // 0: flow.v1.Flow.type:type_name -> flow.v1.Flow.FlowType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flow_v1_proto_init() }
func file_flow_v1_proto_init() {
	if File_flow_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_flow_v1_proto_rawDesc), len(file_flow_v1_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_v1_proto_goTypes,
		DependencyIndexes: file_flow_v1_proto_depIdxs,
		EnumInfos:         file_flow_v1_proto_enumTypes,
		MessageInfos:      file_flow_v1_proto_msgTypes,
	}.Build()
	File_flow_v1_proto = out.File
	file_flow_v1_proto_goTypes = nil
	file_flow_v1_proto_depIdxs = nil
}
