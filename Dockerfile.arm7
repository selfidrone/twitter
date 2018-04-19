FROM alpine

RUN adduser -h /home/tweets -D tweets tweets

COPY ./tweets /home/tweets/
RUN chmod +x /home/tweets/tweets

USER tweets

ENTRYPOINT ["/home/tweets/tweets"]
