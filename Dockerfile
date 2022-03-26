FROM golang:1.16-alpine as go-builder

LABEL maintainer="Memory Clutter <memclutter@gmail.com>"

WORKDIR /var/app/

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache curl git ca-certificates tzdata && update-ca-certificates 2>/dev/null || true

# Migrate tool
ENV MIGRATE_VERSION="v4.14.1"
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/$MIGRATE_VERSION/migrate.linux-amd64.tar.gz | tar xvz \
 && mv migrate.linux-amd64 /go/bin/migrate

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/goapp .


FROM alpine

WORKDIR /var/app/

RUN apk update && apk add --no-cache curl

COPY --from=go-builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /go/bin/goapp /go/bin/goapp
COPY --from=go-builder /go/bin/migrate /go/bin/migrate
COPY --from=go-builder /var/app/migrations/ /var/app/migrations/

ENTRYPOINT ["/go/bin/goapp"]
