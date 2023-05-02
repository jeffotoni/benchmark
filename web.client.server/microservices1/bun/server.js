// @jeffotoni
import { fetch as externalFetch } from "bun";
import { Agent as HttpAgent } from "http";

const httpAgent = new HttpAgent({ keepAlive: true, keepAliveMsecs: 1000 });

export default {
    port: 8080,
    async fetch(request) {
        const url = new URL(request.url);
        const { pathname } = url;
	//Bun.gc(false);
        if (pathname === "/v1/user" && request.method === "GET") {
            const externalServiceURL = "http://localhost:3000/v1/avatar";
	    Bun.gc(false);
            try {
                
                const externalResponse = await externalFetch(externalServiceURL, {
                  agent: httpAgent,
                });
		//Bun.gc(false);
                if (externalResponse.ok) {
		    //Bun.gc(false);
                    const responseBody = await externalResponse.body;
                    return new Response(responseBody, {
                        headers: { "Content-Type": "application/json" },
                    });
                } else {
                    return new Response("External service error", { status: 500 });
                }
            } catch (error) {
                return new Response("Error fetching external service", { status: 500 });
            }
        } else {
            return new Response("Not Found", { status: 404 });
        }
    },
};
