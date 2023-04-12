FROM artifactory.authnull.com/authnull/golang:1.19-alpine

RUN apk update
# RUN go mod download && \
#     go mod verify && \ 
#     go mod tidy
RUN chmod 777 /root/.m2/build/resources/auth.sh
RUN go build -o main ./app/service.go
RUN sh /root/.m2/build/resources/auth.sh
COPY /root/.m2/app/service.go  /root/.m2/build/deploy/docker/
COPY /root/.m2/main  /root/.m2/build/deploy/docker/

WORKDIR /root/.m2/
USER root

ENV GOCACHE /tmp

RUN mkdir /app

COPY ./main /app

RUN chmod a+w /app

RUN chmod +x /app/main

WORKDIR /app

EXPOSE 3893

CMD ["./main",  "-c", "sample-simple.cfg"]