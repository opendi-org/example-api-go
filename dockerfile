FROM debian:bullseye-slim

WORKDIR /

COPY opendi-api .
COPY opendi-database.db .

EXPOSE 8080

CMD ["./opendi-api"]