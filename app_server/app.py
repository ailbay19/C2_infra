from flask import Flask, make_response, send_file

import requests
import io


app = Flask(__name__)

app.config["FS"] = "http://fs:18081"

fs_url = app.config["FS"]

@app.route("/")
def root():
    
    res = ""
    
    file_res = requests.get(fs_url)
    
    
    if file_res.status_code == 200:
        file = io.BytesIO(file_res.content)
        filename = file_res.headers.get('Content-Disposition')[21:]
        
        res = send_file(file, as_attachment=True, download_name=filename)
        
        res.headers.set('Server', 'NGINX')
    
    return make_response(res, 200)


if __name__ == "__main__":
    app.run(port=18080)