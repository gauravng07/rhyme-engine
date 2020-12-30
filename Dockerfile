FROM golang:1.15-alpine AS build
RUN apk add --no-cache make git gcc musl-dev
ARG SERVICE

ADD . /go/src/${SERVICE}
WORKDIR /go/src/${SERVICE}

# leverage docker cache and keep the time consuming step together
RUN make clean
RUN make unit-test
RUN make coverage
RUN make ${SERVICE}

RUN mv ${SERVICE} /${SERVICE}
RUN mv assets /assets

FROM alpine:3.9.2

ARG SERVICE

ENV APP=${SERVICE}

RUN apk add --no-cache ca-certificates && \
    mkdir /app && \
    mkdir assets

COPY --from=build /${SERVICE} /app/${SERVICE}
COPY --from=build /assets /app/assets

ENTRYPOINT /app/${APP}
