FROM ubuntu:20.10
WORKDIR /yutu/
RUN apt update
RUN apt-get install golang-go -y
RUN apt-get install docker.io -y
RUN apt-get install build-essential -y
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
COPY main /yutu/
