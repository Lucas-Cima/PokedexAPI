FROM golang:1.12.0-alpine3.9
RUN mkdir /pokedexAPI
ADD . /pokedexAPI
WORKDIR /pokedexAPI
RUN go build -o main .
CMD ["/pokedexAPI/main"]