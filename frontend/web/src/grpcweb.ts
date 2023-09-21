import { createChannel, createClientFactory, FetchTransport } from "nice-grpc-web";
import { UserServiceDefinition } from "./types/proto/api/v2/user_service";
import { WorkspaceServiceDefinition } from "./types/proto/api/v2/workspace_service";

const address = import.meta.env.MODE === "development" ? "http://localhost:8082" : window.location.origin;

const channel = createChannel(
  address,
  FetchTransport({
    credentials: "include",
  })
);

const clientFactory = createClientFactory();

export const userServiceClient = clientFactory.create(UserServiceDefinition, channel);

export const workspaceServiceClient = clientFactory.create(WorkspaceServiceDefinition, channel);
