FROM golang

COPY . /go/src/github.com/gernest/utron-todo

WorkDir  /go/src/github.com/gernest/utron-todo

RUN go get -v

RUN go install

EXPOSE 8090

CMD ["/go/bin/utron-todo"]
