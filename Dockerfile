FROM golang:1.14.3-alpine as app_build
ENV APP_HOME /go/src/github.com/phuonglvh/golang-first-pet
WORKDIR $APP_HOME
RUN (apk add --update --no-cache git)

# install dependencies

# Copy the local package files to the container's workspace.
COPY . ./

# Build
RUN go build -o go-multi-room-chat

# production image
FROM golang:1.14.3-alpine as app_prod
ENV APP_HOME /go/src/github.com/phuonglvh/golang-first-pet
WORKDIR $APP_HOME
COPY --from=app_build $APP_HOME/go-multi-room-chat ./
COPY ["./static/", "./static"]
COPY ["./app/views", "./app/views"]

# Document that the service listens on port 8080.
EXPOSE 8080
RUN ls -la
# Run the outyet command by default when the container starts.
ENTRYPOINT ["./go-multi-room-chat"]
