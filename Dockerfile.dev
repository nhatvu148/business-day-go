FROM ubuntu:21.04

RUN apt-get update
RUN apt-get install -y wget make git build-essential

# Go
WORKDIR /app
RUN wget -c https://dl.google.com/go/go1.18.2.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin
RUN go version
RUN go env

CMD ["make", "server"]