# @jeffotoni
# uvicorn server:app --host 127.0.0.1 --port 8080 --workers 4 --log-level critical 
from fastapi import FastAPI, Response
import httpx
from datetime import datetime

app = FastAPI()

@app.get("/v1/user")
async def get_client():
    async with httpx.AsyncClient() as client:
        response = await client.get("http://localhost:3000/v1/avatar")

    customer_data = response.text

    response_headers = {
        "Content-Type": "application/json",
        "Engine": "Python",
        "Location": "/v1/user",
        "Date": datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%S.%fZ"),
    }
    
    return Response(content=customer_data, status_code=200, headers=response_headers)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8080,log_level="critical")
