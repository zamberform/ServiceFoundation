FROM golang as builder

WORKDIR /server

COPY server /server

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/server .

FROM scratch  

WORKDIR /root/

COPY server/configs/. ./configs

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/server .

CMD ["./server"] 
