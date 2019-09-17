FROM golang:1.13-alpine
LABEL version="1.0"
LABEL description="cdpzyafk"
RUN mkdir -p /app
WORKDIR /app
COPY main /app/
CMD ["./main"]
EXPOSE 8080