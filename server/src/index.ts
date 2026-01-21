import { Elysia } from "elysia";
import { Route } from "./routes/main.route";

const app = new Elysia()
  .use(Route)
  .listen(3000);

const host = app.server?.hostname ?? "localhost";
const port = app.server?.port ?? 3000;

console.log(`
────────────────────────────────────────
🦊  ELYSIA SERVER ONLINE
────────────────────────────────────────
🌐  URL      : http://${host}:${port}
⚡  Runtime  : Bun
📦  Mode     : Development
🕒  Started  : ${new Date().toLocaleString()}
────────────────────────────────────────
`);
