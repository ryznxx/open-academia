import Elysia from "elysia";

export default function AuthRoute() {
  return new Elysia().group("/auth", (app) =>
    app
      .get("/sso", () => "login route")
  );
}