import { WorkerPool } from "./worker_pool.ts";

const server = Deno.listen({ hostname: "0.0.0.0", port: 8080 });
console.log("Servidor rodando em 0.0.0.0:8080");

const workerPool = new WorkerPool(4, new URL("worker.ts", import.meta.url).href);

for await (const conn of server) {
  (async () => {
    const httpConn = Deno.serveHttp(conn);
    for await (const reqEvent of httpConn) {
      const request = {
        url: reqEvent.request.url,
        method: reqEvent.request.method,
        headers: [...reqEvent.request.headers],
      };
      
      const response = await workerPool.run(request);
      reqEvent.respondWith(new Response(response.body, { status: response.status }));
    }
  })();
}
