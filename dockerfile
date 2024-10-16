FROM golang:1.22.3
WORKDIR /dockerize
COPY  go.mod .
COPY  . .
RUN go build -o hassan /dockerize
CMD ["go","run","."]

# Build Stage
# FROM golang:1.22.3 AS builder
# WORKDIR /dockerize
# COPY go.mod .
# COPY . .
# RUN go build -o hassan .

# # Final Stage
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /dockerize/hassan .
# ENV GOGC=50
# RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# USER appuser
# EXPOSE 8080
# CMD ["go", "run", ""]

# Optionally add a health check if applicable
# HEALTHCHECK CMD curl --fail http://localhost:8080/health || exit 1
