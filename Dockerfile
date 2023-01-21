FROM golang:1.19
WORKDIR /app
COPY . .
RUN go build -o main .
#ENTRYPOINT ["/app"]
EXPOSE 8081
CMD ["/app/main", "serve"]