import Elysia from "elysia";

export function Route() {
  return new Elysia().get("/", "hello")
}