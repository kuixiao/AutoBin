FROM alpine

RUN mkdir -p /go/src/autobin/user/build
WORKDIR /go/src/autobin/user/build
ADD build/autobin-user .
ADD conf/config.json .

EXPOSE 8182
ENTRYPOINT "/go/src/autobin/user/build/autobin-user"
