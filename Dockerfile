FROM ubuntu:20.10
WORKDIR /yutu/
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt update
RUN apt install software-properties-common -y
RUN apt-get install -y docker.io build-essential software-properties-common curl
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs
RUN npm install --global yarn
RUN apt-get purge golang*
RUN curl -L https://golang.org/dl/go1.17.2.linux-amd64.tar.gz --output go1.17.2.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz
RUN rm -rf go1.17.2.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN apt-get remove curl -y
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
COPY main /yutu/
