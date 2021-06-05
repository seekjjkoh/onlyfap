import { serve } from "./deps.ts";

const PORT = 1803;

const s = serve({ port: PORT });

console.log("Deno server serving port:", PORT);

for await (const req of s) {
  req.respond({ body: "Hello World" });
}
