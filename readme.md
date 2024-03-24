##
docker build -t my-app .
docker run -d -p 8080:8080 my-app
##
docker images
##
docker login
asifsmtp2@gmail.com
8D
##
docker build -t asifsmtp2/go-app:v1.0.0 .
##
docker tag asifsmtp2/new-repo:v1.0.0 asifsmtp2/go-app:latest
##
docker push asifsmtp2/go-app:latest
##
docker run -d -p 8080:8080 asifsmtp2/go-app:latest
##