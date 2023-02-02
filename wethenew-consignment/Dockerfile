FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and 
#only redownloading them in subsequent builds if they change

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

# COPY . .
COPY . ./
RUN go build -v -o /usr/local/bin/app 

CMD ["app"]
