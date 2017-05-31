# ether-waveform

Gin Tonic example project to test mocking libraries and an alternative to the existing
REST framework for a oriented microservice architecture.

Encourage to do:

- Dependency injection
- Mocking libraries
- Build cycle with Docker which makes host OS environment agnostic.
- Unit testing following TDD principles
- SRP classes

# TODOs

[ X ] Simple health check with debug

[ X ] Make it configurable with *YML*

[ X ] Dockerize the service (User docker-compose & Healthchecks) (build, test, deploy)

[ ? ] Deploy documentation with Docker

- Having problems deploying it as webserver

[ X ] Mocking objects

- [Mockery](https://github.com/vektra/mockery])
- [Testity](https://github.com/stretchr/testify)

[ X ] Integration tests with docker

- [Ginkgo]((Ginko)[https://onsi.github.io/ginkgo/])

[ ] Expose Swagger spec with docker and go

[ ] Concourse pipeline

[ ] Openshift cluster with the Application

[ ] Prometheus metrics within the microservice

# Docker Cycle

- Build

```bash
docker-compose -f docker/build up
```

```bash
go build pkg/api/main/main.go
```

- Unit Tests

```bash
docker-compose -f docker/tests.yml up
```

```bash
go test -v $(glide novendor)
```

- Integration Tests

```bash
docker-compose -f docker/integration.yml up
```

```bash
go test -v $(glide novendor) --tags=integration
```

### Tooling

Go Tools specified within the [awesome-go](https://awesome-go.com/) framework which help to integrate
different parts of the dev cycle.

## [Ginko](https://onsi.github.io/ginkgo/)

- Generate tests

```bash
ginkgo bootstrap
```

```bash
ginkgo generate controller
```

## [Mockery](https://github.com/vektra/mockery) 

- Generate mocks (The mock targets should be interfaces, so encourage to enforce interfacing
classes)

```bash
mockery -name=Stringer
```

