import Elysia from "elysia";
import AuthRoute from "./auth/auth.route";

export function Route() {
  return new Elysia().use(AuthRoute());
}