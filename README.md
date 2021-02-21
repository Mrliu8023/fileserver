# fileserver
Fileserver based on http.FileServer, and LXC by docker.

# Usage

```shell
docker run -d -p 8000:8000 -v <your_path>:/files fileserver
```
