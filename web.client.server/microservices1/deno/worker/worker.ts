// worker.ts
self.onmessage = async (event) => {
  const reqEvent = event.data;
  const req = new Request(reqEvent.url, { method: reqEvent.method, headers: reqEvent.headers });
  const url = new URL(req.url);

  if (req.method === "GET" && url.pathname === "/v1/client/get") {
    try {
      const response = await fetch("http://localhost:3000/v1/customer/get");
      if (response.ok) {
        const data = await response.json();
        self.postMessage({ status: 200, body: JSON.stringify(data) });
      } else {
        self.postMessage({ status: 500, body: "Erro ao buscar os dados do cliente." });
      }
    } catch (error) {
      self.postMessage({ status: 500, body: "Erro ao buscar os dados do cliente." });
    }
  } else {
    self.postMessage({ status: 404, body: "Not Found" });
  }
};

