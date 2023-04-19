# @jeffotoni
# uvicorn server:app --host 127.0.0.1 --port 8080 --workers 4 --log-level critical 
import requests
from fastapi import FastAPI
from fastapi.responses import JSONResponse
from datetime import datetime

app = FastAPI()

@app.get("/v1/client/get")
async def get_client():
    response = requests.get("http://localhost:3000/v1/customer/get")
    customer_data = response.json()

    response_headers = {
        "Content-Type": "application/json",
        "Engine": "Python",
        "Location": "/v1/client/post",
        "Date": datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%S.%fZ"),
    }

    return JSONResponse(content=customer_data, status_code=200, headers=response_headers)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8080,log_level="critical")
