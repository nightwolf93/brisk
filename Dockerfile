FROM golang:latest
LABEL maintainer="Nightwolf93 <nightwolf93@protonmail.com>"
WORKDIR /app
COPY ./bin ./
EXPOSE 3000
CMD ["./brisk"]