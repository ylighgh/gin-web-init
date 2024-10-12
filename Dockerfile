FROM ubuntu
WORKDIR /gin-web-init
# exporter
COPY gin-web-init_linux /usr/bin/gin-web-init_linux
COPY ./config.toml /gin-web-init/config.toml
ENV CONFIG="/gin-web-init/config.toml"
RUN chmod a+x /usr/bin/gin-web-init_linux && mkdir /resource

ENTRYPOINT ["/usr/bin/gin-web-init_linux"]

EXPOSE 8080
