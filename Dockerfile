FROM alpine

RUN apk --update upgrade && \
    apk add ca-certificates && \
    update-ca-certificates

RUN adduser -h /home/tweets -D tweets tweets

COPY ./tweets /home/tweets/
RUN chmod +x /home/tweets/tweets

USER tweets

ENTRYPOINT ["/home/tweets/tweets"]
