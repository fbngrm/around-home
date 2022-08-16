# Overview

The solution implements all requirements from the task definition.

Since not further specified, I chose to implement a `gRPC API`.
The gRPC API naturally uses `protocol buffers` and `HTTP/2`.
However, an `HTTP API` that supports JSON data via `HTTP/1.1` is also provided by mapping the proto definition to the `OpenAPI` standard.
Though, the HTTP API's only purpose is making interaction with the API easier for exploring the endpoints.
It does not implement a real REST API.
Mapping a RPC based API to HTTP needs to be given more thought and I consider it out of scope of this challenge.

The API is described in a [.proto file](./apis/match/v1/match.proto).
The JSON API is documented at `http://localhost:8081/openapi/`.
Enter `/openapi.json` in the box at the top of the page to browse the endpoint definitions.

The project utilizes code generation wherever possible.
The code generation process generates the gRPC and HTTP APIs, the gRPC client code, an OpenAPI specification, the database access layer and migrations as well as mocks for unit tests.

The project provides `Dockerfiles` and a `docker-compose` configuration for local development/testing.
There are 4 containers used, a postgres database, a migration container, a seed container and the main app.
Please note, building the Dockerfiles results in large images.
Due to time constraints, I didn't optimize towards smaller images.

The `just` command runner is used to automate command execution.
A `Makefile` is provided for making build dependencies to avoid re-running time-consuming code generation steps.

## Requirements

To run the application it is required that `Docker` and `docker-compose` are installed on the machine.
Please note, I added the `/gen` and `/ent` directories since they are required to build the docker images.
In a normal setup this would not be shipped into version control but generated.
I added them anyway to keep the requirements and installations to the reviewer's machine as limited as possible.
It is also required to install `just` or run the command in the justfile manually.
A tool for making HTTP requests is required as well, e.g. `cURL` or `wget`.

## Running the app

The easiest way to run the app is to start it via docker-compose and access the HTTP API via cURL.

```
just seed-and-run
```

```
curl -X 'GET' \
  'http://localhost:8081/v1/partners/Wood/52.532566:13.396261' \
  -H 'accept: application/json'
```

This should result in a response like:

```
{
    "partner": [
        {
            "distance": 0,
            "id": 3,
            "location": "52.532566:13.396261",
            "materials": "WOOD, TILES",
            "radius": 100,
            "rating": 5
        },
        {
            "distance": 1643.552183,
            "id": 1,
            "location": "52.519065:13.406083",
            "materials": "WOOD, TILES, CARPET",
            "radius": 10000,
            "rating": 5
        },
        {
            "distance": 1643.552183,
            "id": 2,
            "location": "52.519065:13.406083",
            "materials": "WOOD, TILES, CARPET",
            "radius": 10000,
            "rating": 4
        }
    ]
}
```

## Configuration

There is very minimal config supported currently.
I would normally use `env` variables and flags for configuration according to `12 Factor App` methodology.

## Architecture

The project uses `Clean Architecture` throughout the codebase.
An exception is the database schemas, which is left out due to time constraints.
Types are designed around domain capabilities.
Dependency injection is used to achieve loose coupling, extensibility and testability.

## Tests

I added only one unit test due to time constraints.
It shows table-driven tests for the api on package boundary.
The application would need higher test coverage before going to production, especially the calculations, matching and sorting.
There are no integration tests due to time constraints.

## Distance calculations

I used the `Haversine` algorithm for calculating the distance.
I copied the code and marked this in the header of the file.
Due to time constraints, I didn't implement it myself and didn't test it thoroughly, which would be a requirement for production.

## Error handling

The error handling in the API endpoint is overly simplified due to time constraints.
This would be one of the first things that needed to be addressed before production since we must not expose errors to the outside user.

## Next steps

A few things that I would implement as the next steps are:
- Restructure ent directory to adhere to Clean Architecture
- Config
- Error handling in the endpoints
- Integration tests
- Logging
- Authentication
- Customer management
- Storing of request, associated with customers
- Health/Readiness endpoints
- Observability (metrics, tracing)
- TLS support

## Reasoning

I would be happy to explain all reasoning of design decisions and shortcomings in depth in a potential next interview :)
