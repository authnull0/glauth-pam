FROM artifactory.authnull.com/authnull/golang:1.19-alpine

USER root

ENV GOCACHE /tmp

RUN mkdir /app

COPY ./main /app

RUN chmod a+w /app

RUN chmod +x /app/main

WORKDIR /app

EXPOSE 3893

CMD ["./main",  "-c", "sample-simple.cfg"]