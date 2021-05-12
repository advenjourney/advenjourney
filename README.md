# AdvenJourney
Connect and travel with like minded people.

Monorepo, contains the code for the AdvenJourney platform:
- [api](https://github.com/advenjourney/advenjourney/tree/main/api) - Golang backend
- [web](https://github.com/advenjourney/advenjourney/tree/main/web) - VuePress frontend

## Develop

Build and run server (web + api) and db with docker compose
```
docker-compose up
```

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

