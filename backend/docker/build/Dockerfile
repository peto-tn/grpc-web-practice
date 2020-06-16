ARG ALPINE_VERSION=3.10.3
###################
# STAGE 1: builder
###################

FROM alpine:$ALPINE_VERSION as builder

WORKDIR /app/source

RUN apk add --no-cache go \
        git \
        binutils-gold \
        curl \
        g++ \
        gcc \
        gnupg \
        libgcc \
        linux-headers \
        make \
        python
ADD . .
RUN go build && strip blog

###################
# STAGE 2: runner
###################

FROM alpine:$ALPINE_VERSION as runner

WORKDIR /app

RUN apk update && apk add --no-cache libc6-compat ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/source/blog /app/
COPY --from=builder /app/source/cmd/start /app/cmd/

# run it
ENTRYPOINT ["/app/cmd/start"]
