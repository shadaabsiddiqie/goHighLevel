import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { TodoService } from "./gen/proto/todo_connect.js";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

export const todoClient = createPromiseClient(TodoService, transport);
