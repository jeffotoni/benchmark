// Criando o servidor HTTP na porta 8080
const server = Deno.listen({ hostname: "0.0.0.0", port: 8080 });
console.log("Servidor rodando em 0.0.0.0:8080");

for await (const conn of server) {
  (async () => {
    
    const httpConn = Deno.serveHttp(conn);
    
    for await (const reqEvent of httpConn) {
      
      const req = reqEvent.request;
      const url = new URL(req.url);
      
      if (req.method === "GET" && url.pathname === "/v1/user") {
        try {
          const response = await fetch("http://localhost:3000/v1/avatar");      
          if (response.ok) {
            const data = await response.json();
            reqEvent.respondWith(new Response(JSON.stringify(data), { status: 200 }));
          } else {
            reqEvent.respondWith(new Response("Erro ao buscar os dados do cliente.", { status: 500 }));
          }
        } catch (error) {
          reqEvent.respondWith(new Response("Erro ao buscar os dados do cliente.", { status: 500 }));
        }
      } else {
        reqEvent.respondWith(new Response("Not Found", { status: 404 }));
      }
    }
  })();
}
