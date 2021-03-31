FROM golang:1.16.2-alpine3.13 AS build

RUN apk add --no-cache gcc musl-dev

RUN mkdir /app

ENV GO111MODULE=on
ENV GOPATH=/app

RUN mkdir -p $GOPATH/src/github.com/lffranca/api
RUN mkdir $GOPATH/bin
RUN mkdir $GOPATH/pkg

WORKDIR $GOPATH/src/github.com/lffranca/api
COPY . $GOPATH/src/github.com/lffranca/api


WORKDIR $GOPATH/src/github.com/lffranca/api
RUN go build -v -o $GOPATH/bin/app -ldflags '-extldflags "-static"' $GOPATH/src/github.com/lffranca/api/cmd/api/main.go
WORKDIR $GOPATH

RUN mkdir /build

RUN cp $GOPATH/bin/app /build

VOLUME $GOPATH/bin

# wkhtmltopdf
FROM surnet/alpine-wkhtmltopdf:3.8-0.12.5-full as wkhtmltopdf

# ALPINE
FROM alpine:3.13.3

RUN  echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.8/main" > /etc/apk/repositories \
     && echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.8/community" >> /etc/apk/repositories \
     && apk update && apk add --no-cache \
      libstdc++ \
      libx11 \
      libxrender \
      libxext \
      libssl1.0 \
      ca-certificates \
      fontconfig \
      freetype \
      ttf-dejavu \
      ttf-droid \
      ttf-freefont \
      ttf-liberation \
      ttf-ubuntu-font-family \
    && apk add --no-cache --virtual .build-deps \
      msttcorefonts-installer \
    \
    # Install microsoft fonts
    && update-ms-fonts \
    && fc-cache -f \
    \
    # Clean up when done
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/* \
    && apk del .build-deps

RUN apk add --no-cache libc6-compat
RUN apk add --no-cache ca-certificates
# RUN apk add --no-cache wkhtmltopdf

COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=wkhtmltopdf /bin/wkhtmltoimage /bin/wkhtmltoimage

RUN mkdir /app
COPY --from=build /app/bin/app /app/app
# COPY ./production.env /app/.env
RUN chmod +x /app/app

WORKDIR /app

EXPOSE 8000

CMD [ "./app" ]
