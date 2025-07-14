FROM quay.io/projectquay/golang:1.15
RUN go get -u github.com/go-sql-driver/mysql
ADD basic_reader.go ./
RUN go build -o basic_reader 
ENV PORT 8080
USER 1001
EXPOSE 8080
CMD ["./basic_reader"]

