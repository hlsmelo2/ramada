FROM golang:1.23
WORKDIR /www
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN cp .env.example .env
RUN go build
EXPOSE 3001
CMD ["./api"]