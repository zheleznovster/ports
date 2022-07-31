FROM golang:1.18-alpine as builder

WORKDIR /

COPY . ./

# project uses no dependencies, so `go mod download` step is optional, but it's OK to keep it
RUN go mod download

RUN go build -o /ports

# use a fresh scratch layer and copy only the compiled binary from builder, but do not copy the code
FROM scratch

COPY --from=builder /ports ./
COPY --from=builder /data/ports.json ./data/

CMD [ "/ports" ]
