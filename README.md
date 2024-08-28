# C2 Server and Client

## Server

We have an app_server for outside requests, and a file_server for internal file uploading and queueing. 

### App Server

- Runs on port 18080
- On root, it requests the File Server for files via http://fs:18081

### File Server

- Runs on port 18081
- On root, it sends 204 if there are no files on the queue. Otherwise sends the file.
- On GET /upload, we have a simple upload html. With POST /upload, we can upload the file.

### NGINX

- Runs on port 80.
- Proxy server is configured as app:18080, where app is docker container name.
- Currently runs on the same bridge network as App and FS servers. Otherwise change nginx.conf to point to an ip.
- Other VM's on the same network can access ports 80, 18080, 18081. Other ports should be closed.

## Scripts

### Local
These are on root.
- Run script runs both the machines on relevant ports on the background. Also creates a pipe for reading server outputs.
- Stop script finds and kills the processes on the relevant ports.
- Upload script with file_path argument sends an upload request to the File Server
- Save transfer script saves the local docker images in a tar file, then sends it to a VM on port 14555 via scp. This should be run after running "docker compose build"

### VM
These are on ./vm_servers
- Setup script loads the tar files on Docker, creates a bridged network.
- Run script runs the docker images on ports 18080 and 18081.
- Upload script is the same as the local one.
