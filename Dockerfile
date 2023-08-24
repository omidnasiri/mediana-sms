FROM golang:1.20 as build
RUN mkdir -p /app
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /app/bin/mediana-sms /app/cmd/

FROM alpine:3.16 as run
# TODO: Add envs for other that dev environments
COPY --from=build /app/bin/mediana-sms /bin/mediana-sms

EXPOSE 8080
CMD [ "/bin/mediana-sms"]
