# C2 Server and Client

## Server

We have an app_server and an NGINX reverse proxy.

### App Server

-   Runs on port 18080
-   On root, it sends the file in the queue or returns 204.
-   On /upload, we can POST a file in the queue. Only accessible through container shell.

### NGINX

-   Runs on port 80 and 443.
-   Proxy server is configured as app:18080, where app is docker container name.
-   Runs on the same bridge network as App servers.
-   Requires SSL certificates. Port 80 redirects to Google. If server certificate but no client certificate, then port 443 redirects to Google. Otherwise redirect to app.

## Scripts

### Local

These are on root.

-   Run script runs both the machines on relevant ports on the background. Also creates a pipe for reading server outputs.
-   [DEPRECATED] Stop script finds and kills the processes on the relevant ports.
-   [DEPRECATED] Upload script with file_path argument sends an upload request to the File Server
-   Save transfer script saves the local docker images in a tar file, then sends it to a VM on port 14555 via scp. This should be run after running "docker compose build"
-   Clean script removes all the container and images
-   Generate keys script generates certificates and keys and puts them in client, nginx and vm_servers directories.
-   All script cleans, builds and save_transfer's.

### VM

These are on ./vm_servers

-   Setup script loads the tar files on Docker, creates a bridged network.
-   Run script runs the docker images on ports 443 and 18080.
-   [DEPRECATED] Upload script is the same as the local one.
-   Clean script removes all the container and images
