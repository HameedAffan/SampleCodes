FROM golang:1.12.0
RUN mkdir /newfile
WORKDIR / newfile
COPY go.mod
ADD ./main.go
RUN go build -o main.
CMD ["/newfile/main"]