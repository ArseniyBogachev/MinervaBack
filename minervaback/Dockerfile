FROM golang:latest AS builder

COPY ./deploy-keys/ /root/.ssh/
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host gitea.unprism.ru\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url."ssh://git@gitea.unprism.ru:8022".insteadOf "https://gitea.unprism.ru"
ENV GOPRIVATE gitea.unprism.ru/KRBL/FemaProtocol

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY ./go.* ./

RUN go mod download

COPY ./ ./
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

FROM alpine:latest AS runner

COPY --from=builder build /bin/build

EXPOSE 8080

ENTRYPOINT [ "/bin/build" ]
