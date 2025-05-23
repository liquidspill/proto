// @generated by protoc-gen-es v2.4.0 with parameter "target=ts"
// @generated from file fluid/controlplane/v1/controlplane.proto (package controlplane.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { KeyValue } from "../../../std/v1/std_pb";
import { file_std_v1_std } from "../../../std/v1/std_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file fluid/controlplane/v1/controlplane.proto.
 */
export const file_fluid_controlplane_v1_controlplane: GenFile = /*@__PURE__*/
  fileDesc("CihmbHVpZC9jb250cm9scGxhbmUvdjEvY29udHJvbHBsYW5lLnByb3RvEg9jb250cm9scGxhbmUudjEiSQoVQ29udHJvbE1lc3NhZ2VSZXF1ZXN0EjAKB21lc3NhZ2UYASABKAsyHy5jb250cm9scGxhbmUudjEuQ29udHJvbE1lc3NhZ2UicwoWQ29udHJvbE1lc3NhZ2VSZXNwb25zZRIwCgdtZXNzYWdlGAEgASgLMh8uY29udHJvbHBsYW5lLnYxLkNvbnRyb2xNZXNzYWdlEicKBmNvbmZpZxgCIAEoCzIXLmNvbnRyb2xwbGFuZS52MS5Db25maWciWQoOQ29udHJvbE1lc3NhZ2USPAoQZXhwb3J0X3RlbGVtZXRyeRgBIAEoCzIgLmNvbnRyb2xwbGFuZS52MS5FeHBvcnRUZWxlbWV0cnlIAEIJCgdtZXNzYWdlIiEKD0V4cG9ydFRlbGVtZXRyeRIOCgZleHBvcnQYASABKAgiGAoGQ29uZmlnEg4KBmV4cG9ydBgBIAEoCCJjChNSZWdpc3RlclJ1bGVSZXF1ZXN0EicKBnRhcmdldBgBIAEoCzIXLmNvbnRyb2xwbGFuZS52MS5UYXJnZXQSIwoEcnVsZRgCIAEoCzIVLmNvbnRyb2xwbGFuZS52MS5SdWxlIqQBChRSZWdpc3RlclJ1bGVSZXNwb25zZRItCglhY3Rpb25fYXQYASABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEg8KB3J1bGVfaWQYAiABKAkSJwoGdGFyZ2V0GAMgASgLMhcuY29udHJvbHBsYW5lLnYxLlRhcmdldBIjCgRydWxlGAQgASgLMhUuY29udHJvbHBsYW5lLnYxLlJ1bGUieAoGVGFyZ2V0EisKB2RldmljZXMYASABKAsyGC5jb250cm9scGxhbmUudjEuRGV2aWNlc0gAEjcKCmRlc2NyaXB0b3IYAiABKAsyIS5jb250cm9scGxhbmUudjEuRGV2aWNlRGVzY3JpcHRvckgAQggKBnRhcmdldCIdCgdEZXZpY2VzEhIKCmRldmljZV9pZHMYASADKAkiXwoQRGV2aWNlRGVzY3JpcHRvchIOCgZ2ZW5kb3IYASABKAkSDQoFbW9kZWwYAiABKAkSCgoCb3MYAyABKAkSIAoGbGFiZWxzGAQgAygLMhAuc3RkLnYxLktleVZhbHVlIgYKBFJ1bGUiKAoVRGVyZWdpc3RlclJ1bGVSZXF1ZXN0Eg8KB3J1bGVfaWQYASABKAkiWAoWRGVyZWdpc3RlclJ1bGVSZXNwb25zZRItCglhY3Rpb25fYXQYASABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEg8KB3J1bGVfaWQYAiABKAkytwIKE0NvbnRyb2xQbGFuZVNlcnZpY2USXAoHQ29udHJvbBImLmNvbnRyb2xwbGFuZS52MS5Db250cm9sTWVzc2FnZVJlcXVlc3QaJy5jb250cm9scGxhbmUudjEuQ29udHJvbE1lc3NhZ2VSZXNwb25zZSIAEl0KDFJlZ2lzdGVyUnVsZRIkLmNvbnRyb2xwbGFuZS52MS5SZWdpc3RlclJ1bGVSZXF1ZXN0GiUuY29udHJvbHBsYW5lLnYxLlJlZ2lzdGVyUnVsZVJlc3BvbnNlIgASYwoORGVyZWdpc3RlclJ1bGUSJi5jb250cm9scGxhbmUudjEuRGVyZWdpc3RlclJ1bGVSZXF1ZXN0GicuY29udHJvbHBsYW5lLnYxLkRlcmVnaXN0ZXJSdWxlUmVzcG9uc2UiAEJGWkRnaXRodWIuY29tL2xpcXVpZHNwaWxsL3Byb3RvL2dvL2ZsdWlkL2NvbnRyb2xwbGFuZS92MTtjb250cm9scGxhbmV2MWIGcHJvdG8z", [file_google_protobuf_timestamp, file_std_v1_std]);

