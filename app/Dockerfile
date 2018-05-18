# Lightweight alpine OS, weight only 5mb, everything else is Go environment
FROM golang:1.9.2
# Workdir is path in your docker image from where all your commands will be executed
WORKDIR /go/src/github.com/centrypoint/refrigerator/back/go/main
# Add all from your project inside workdir of docker image
ADD . /go/src/github.com/centrypoint/refrigerator/back/go/main
# Then run your script to install dependencies and build application
RUN ./install-deps.sh && go build
# Next start another building context
FROM alpine:latest
# Copy only build result from previous step to new lightweight image
COPY --from=0 /go/src/github.com/centrypoint/regrigerator/back/go/main/main .
# Expose port for access to your app outside of container
EXPOSE 8080

# Starting bundled binary file
CMD [ "./main" ]