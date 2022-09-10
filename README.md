# Flight tracker
The microservice API accepts a request that includes a list of flights and returns the total flight path with origin and destination airports. 
The list of flights defined by a source and destination airport code and may not be listed in order.

Examples:
```
[["SFO", "EWR"]] => ["SFO", "EWR"]
[["ATL", "EWR"], ["SFO", "ATL"]] => ["SFO", "EWR"]
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
```
## Get started
### Prerequisites
- Makefile
- Docker

### Run with docker:
- go to the project root
- `make docker-build`
- `make docker-run`

### Send request using `curl`
```
curl http://localhost:8080/get-flight-path -d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'
```
The response should be next:
``` 
["SFO","EWR"]
```

## For developers
### Project structure
I used clean architecture approach and below you can find actual project structure:
```
internal: is for modules that is not likely to be used in other projects
└── service: wrapper around modules following Clean Architecture approach
    └── transport - transport layer
    └── usecase - business logic layer
    └── repo - repository layer(currently is absent)
pkg: is for modules that can be used in other projects, in future we can move them to separate library
```

### Testing
For unit tests run next command:
```
make test-unit
```

For integration tests run next command:
```
make test-integration
```

### Linter
To run linter run `make lint` command

We use golangci linter, that aggregates all available linters.
Linter's configuration is located in `.golangci.yml`.

### Env variables
You can set your log leve and port(don't forget to change docker run command as well) as environment variables and service will use them.

