FROM golang:1.15-alpine AS build_base
# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN apk add tzdata
# We want to populate the module cache based on the go.{mod,sum} files.
WORKDIR /app
# Here we copy the rest of the source code
COPY . .
#change time to local timezone
RUN echo Asia/Almaty > /etc/timezone
#And compile the project
RUN go build ./main.go
# Finally we copy the statically compiled Go binary.
# Run the executable
ENTRYPOINT ["./main"]