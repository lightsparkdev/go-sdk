FROM --platform=$BUILDPLATFORM golang:1.22-bookworm as builder

ARG TARGETOS TARGETARCH
RUN echo "$TARGETARCH" | sed 's,arm,aarch,;s,amd,x86_,' > /tmp/arch
RUN apt-get update && apt-get install -y "gcc-$(tr _ - < /tmp/arch)-linux-gnu" && apt-get clean && rm -rf /var/lib/apt/lists/*

ENV GOOS $TARGETOS
ENV GOARCH $TARGETARCH

COPY . /src
RUN go env
RUN cd /src/examples/uma-server && CGO_ENABLED=1 CC=$(cat /tmp/arch)-linux-gnu-gcc go install
RUN if [ -e /go/bin/${TARGETOS}_${TARGETARCH} ]; then mv /go/bin/${TARGETOS}_${TARGETARCH}/* /go/bin/; fi

FROM debian:bookworm as final

RUN addgroup --system --gid 1000 go && adduser --system --uid 1000 --ingroup go go
RUN apt-get update && apt-get -y install ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists

COPY --from=builder /go/bin/uma-server /usr/local/bin
COPY --from=builder /src/examples/uma-server/templates /templates

ENTRYPOINT ["uma-server"]
