# FileUpload

# What it does?
FileUpload does makes possible to resumable upload big files. User can pause/resume/cancel upload.

# What is use of FileUpload?
As it provides functionality of cancel current upload, If user upload a wrong file and any instant of time user get realized then user can cancel current uploading of file.


# How to setup with Docker?

1.Clone this Repository with url : https://github.com/ultra-utsav/FileUpload.git  
git clone https://github.com/ultra-utsav/FileUpload.git

Step 2 : Create Docker image (cmd: "docker build -t file-upload .").

Step 3 : You have docker image. check with (cmd: "docker images").

Step 4 : Run your docker image and that is call docker container.(cmd: "docker run -d -p PORT_U_WANT:8000 file-upload").

Step 5 : Instance is running on https://localhost:PORT_U_WANT/

Step 6 : Find Container ID (cmd: "docker ps") 

Step 7 : Stop running (cmd: "docker container stop CONTAINER_ID")
