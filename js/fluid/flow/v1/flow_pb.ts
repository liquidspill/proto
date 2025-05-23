// @generated by protoc-gen-es v2.4.0 with parameter "target=ts"
// @generated from file fluid/flow/v1/flow.proto (package flow.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_tagger_tagger } from "../../../tagger/tagger_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file fluid/flow/v1/flow.proto.
 */
export const file_fluid_flow_v1_flow: GenFile = /*@__PURE__*/
  fileDesc("ChhmbHVpZC9mbG93L3YxL2Zsb3cucHJvdG8SB2Zsb3cudjEi8h8KBEZsb3cSRwoNZmx1aWRfdmVyc2lvbhgBIAEoCUIwmoSeAytwYXJxdWV0OiJmbHVpZF92ZXJzaW9uLGRpY3QiIGZsdWlkOiJtaW5tYXgiEkwKBHR5cGUYAiABKA4yFi5mbG93LnYxLkZsb3cuRmxvd1R5cGVCJpqEngMhcGFycXVldDoidHlwZSxkaWN0IiBmbHVpZDoiYmxvb20iElIKDXRpbWVfcmVjZWl2ZWQYAyABKANCO5qEngM2cGFycXVldDoidGltZV9yZWNlaXZlZCx0aW1lc3RhbXAsZGVsdGEiIGZsdWlkOiJtaW5tYXgiEjgKCHNlcXVlbmNlGAQgASgNQiaahJ4DIXBhcnF1ZXQ6InNlcXVlbmNlIiBmbHVpZDoibWlubWF4IhJCCg1zYW1wbGluZ19yYXRlGAUgASgNQiuahJ4DJnBhcnF1ZXQ6InNhbXBsaW5nX3JhdGUiIGZsdWlkOiJtaW5tYXgiEj4KCWRpcmVjdGlvbhgGIAEoDUIrmoSeAyZwYXJxdWV0OiJkaXJlY3Rpb24sZGljdCIgZmx1aWQ6ImJsb29tIhJJCg9zYW1wbGVyX2FkZHJlc3MYByABKAxCMJqEngMrcGFycXVldDoic2FtcGxlcl9hZGRyZXNzIiBmbHVpZDoibWlubWF4LGlwIhJWCg9mbG93X3N0YXJ0X3RpbWUYCCABKANCPZqEngM4cGFycXVldDoiZmxvd19zdGFydF90aW1lLHRpbWVzdGFtcCxkZWx0YSIgZmx1aWQ6Im1pbm1heCISUgoNZmxvd19lbmRfdGltZRgJIAEoA0I7moSeAzZwYXJxdWV0OiJmbG93X2VuZF90aW1lLHRpbWVzdGFtcCxkZWx0YSIgZmx1aWQ6Im1pbm1heCISMgoFYnl0ZXMYCiABKA1CI5qEngMecGFycXVldDoiYnl0ZXMiIGZsdWlkOiJtaW5tYXgiEjYKB3BhY2tldHMYCyABKA1CJZqEngMgcGFycXVldDoicGFja2V0cyIgZmx1aWQ6Im1pbm1heCISOwoIc3JjX2FkZHIYDCABKAxCKZqEngMkcGFycXVldDoic3JjX2FkZHIiIGZsdWlkOiJtaW5tYXgsaXAiEjsKCGRzdF9hZGRyGA0gASgMQimahJ4DJHBhcnF1ZXQ6ImRzdF9hZGRyIiBmbHVpZDoibWlubWF4LGlwIhI+CglldGhlcnR5cGUYDiABKAlCK5qEngMmcGFycXVldDoiZXRoZXJ0eXBlLGRpY3QiIGZsdWlkOiJibG9vbSISPAoIcHJvdG9jb2wYDyABKA1CKpqEngMlcGFycXVldDoicHJvdG9jb2wsZGljdCIgZmx1aWQ6ImJsb29tIhI4CghzcmNfcG9ydBgQIAEoDUImmoSeAyFwYXJxdWV0OiJzcmNfcG9ydCIgZmx1aWQ6Im1pbm1heCISOAoIZHN0X3BvcnQYESABKA1CJpqEngMhcGFycXVldDoiZHN0X3BvcnQiIGZsdWlkOiJtaW5tYXgiEkAKDGluX2ludGVyZmFjZRgSIAEoDUIqmoSeAyVwYXJxdWV0OiJpbl9pbnRlcmZhY2UiIGZsdWlkOiJtaW5tYXgiEkIKDW91dF9pbnRlcmZhY2UYEyABKA1CK5qEngMmcGFycXVldDoib3V0X2ludGVyZmFjZSIgZmx1aWQ6Im1pbm1heCISNgoHc3JjX21hYxgUIAEoBEIlmoSeAyBwYXJxdWV0OiJzcmNfbWFjIiBmbHVpZDoibWlubWF4IhI2Cgdkc3RfbWFjGBUgASgEQiWahJ4DIHBhcnF1ZXQ6ImRzdF9tYWMiIGZsdWlkOiJtaW5tYXgiEikKCHNyY192bGFuGBYgASgNQheahJ4DEnBhcnF1ZXQ6InNyY192bGFuIhIpCghkc3RfdmxhbhgXIAEoDUIXmoSeAxJwYXJxdWV0OiJkc3RfdmxhbiISJwoHdmxhbl9pZBgYIAEoDUIWmoSeAxFwYXJxdWV0OiJ2bGFuX2lkIhI1Cg5pbmdyZXNzX3ZyZl9pZBgZIAEoDUIdmoSeAxhwYXJxdWV0OiJpbmdyZXNzX3ZyZl9pZCISMwoNZWdyZXNzX3ZyZl9pZBgaIAEoDUIcmoSeAxdwYXJxdWV0OiJlZ3Jlc3NfdnJmX2lkIhIlCgZpcF90b3MYGyABKA1CFZqEngMQcGFycXVldDoiaXBfdG9zIhIlCgZpcF90dGwYHCABKA1CFZqEngMQcGFycXVldDoiaXBfdHRsIhI+Cgl0Y3BfZmxhZ3MYHSABKA1CK5qEngMmcGFycXVldDoidGNwX2ZsYWdzLGRpY3QiIGZsdWlkOiJibG9vbSISPgoJaWNtcF90eXBlGB4gASgNQiuahJ4DJnBhcnF1ZXQ6ImljbXBfdHlwZSxkaWN0IiBmbHVpZDoiYmxvb20iEj4KCWljbXBfY29kZRgfIAEoDUIrmoSeAyZwYXJxdWV0OiJpY21wX2NvZGUsZGljdCIgZmx1aWQ6ImJsb29tIhI3Cg9pcHY2X2Zsb3dfbGFiZWwYICABKA1CHpqEngMZcGFycXVldDoiaXB2Nl9mbG93X2xhYmVsIhI7ChFmb3J3YXJkaW5nX3N0YXR1cxghIAEoDUIgmoSeAxtwYXJxdWV0OiJmb3J3YXJkaW5nX3N0YXR1cyISLwoLZnJhZ21lbnRfaWQYIiABKA1CGpqEngMVcGFycXVldDoiZnJhZ21lbnRfaWQiEjcKD2ZyYWdtZW50X29mZnNldBgjIAEoDUIemoSeAxlwYXJxdWV0OiJmcmFnbWVudF9vZmZzZXQiEjsKEWJpX2Zsb3dfZGlyZWN0aW9uGCQgASgNQiCahJ4DG3BhcnF1ZXQ6ImJpX2Zsb3dfZGlyZWN0aW9uIhI4CgZzcmNfYXMYJSABKAlCKJqEngMjcGFycXVldDoic3JjX2FzLGRpY3QiIGZsdWlkOiJibG9vbSISOgoHc3JjX2FzbhgmIAEoDUIpmoSeAyRwYXJxdWV0OiJzcmNfYXNuLGRpY3QiIGZsdWlkOiJibG9vbSISOAoGZHN0X2FzGCcgASgJQiiahJ4DI3BhcnF1ZXQ6ImRzdF9hcyxkaWN0IiBmbHVpZDoiYmxvb20iEjoKB2RzdF9hc24YKCABKA1CKZqEngMkcGFycXVldDoiZHN0X2FzbixkaWN0IiBmbHVpZDoiYmxvb20iEi0KCm5leHRob3BfYXMYKSABKA1CGZqEngMUcGFycXVldDoibmV4dGhvcF9hcyISOQoHbmV4dGhvcBgqIAEoDEIomoSeAyNwYXJxdWV0OiJuZXh0aG9wIiBmbHVpZDoibWlubWF4LGlwIhI8CghzcmNfY2l0eRgrIAEoCUIqmoSeAyVwYXJxdWV0OiJzcmNfY2l0eSxkaWN0IiBmbHVpZDoiYmxvb20iEkIKC3NyY19jb3VudHJ5GCwgASgJQi2ahJ4DKHBhcnF1ZXQ6InNyY19jb3VudHJ5LGRpY3QiIGZsdWlkOiJibG9vbSISPAoIZHN0X2NpdHkYLSABKAlCKpqEngMlcGFycXVldDoiZHN0X2NpdHksZGljdCIgZmx1aWQ6ImJsb29tIhJCCgtkc3RfY291bnRyeRguIAEoCUItmoSeAyhwYXJxdWV0OiJkc3RfY291bnRyeSxkaWN0IiBmbHVpZDoiYmxvb20iEjYKB3NyY19uZXQYLyABKA1CJZqEngMgcGFycXVldDoic3JjX25ldCIgZmx1aWQ6Im1pbm1heCISNgoHZHN0X25ldBgwIAEoDUIlmoSeAyBwYXJxdWV0OiJkc3RfbmV0IiBmbHVpZDoibWlubWF4IhIrCgloYXNfZW5jYXAYMSABKAhCGJqEngMTcGFycXVldDoiaGFzX2VuY2FwIhJHCg5zcmNfYWRkcl9lbmNhcBgyIAEoDEIvmoSeAypwYXJxdWV0OiJzcmNfYWRkcl9lbmNhcCIgZmx1aWQ6Im1pbm1heCxpcCISRwoOZHN0X2FkZHJfZW5jYXAYMyABKAxCL5qEngMqcGFycXVldDoiZHN0X2FkZHJfZW5jYXAiIGZsdWlkOiJtaW5tYXgsaXAiEi8KC3Byb3RvX2VuY2FwGDQgASgNQhqahJ4DFXBhcnF1ZXQ6InByb3RvX2VuY2FwIhI3Cg9ldGhlcnR5cGVfZW5jYXAYNSABKA1CHpqEngMZcGFycXVldDoiZXRoZXJ0eXBlX2VuY2FwIhIxCgxpcF90b3NfZW5jYXAYNiABKA1CG5qEngMWcGFycXVldDoiaXBfdG9zX2VuY2FwIhIxCgxpcF90dGxfZW5jYXAYNyABKA1CG5qEngMWcGFycXVldDoiaXBfdHRsX2VuY2FwIhJDChVpcHY2X2Zsb3dfbGFiZWxfZW5jYXAYOCABKA1CJJqEngMfcGFycXVldDoiaXB2Nl9mbG93X2xhYmVsX2VuY2FwIhI7ChFmcmFnbWVudF9pZF9lbmNhcBg5IAEoDUIgmoSeAxtwYXJxdWV0OiJmcmFnbWVudF9pZF9lbmNhcCISQwoVZnJhZ21lbnRfb2Zmc2V0X2VuY2FwGDogASgNQiSahJ4DH3BhcnF1ZXQ6ImZyYWdtZW50X29mZnNldF9lbmNhcCISKQoIbXBsc190dGwYUCADKA1CF5qEngMScGFycXVldDoibXBsc190dGwiEi0KCm1wbHNfbGFiZWwYUSADKA1CGZqEngMUcGFycXVldDoibXBsc19sYWJlbCISJwoHbXBsc19pcBhSIAMoDEIWmoSeAxFwYXJxdWV0OiJtcGxzX2lwIhInCgdoYXNfcHBwGEcgASgIQhaahJ4DEXBhcnF1ZXQ6Imhhc19wcHAiEj8KE3BwcF9hZGRyZXNzX2NvbnRyb2wYSCABKA1CIpqEngMdcGFycXVldDoicHBwX2FkZHJlc3NfY29udHJvbCISMQoMYmdwX25leHRfaG9wGEkgASgMQhuahJ4DFnBhcnF1ZXQ6ImJncF9uZXh0X2hvcCISNwoPYmdwX2NvbW11bml0aWVzGEogAygNQh6ahJ4DGXBhcnF1ZXQ6ImJncF9jb21tdW5pdGllcyISJwoHYXNfcGF0aBhLIAMoDUIWmoSeAxFwYXJxdWV0OiJhc19wYXRoIhI8CghzZWdtZW50cxhMIAMoCUIqmoSeAyVwYXJxdWV0OiJzZWdtZW50cyxkaWN0IiBmbHVpZDoiYmxvb20iIoUBCghGbG93VHlwZRIZChVGTE9XX1RZUEVfVU5TUEVDSUZJRUQQABIVChFGTE9XX1RZUEVfU0ZMT1dfNRABEhgKFEZMT1dfVFlQRV9ORVRGTE9XX1Y1EAISGAoURkxPV19UWVBFX05FVEZMT1dfVjkQAxITCg9GTE9XX1RZUEVfSVBGSVgQBEI2WjRnaXRodWIuY29tL2xpcXVpZHNwaWxsL3Byb3RvL2dvL2ZsdWlkL2Zsb3cvdjE7Zmxvd3YxYgZwcm90bzM", [file_tagger_tagger]);

