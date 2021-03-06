FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build cmd/main.go 
ENV DBNAME=$DBNAME DBUSER=$DBUSER DBPASS=$DBPASS
ENTRYPOINT ["./main"]
EXPOSE 8080 8080