/**
 * @generated from message controlplane.v1.ControlMessageRequest
 */
export type ControlMessageRequest = Message<"controlplane.v1.ControlMessageRequest"> & {
  /**
   * @generated from field: controlplane.v1.ControlMessage message = 1;
   */
  message?: ControlMessage;
};

/**
 * Describes the message controlplane.v1.ControlMessageRequest.
 * Use `create(ControlMessageRequestSchema)` to create a new message.
 */
export const ControlMessageRequestSchema: GenMessage<ControlMessageRequest> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 0);

/**
 * @generated from message controlplane.v1.ControlMessageResponse
 */
export type ControlMessageResponse = Message<"controlplane.v1.ControlMessageResponse"> & {
  /**
   * @generated from field: controlplane.v1.ControlMessage message = 1;
   */
  message?: ControlMessage;

  /**
   * @generated from field: controlplane.v1.Config config = 2;
   */
  config?: Config;
};

/**
 * Describes the message controlplane.v1.ControlMessageResponse.
 * Use `create(ControlMessageResponseSchema)` to create a new message.
 */
export const ControlMessageResponseSchema: GenMessage<ControlMessageResponse> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 1);

/**
 * @generated from message controlplane.v1.ControlMessage
 */
export type ControlMessage = Message<"controlplane.v1.ControlMessage"> & {
  /**
   * @generated from oneof controlplane.v1.ControlMessage.message
   */
  message: {
    /**
     * @generated from field: controlplane.v1.ExportTelemetry export_telemetry = 1;
     */
    value: ExportTelemetry;
    case: "exportTelemetry";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message controlplane.v1.ControlMessage.
 * Use `create(ControlMessageSchema)` to create a new message.
 */
export const ControlMessageSchema: GenMessage<ControlMessage> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 2);

/**
 * @generated from message controlplane.v1.ExportTelemetry
 */
export type ExportTelemetry = Message<"controlplane.v1.ExportTelemetry"> & {
  /**
   * @generated from field: bool export = 1;
   */
  export: boolean;
};

/**
 * Describes the message controlplane.v1.ExportTelemetry.
 * Use `create(ExportTelemetrySchema)` to create a new message.
 */
export const ExportTelemetrySchema: GenMessage<ExportTelemetry> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 3);

/**
 * @generated from message controlplane.v1.Config
 */
export type Config = Message<"controlplane.v1.Config"> & {
  /**
   * @generated from field: bool export = 1;
   */
  export: boolean;
};

/**
 * Describes the message controlplane.v1.Config.
 * Use `create(ConfigSchema)` to create a new message.
 */
export const ConfigSchema: GenMessage<Config> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 4);

/**
 * @generated from message controlplane.v1.RegisterRuleRequest
 */
export type RegisterRuleRequest = Message<"controlplane.v1.RegisterRuleRequest"> & {
  /**
   * @generated from field: controlplane.v1.Target target = 1;
   */
  target?: Target;

  /**
   * @generated from field: controlplane.v1.Rule rule = 2;
   */
  rule?: Rule;
};

/**
 * Describes the message controlplane.v1.RegisterRuleRequest.
 * Use `create(RegisterRuleRequestSchema)` to create a new message.
 */
export const RegisterRuleRequestSchema: GenMessage<RegisterRuleRequest> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 5);

/**
 * @generated from message controlplane.v1.RegisterRuleResponse
 */
export type RegisterRuleResponse = Message<"controlplane.v1.RegisterRuleResponse"> & {
  /**
   * @generated from field: google.protobuf.Timestamp action_at = 1;
   */
  actionAt?: Timestamp;

  /**
   * @generated from field: string rule_id = 2;
   */
  ruleId: string;

  /**
   * @generated from field: controlplane.v1.Target target = 3;
   */
  target?: Target;

  /**
   * @generated from field: controlplane.v1.Rule rule = 4;
   */
  rule?: Rule;
};

/**
 * Describes the message controlplane.v1.RegisterRuleResponse.
 * Use `create(RegisterRuleResponseSchema)` to create a new message.
 */
