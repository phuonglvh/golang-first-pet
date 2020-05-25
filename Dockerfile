FROM golang:1.14.3-alpine as first-pet-build

WORKDIR /go/src/github.com/phuonglvh/golang-first-pet

# Copy the local package files to the container's workspace.
COPY . ./
RUN (apk add --update --no-cache git)

# install dependencies
RUN go get -u ./...

# Build
RUN go install github.com/phuonglvh/golang-first-pet

# production image
FROM golang:1.14.3-alpine as golang-multi-room-chat
WORKDIR /go/bin
COPY --from=first-pet-build /go/bin/golang-first-pet ./golang-multi-room-chat
RUN ls -la

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the outyet command by default when the container starts.
ENTRYPOINT ["/go/bin/golang-multi-room-chat"]