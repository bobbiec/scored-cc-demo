import requests
import base64

from flask import Flask, request

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.route("/attestation", methods=['POST'])
def attestation():
    data = request.get_json(force=True)
    nonce = base64.urlsafe_b64encode(data['nonce'].encode())
    userData = base64.urlsafe_b64encode(b"This data is authenticated!")
    response = requests.get(
        'http://localhost:50123/api/v1/attestation/report',
        params={
            'nonce': nonce,
            'userData': userData,
        }
    )
    return response.content


if __name__ == "__main__":
    app.run(debug=True, port=8000)
