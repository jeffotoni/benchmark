// @jeffotoni
import { fetch as externalFetch } from "bun";
export default {
    port: 8080,
    async fetch(request) {
        const url = new URL(request.url);
        const { pathname } = url;
        if (pathname === "/v1/user" && request.method === "GET") {
            const externalServiceURL = "http://localhost:3000/v1/avatar";
            try {
                const externalResponse = await externalFetch(externalServiceURL);
                if (externalResponse.ok) {
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