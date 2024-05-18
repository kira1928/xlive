FROM alpine

ARG tag

ENV WORKDIR="/srv/xlive"
ENV OUTPUT_DIR="/srv/xlive" \
    CONF_DIR="/etc/xlive" \
    PORT=8080

ENV PUID=0 PGID=0 UMASK=022

RUN mkdir -p $OUTPUT_DIR && \
    mkdir -p $CONF_DIR && \
    apk update && \
    apk --no-cache add ffmpeg libc6-compat curl su-exec tzdata && \
    cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN sh -c "case $(arch) in aarch64) go_arch=arm64 ;; arm*) go_arch=arm ;; i386|i686) go_arch=386 ;; x86_64) go_arch=amd64;; esac && \
    cd /tmp && curl -sSLO https://github.com/kira1928/xlive/releases/download/$tag/xlive-linux-\${go_arch}.tar.gz && \
    tar zxvf xlive-linux-\${go_arch}.tar.gz xlive-linux-\${go_arch} && \
    chmod +x xlive-linux-\${go_arch} && \
    mv ./xlive-linux-\${go_arch} /usr/bin/xlive && \
    rm ./xlive-linux-\${go_arch}.tar.gz" && \
    sh -c "if [ $tag != $(/usr/bin/xlive --version | tr -d '\n') ]; then return 1; fi"

COPY config.docker.yml $CONF_DIR/config.yml

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

VOLUME $OUTPUT_DIR

EXPOSE $PORT

WORKDIR ${WORKDIR}
ENTRYPOINT [ "sh" ]
CMD [ "/entrypoint.sh" ]
