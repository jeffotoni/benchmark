const server = Deno.listen({ hostname: "0.0.0.0", port: 8080 });
console.log("Servidor rodando em 0.0.0.0:8080");

for await (const conn of server) {
  (async () => {
    const httpConn = Deno.serveHttp(conn);
    for await (const reqEvent of httpConn) {
      const worker = new Worker(new URL("worker.ts", import.meta.url).href, { type: "module" });
      worker.postMessage({
        url: reqEvent.request.url,
        method: reqEvent.request.method,
        headers: [...reqEvent.request.headers],
      });

      worker.onmessage = (event) => {
        const { status, body } = event.data;
        reqEvent.respondWith(new Response(body, { status }));
        worker.terminate();
      };
    }
  })();
}