# Compile stage
FROM golang:1.17 AS build-env
 
# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY . /app
 
# Set working directory
WORKDIR /app

# Run go build to create the executable
RUN go build -o ./server .

# Final stage
FROM gcr.io/distroless/base-debian11

WORKDIR /
EXPOSE 9000
COPY --from=build-env /app/server /

# Run the server executable
CMD ["/server"]