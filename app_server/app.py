from flask import Flask, make_response, send_file, request, redirect

from file_manager import FileManager
import io


app = Flask(__name__)
fm = FileManager()

@app.route("/")
def root():
    response = ""

    file_resource = fm.get()
    if file_resource:
        file = io.BytesIO(file_resource["file"])
        name = file_resource["filename"]
        
        response = send_file(file, as_attachment = True, download_name = name)
    
        return make_response(response, 200)
    
    else:
        return make_response(response, 204)





@app.route("/upload", methods = ['GET', 'POST'])
def upload():
    if request.method == 'GET':
        response =  '''
                        <!DOCTYPE html>
                        <form method="post" enctype="multipart/form-data">
                            <input type="file" name="file" />
                            <input type="submit" value="Upload" />
                        </form>
                    '''
                    
        
        return make_response(response, 200)
    
    if request.method == 'POST':
        response = ""
        
        if 'file' not in request.files:
            return redirect(request.url)
        
        file = request.files['file']

        if file.filename == '':
            return redirect(request.url)

        if not file:
            return redirect(request.url)
        
        data = file.stream.read()
        

        file_res = fm.put(data)
        if file_res:
            response = file_res["filename"]
            return make_response(response, 201)
        else:
            return make_response(response, 400)


if __name__ == "__main__":
    app.run(port=18080)