FROM golang
RUN apt-get update && apt-get install -y libpcap-dev --no-install-recommends && rm -rf /var/lib/apt/lists/*
ADD . /go/src/github.com/xianwangs/mylg
RUN go get -x github.com/xianwangs/mylg
CMD ["mylg"]
