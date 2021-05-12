# AdvenJourney
Connect and travel with like minded people.

Monorepo, contains the code for the AdvenJourney platform:
- [api](https://github.com/advenjourney/advenjourney/tree/main/api) - Golang backend
- [web](https://github.com/advenjourney/advenjourney/tree/main/web) - VuePress frontend

## Create DB for development
Setup DB for development, api uses connection configs in `.env`
```
docker-compose up -d db
```

## Develop
Development for web only
```
make dev
```

Build api only
```
make build
```

Build web and embed in api
```
make    # equal to make assets build
```

Run api backend
```
./build/api
```

## Run 
Build and run server (web + api) and db with docker compose
```
docker-compose up
```

To run using `go run ...`, work directly in `./api`folder
```
cd api
go run cmd/api/api.go 
```