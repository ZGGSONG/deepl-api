FROM alpine:latest

ARG TARGETARCH

RUN mkdir "/app"
WORKDIR "/app"

# 修正时区
RUN apk update && apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone


COPY app* "/app/."

RUN mv app_${TARGETARCH} app

EXPOSE 8000

ENTRYPOINT ["./app"]