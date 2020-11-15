
# What it does?
FileUpload does makes possible to resumable upload big files. User can pause/resume/cancel upload.

# What is use of FileUpload?
As it provides functionality of cancel current upload, If user upload a wrong file and any instant of time user get realized then user can cancel current uploading of file.


# How to setup with Docker?

Step 1 : Clone this Repository 
```
git clone https://github.com/ultra-utsav/FileUpload.git
```

Step 2 : Create Docker image 
```
docker build -t file-upload .
```

Step 3 : You have docker image. check with 
```
docker images
```

Step 4 : Run your docker image and that is call docker container.
```
docker run -d -p PORT_U_WANT:8000 file-upload
```

Step 5 : Instance is running on ```https://localhost:PORT_U_WANT/```

Step 6 : Find Container ID 
```
docker ps
```

Step 7 : Stop running 
```
docker container stop CONTAINER_ID
```



# Where Uploaded files stored in docker?
It can be find using applying SSH into a running docker container.

Step 1 : Find Name of current running container. 
```
docker ps   
```

Step 2 : Get get a bash shell in the container 
```
docker exec -it <container name> /bin/bash
```

Step 3 : You can run any command here like 
```
docker exec -it <container name> <command>
```

Step 4 : In ```./files``` where all uploaded files will be there if not uploaded any files then there is no folder with ```files```.

# What are used to develope?

1. Golang for backend routing.
2. Resumable.js for dividing whole file into chunks and send to backend.
3. HTML/CSS.
4. Docker.




