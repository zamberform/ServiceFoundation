FROM golang as builder

WORKDIR /src

COPY src /src

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api .

FROM scratch  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/api .

CMD ["./api"] 
