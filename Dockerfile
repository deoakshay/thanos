FROM scratch
 # create a working directory
WORKDIR /
 # run main.go
COPY  app .
CMD ["./app"]