/**
 * @generated from message flow.v1.Flow
 */
export type Flow = Message<"flow.v1.Flow"> & {
  /**
   * @generated from field: string fluid_version = 1;
   */
  fluidVersion: string;

  /**
   * @generated from field: flow.v1.Flow.FlowType type = 2;
   */
  type: Flow_FlowType;

  /**
   * @generated from field: int64 time_received = 3;
   */
  timeReceived: bigint;

  /**
   * @generated from field: uint32 sequence = 4;
   */
  sequence: number;

  /**
   * @generated from field: uint32 sampling_rate = 5;
   */
  samplingRate: number;

  /**
   * @generated from field: uint32 direction = 6;
   */
  direction: number;

  /**
   * Sampler information
   *
   * Fixed 16 byte array
   *
   * @generated from field: bytes sampler_address = 7;
   */
  samplerAddress: Uint8Array;

  /**
   * Flow timing
   *
   * @generated from field: int64 flow_start_time = 8;
   */
  flowStartTime: bigint;

  /**
   * @generated from field: int64 flow_end_time = 9;
   */
  flowEndTime: bigint;

  /**
   * Size metrics
   *
   * @generated from field: uint32 bytes = 10;
   */
  bytes: number;

  /**
   * @generated from field: uint32 packets = 11;
   */
  packets: number;

  /**
   * Network addresses
   *
   * Fixed 16 byte array
   *
   * @generated from field: bytes src_addr = 12;
   */
  srcAddr: Uint8Array;

  /**
   * Fixed 16 byte array
   *
   * @generated from field: bytes dst_addr = 13;
   */
  dstAddr: Uint8Array;

  /**
   * @generated from field: string ethertype = 14;
   */
  ethertype: string;

  /**
   * @generated from field: uint32 protocol = 15;
   */
  protocol: number;

  /**
   * Ports
   *
   * @generated from field: uint32 src_port = 16;
   */
  srcPort: number;

  /**
   * @generated from field: uint32 dst_port = 17;
   */
  dstPort: number;

  /**
   * Interfaces
   *
   * @generated from field: uint32 in_interface = 18;
   */
  inInterface: number;

  /**
   * @generated from field: uint32 out_interface = 19;
   */
  outInterface: number;

  /**
   * Ethernet information
   *
   * @generated from field: uint64 src_mac = 20;
   */
  srcMac: bigint;

  /**
   * @generated from field: uint64 dst_mac = 21;
   */
  dstMac: bigint;

  /**
   * VLAN information
   *
   * @generated from field: uint32 src_vlan = 22;
   */
  srcVlan: number;

  /**
   * @generated from field: uint32 dst_vlan = 23;
   */
  dstVlan: number;

  /**
   * @generated from field: uint32 vlan_id = 24;
   */
  vlanId: number;

  /**
   * VRF information
   *
   * @generated from field: uint32 ingress_vrf_id = 25;
   */
  ingressVrfId: number;

  /**
   * @generated from field: uint32 egress_vrf_id = 26;
   */
  egressVrfId: number;

  /**
   * IP and TCP flags
   *
   * @generated from field: uint32 ip_tos = 27;
   */
  ipTos: number;

  /**
   * @generated from field: uint32 ip_ttl = 28;
   */
  ipTtl: number;

  /**
   * @generated from field: uint32 tcp_flags = 29;
   */
  tcpFlags: number;

  /**
   * @generated from field: uint32 icmp_type = 30;
   */
  icmpType: number;

  /**
   * @generated from field: uint32 icmp_code = 31;
   */
  icmpCode: number;

  /**
   * @generated from field: uint32 ipv6_flow_label = 32;
   */
  ipv6FlowLabel: number;

  /**
   * @generated from field: uint32 forwarding_status = 33;
   */
  forwardingStatus: number;

  /**
   * Fragment information
   *
   * @generated from field: uint32 fragment_id = 34;
   */
  fragmentId: number;

  /**
   * @generated from field: uint32 fragment_offset = 35;
   */
  fragmentOffset: number;

  /**
   * @generated from field: uint32 bi_flow_direction = 36;
   */
  biFlowDirection: number;

  /**
   * AS information
   *
   * @generated from field: string src_as = 37;
   */
  srcAs: string;

  /**
   * @generated from field: uint32 src_asn = 38;
   */
  srcAsn: number;

  /**
   * @generated from field: string dst_as = 39;
   */
  dstAs: string;

  /**
   * @generated from field: uint32 dst_asn = 40;
   */
  dstAsn: number;

  /**
   * @generated from field: uint32 nexthop_as = 41;
   */
  nexthopAs: number;

  /**
   * Fixed 16 byte array
   *
   * @generated from field: bytes nexthop = 42;
   */
  nexthop: Uint8Array;

  /**
   * Geo information
   *
   * @generated from field: string src_city = 43;
   */
  srcCity: string;

  /**
   * @generated from field: string src_country = 44;
   */
  srcCountry: string;

  /**
   * @generated from field: string dst_city = 45;
   */
  dstCity: string;

  /**
   * @generated from field: string dst_country = 46;
   */
  dstCountry: string;

  /**
   * Network prefix
   *
   * @generated from field: uint32 src_net = 47;
   */
  srcNet: number;

  /**
   * @generated from field: uint32 dst_net = 48;
   */
  dstNet: number;

  /**
   * Encapsulation information
   *
   * @generated from field: bool has_encap = 49;
   */
  hasEncap: boolean;

  /**
   * Fixed 16 byte array
   *
   * @generated from field: bytes src_addr_encap = 50;
   */
  srcAddrEncap: Uint8Array;

  /**
   * Fixed 16 byte array
   *
   * @generated from field: bytes dst_addr_encap = 51;
   */
  dstAddrEncap: Uint8Array;

  /**
   * @generated from field: uint32 proto_encap = 52;
   */
  protoEncap: number;

  /**
   * @generated from field: uint32 ethertype_encap = 53;
   */
  ethertypeEncap: number;

  /**
   * @generated from field: uint32 ip_tos_encap = 54;
   */
  ipTosEncap: number;

  /**
   * @generated from field: uint32 ip_ttl_encap = 55;
   */
  ipTtlEncap: number;

  /**
   * @generated from field: uint32 ipv6_flow_label_encap = 56;
   */
  ipv6FlowLabelEncap: number;

  /**
   * @generated from field: uint32 fragment_id_encap = 57;
   */
  fragmentIdEncap: number;

  /**
   * @generated from field: uint32 fragment_offset_encap = 58;
   */
  fragmentOffsetEncap: number;

  /**
   * MPLS information
   *
   * @generated from field: repeated uint32 mpls_ttl = 80;
   */
  mplsTtl: number[];

  /**
   * @generated from field: repeated uint32 mpls_label = 81;
   */
  mplsLabel: number[];

  /**
   * @generated from field: repeated bytes mpls_ip = 82;
   */
  mplsIp: Uint8Array[];

  /**
   * PPP information
   *
   * @generated from field: bool has_ppp = 71;
   */
  hasPpp: boolean;

  /**
   * @generated from field: uint32 ppp_address_control = 72;
   */
  pppAddressControl: number;

  /**
   * BGP information
   *
   * @generated from field: bytes bgp_next_hop = 73;
   */
  bgpNextHop: Uint8Array;

  /**
   * @generated from field: repeated uint32 bgp_communities = 74;
   */
  bgpCommunities: number[];

  /**
   * @generated from field: repeated uint32 as_path = 75;
   */
  asPath: number[];

  /**
   * Spilled additional metadata
   *
   * @generated from field: repeated string segments = 76;
   */
  segments: string[];
};

/**
 * Describes the message flow.v1.Flow.
 * Use `create(FlowSchema)` to create a new message.
 */
export const FlowSchema: GenMessage<Flow> = /*@__PURE__*/
  messageDesc(file_fluid_flow_v1_flow, 0);

/**
 * @generated from enum flow.v1.Flow.FlowType
 */
export enum Flow_FlowType {
  /**
   * @generated from enum value: FLOW_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: FLOW_TYPE_SFLOW_5 = 1;
   */
  SFLOW_5 = 1,

  /**
   * @generated from enum value: FLOW_TYPE_NETFLOW_V5 = 2;
   */
  NETFLOW_V5 = 2,

  /**
   * @generated from enum value: FLOW_TYPE_NETFLOW_V9 = 3;
   */
  NETFLOW_V9 = 3,

  /**
   * @generated from enum value: FLOW_TYPE_IPFIX = 4;
   */
  IPFIX = 4,
}

/**
 * Describes the enum flow.v1.Flow.FlowType.
 */
export const Flow_FlowTypeSchema: GenEnum<Flow_FlowType> = /*@__PURE__*/
  enumDesc(file_fluid_flow_v1_flow, 0, 0);

