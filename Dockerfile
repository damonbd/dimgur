FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /app/dimgur-go /app/

# # have to copy over sub directories
# WORKDIR /Views
# COPY ./Views .

WORKDIR ..

EXPOSE 8080:8080
ENTRYPOINT ["/app/dimgur-go"]

#sudo docker build -t dimgur-go .
#sudo docker run -p 8080:8080 dimgur-go .

