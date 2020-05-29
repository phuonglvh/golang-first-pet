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
RUN (apk add tree)
COPY --from=app_build $APP_HOME/go-multi-room-chat ./
COPY ["./static/", "./static"]
COPY ["./app/views", "./app/views"]
RUN tree -L 4

# default environment variables
ENV MODE PRODUCTION
ENV SERVER_HOST 0.0.0.0
ENV SERVER_PORT 8080
ENV CHAT_MESSAGE_LIFETIME 1

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the outyet command by default when the container starts.
ENTRYPOINT ["sh", "-c", "MODE=${MODE} SERVER_HOST=${SERVER_HOST} SERVER_PORT=${SERVER_PORT} CHAT_MESSAGE_LIFETIME=${CHAT_MESSAGE_LIFETIME} ./go-multi-room-chat"]
