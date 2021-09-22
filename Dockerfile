FROM ubuntu:20.10
WORKDIR /yutu/
RUN apt update
RUN apt-get install -y docker.io golang-go build-essential software-properties-common curl
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs
RUN npm install --global yarn
RUN apt-get remove curl -y
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
COPY main /yutu/
