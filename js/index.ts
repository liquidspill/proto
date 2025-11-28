import * as ControlPlane from "./nexus/controlplane/v1/control_plane_pb";
import * as Permissions from "./nexus/controlplane/v1/permissions_pb";
export const ControlPlaneV1 = { ...ControlPlane, ...Permissions };
export * as CatalogV1 from "./nexus/catalog/v1/catalog_pb";
