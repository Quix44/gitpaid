# Start with the Go base image for building the binary
FROM golang:1.21.5 as build
WORKDIR /github-handler

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code from the src directory
COPY src/. ./

# Build the Go application
RUN go build -tags lambda.norpc -o main .

# Second stage: build the small runtime image.
FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /github-handler/main ./main

ENTRYPOINT [ "./main" ]
