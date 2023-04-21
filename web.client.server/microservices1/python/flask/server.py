# @jeffotoni

from flask import Flask, jsonify, make_response
import requests
from datetime import datetime

app = Flask(__name__)

@app.route('/v1/user', methods=['GET'])
def get_client():
    response = requests.get('http://localhost:3000/v1/avatar')
    customer_data = response.json()
    
    response_headers = {
        "Content-Type": "application/json",
        "Engine": "Python",
        "Location": "/v1/client/post",
        "Date": datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%S.%fZ")
    }

    return make_response(jsonify(customer_data), 200, response_headers)

if __name__ == '__main__':
    app.run(debug=True, port=8080)