FROM node:lts-alpine3.18

ARG GIT_COMMIT
ENV PORT=8181

RUN apk update && \
    apk --no-cache add curl unzip

RUN echo "Based on commit: $GIT_COMMIT"
RUN sh -c "curl -sSLO https://github.com/kira1928/badanmu/archive/$GIT_COMMIT.zip && \
    unzip $GIT_COMMIT.zip && \
    mv badanmu-$GIT_COMMIT /badanmu && \
    cd /badanmu/packages/badanmu && \
    npm install"

WORKDIR "/badanmu/packages/badanmu"

EXPOSE $PORT

ENTRYPOINT ["npm"]
CMD ["run", "start", "$PORT"]