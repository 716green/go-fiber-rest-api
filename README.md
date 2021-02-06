# go-fiber-rest-api

## Go + Go Fiber + SQLite

### REST API for technologies database

#### Run Server

```bash
go run main.go
```

## Endpoints

#### Get all technologies

```bash
curl http://localhost:3000/api/v1/tech
```

#### Get single technology by ID

```bash
curl http://localhost:3000/api/v1/tech/1
```

#### Add Single Technology

```bash
curl -X POST -H "Content-Type: application/json" --data "{\"technology\":\"Go\", \"category\":\"System Level Language\", \"note\": \"Server Programming Language by Google\"}" http://localhost:3000/api/v1/tech
```

#### Delete single technology by ID

```bash
curl -X DELETE  http://localhost:3000/api/v1/tech/1
```
