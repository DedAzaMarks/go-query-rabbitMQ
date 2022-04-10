FROM rabbitmq
RUN rabbitmq-server -detached &
RUN rabbitmq-plugins enable rabbitmq_management
RUN apt-get update -y && apt-get upgrade -y && apt-get install -y wget curl
RUN wget https://dl.google.com/go/go1.17.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.8.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
EXPOSE 42069
RUN mkdir app
WORKDIR app
COPY . .
RUN /usr/local/go/bin/go mod tidy
CMD while ! curl -s rabbitmq:15672 > /dev/null; do echo waiting for rabbitmq; sleep 3; done; /usr/local/go/bin/go run ./cmd/pathfinder/main.go & /usr/local/go/bin/go run server/*.go