export const RegisterRuleResponseSchema: GenMessage<RegisterRuleResponse> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 6);

/**
 * @generated from message controlplane.v1.Target
 */
export type Target = Message<"controlplane.v1.Target"> & {
  /**
   * @generated from oneof controlplane.v1.Target.target
   */
  target: {
    /**
     * @generated from field: controlplane.v1.Devices devices = 1;
     */
    value: Devices;
    case: "devices";
  } | {
    /**
     * @generated from field: controlplane.v1.DeviceDescriptor descriptor = 2;
     */
    value: DeviceDescriptor;
    case: "descriptor";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message controlplane.v1.Target.
 * Use `create(TargetSchema)` to create a new message.
 */
export const TargetSchema: GenMessage<Target> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 7);

/**
 * @generated from message controlplane.v1.Devices
 */
export type Devices = Message<"controlplane.v1.Devices"> & {
  /**
   * @generated from field: repeated string device_ids = 1;
   */
  deviceIds: string[];
};

/**
 * Describes the message controlplane.v1.Devices.
 * Use `create(DevicesSchema)` to create a new message.
 */
export const DevicesSchema: GenMessage<Devices> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 8);

/**
 * @generated from message controlplane.v1.DeviceDescriptor
 */
export type DeviceDescriptor = Message<"controlplane.v1.DeviceDescriptor"> & {
  /**
   * @generated from field: string vendor = 1;
   */
  vendor: string;

  /**
   * @generated from field: string model = 2;
   */
  model: string;

  /**
   * @generated from field: string os = 3;
   */
  os: string;

  /**
   * @generated from field: repeated std.v1.KeyValue labels = 4;
   */
  labels: KeyValue[];
};

/**
 * Describes the message controlplane.v1.DeviceDescriptor.
 * Use `create(DeviceDescriptorSchema)` to create a new message.
 */
export const DeviceDescriptorSchema: GenMessage<DeviceDescriptor> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 9);

/**
 * @generated from message controlplane.v1.Rule
 */
export type Rule = Message<"controlplane.v1.Rule"> & {
};

/**
 * Describes the message controlplane.v1.Rule.
 * Use `create(RuleSchema)` to create a new message.
 */
export const RuleSchema: GenMessage<Rule> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 10);

/**
 * @generated from message controlplane.v1.DeregisterRuleRequest
 */
export type DeregisterRuleRequest = Message<"controlplane.v1.DeregisterRuleRequest"> & {
  /**
   * @generated from field: string rule_id = 1;
   */
  ruleId: string;
};

/**
 * Describes the message controlplane.v1.DeregisterRuleRequest.
 * Use `create(DeregisterRuleRequestSchema)` to create a new message.
 */
export const DeregisterRuleRequestSchema: GenMessage<DeregisterRuleRequest> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 11);

/**
 * @generated from message controlplane.v1.DeregisterRuleResponse
 */
export type DeregisterRuleResponse = Message<"controlplane.v1.DeregisterRuleResponse"> & {
  /**
   * @generated from field: google.protobuf.Timestamp action_at = 1;
   */
  actionAt?: Timestamp;

  /**
   * @generated from field: string rule_id = 2;
   */
  ruleId: string;
};

/**
 * Describes the message controlplane.v1.DeregisterRuleResponse.
 * Use `create(DeregisterRuleResponseSchema)` to create a new message.
 */
export const DeregisterRuleResponseSchema: GenMessage<DeregisterRuleResponse> = /*@__PURE__*/
  messageDesc(file_fluid_controlplane_v1_controlplane, 12);

/**
 * @generated from service controlplane.v1.ControlPlaneService
 */
export const ControlPlaneService: GenService<{
  /**
   * @generated from rpc controlplane.v1.ControlPlaneService.Control
   */
  control: {
    methodKind: "unary";
    input: typeof ControlMessageRequestSchema;
    output: typeof ControlMessageResponseSchema;
  },
  /**
   * @generated from rpc controlplane.v1.ControlPlaneService.RegisterRule
   */
  registerRule: {
    methodKind: "unary";
    input: typeof RegisterRuleRequestSchema;
    output: typeof RegisterRuleResponseSchema;
  },
  /**
   * @generated from rpc controlplane.v1.ControlPlaneService.DeregisterRule
   */
  deregisterRule: {
    methodKind: "unary";
    input: typeof DeregisterRuleRequestSchema;
    output: typeof DeregisterRuleResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_fluid_controlplane_v1_controlplane, 0);

