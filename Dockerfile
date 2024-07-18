FROM harbor.axzo.cn/library/golang:1.20.4-alpine AS builder
WORKDIR /gin-web-init
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /gin-web-init
RUN go env -w GOPROXY="https://goproxy.cn" && \
    go mod download && \
    go build -o gin-web-init_linux main.go


FROM harbor.axzo.cn/kube-ops/ubuntu18-sonar:v4
WORKDIR /gin-web-init
# exporter
COPY --from=builder /gin-web-init/gin-web-init_linux /usr/bin/gin-web-init_linux
COPY ./config.toml /gin-web-init/config.toml
ENV CONFIG="/gin-web-init/config.toml"
RUN chmod a+x /usr/bin/gin-web-init_linux && mkdir /resource

ENTRYPOINT ["/usr/bin/gin-web-init_linux"]

EXPOSE 8080
