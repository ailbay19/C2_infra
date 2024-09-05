from flask import Flask, make_response, send_file, request, redirect
import os


app = Flask(__name__)

PAYLOADS_DIR = os.path.join(os.path.dirname(__file__), 'payloads')

@app.route("/", methods = ['GET'])
def root():
    body = "Hello"
    
    return make_response(body, 200)

@app.route("/download/<filename>", methods = ['GET'])
def download_file(filename):
    file_path = os.path.join(PAYLOADS_DIR, filename)

    if os.path.isfile(file_path):
        return send_file(file_path, as_attachment=True)
    else:
        return make_response(204)