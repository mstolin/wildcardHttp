FROM golang:1.23.4-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o /bin/whttp ./main.go

FROM scratch
LABEL maintainer="Marcel Stolin <marcelstolin@gmail.com>"
COPY --from=build /bin/whttp /bin/whttp
EXPOSE 80/tcp
ENTRYPOINT ["/bin/whttp"]
CMD [":80"]