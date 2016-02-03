# Xlang [![Build Status](https://travis-ci.org/yarpc/xlang.svg?branch=master)](https://travis-ci.org/yarpc/xlang)

A 2mb Docker appliance for running cross-repo integration tests; Xlang is:

* Portable - runs anywhere Docker is installed, eg Travis & locally.
* General - can be used to test sets of libraries and microservices.
* Flexible - test all combinations of behaviors using custom matrix dimensions.
* Decentralized - each repo can configure and run Xlang independently from the others.
* Light - run Xlang for every commit on every repo in parallel.
* Easy - run integration tests on a large project without installing every component.

## How It Works

Xlang is [published in Docker Hub](https://hub.docker.com/r/yarpc/xlang/) and is
meant to be used with [Docker Compose](https://docs.docker.com/compose/) directly from your repos.

Given the following `docker-compose.yml`, Xlang will initiate an integration test for clients
`alpha` and `omega` for every combination of `behavior` and `speed`:

```yml
xlang:
    image: yarpc/xlang
    links:
        - alpha
        - omega
    environment:
        - XLANG_CLIENTS=alpha,omega
        - XLANG_DIMENSION_BEHAVIOR=dance,run
        - XLANG_DIMENSION_SPEED=fast,slow

alpha:
    image: breerly/hello-server
    ports:
        - 8080
    environment:
        - HELLO_PORT=8080
        - HELLO_MESSAGE=ok

omega:
    image: breerly/hello-server
    ports:
        - 8080
    environment:
        - HELLO_PORT=8080
        - HELLO_MESSAGE=ok
```

Run Xlang:

```
$ docker-compose run xlang

Beginning matrix of tests...

  STATUS | CLIENT | RESPONSE | SPEED | BEHAVIOR
+--------+--------+----------+-------+----------+
  PASSED | alpha  | ok       | fast  | dance
  PASSED | alpha  | ok       | fast  | run
  PASSED | alpha  | ok       | slow  | dance
  PASSED | alpha  | ok       | slow  | run
  PASSED | omega  | ok       | fast  | dance
  PASSED | omega  | ok       | fast  | run
  PASSED | omega  | ok       | slow  | dance
  PASSED | omega  | ok       | slow  | run
```

## How To Use

1. [Write Test Client](docs/write-test-client.md)
2. [Run Xlang](docs/run-xlang.md)
3. [Publish Test Client](docs/publish-test-client.md)
4. [Integrate Other Repos](docs/integrate-other-repos.md)
5. [Add Other Test Dimensions](docs/add-other-dimensions.md)
6. [Run During Continuous Integration](docs/add-to-ci.md)

