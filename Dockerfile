FROM alpine:3.10

RUN mkdir /app
WORKDIR /app
ADD ./.out/auth-service /app/auth-service

CMD ["./auth-service"]
