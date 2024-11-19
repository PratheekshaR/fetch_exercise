FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o receipt-processor .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./receipt-processor"]
