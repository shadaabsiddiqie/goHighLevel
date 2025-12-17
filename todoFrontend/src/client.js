import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { TodoService } from "./gen/proto/todo_connect.js";
import { URLShorternerService } from "./gen/proto/todo_connect.js";

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_BACKEND_URL || "http://localhost:8080",
});

export const todoClient = createPromiseClient(TodoService, transport);
export const urlClient = createPromiseClient(URLShorternerService, transport);
