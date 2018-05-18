FROM golang:1.10.2
# create a working directory
WORKDIR /go/src/app
# add source code
ADD app src
# run main.go
COPY  app src
CMD ["go", "run", "src/main.go"]