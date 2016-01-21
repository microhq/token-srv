FROM alpine:3.2
ADD token-srv /token-srv
ENTRYPOINT [ "/token-srv" ]
