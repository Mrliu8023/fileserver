# fileserver
Fileserver based on http.FileServer, and LXC by Docker.

# LXC Usage
First build DockerFile
```shell
docker build -t fileserver .
```
Then run
```shell
docker run -d -p 8000:8000 -v <your_path>:/files fileserver
```
