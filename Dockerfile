FROM alpine
ADD foobar-service /foobar-service
ENTRYPOINT [ "/foobar-service" ]
