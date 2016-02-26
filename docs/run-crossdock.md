[← Write Test Client](write-test-client.md)

# Run Crossdock

To run Crossdock, we'll need to configure our `docker-compose.yml`:

```yml
crossdock:
    image: yarpc/crossdock
    links:
        - client
    environment:
        - CROSSDOCK_CLIENTS=client
        - CROSSDOCK_AXIS_BEHAVIOR=dance,run

client:
    build: .
    ports:
        - 8080
```

In the configuration above, the `crossdock:` entry can be read as:

> Define a container named crossdock that runs the [yarpc/crossdock](https://hub.docker.com/r/yarpc/crossdock/) image,
> assigns our Test Client as a runtime dependency,
> and defines a custom Axis "behavior".

And the `client:` entry can be read as:

> Define a container named client that is created by running `docker build`
> the `Dockerfile` located in the current directory, then open port 8080.

Of course, we'll need to define a `Dockerfile` in order to build our Test Client:

```Dockerfile
FROM golang:onbuild
EXPOSE 8080
```

Finally, we can call Crossdock:

```
$ docker-compose run crossdock

Beginning matrix of tests...

  STATUS  | CLIENT |      RESPONSE       | BEHAVIOR
+---------+--------+---------------------+----------+
  PASSED  | client | ok                  | dance
  SKIPPED | client | 404                 | run

```

The above output can be read as:

> For every Test Client configured in our `docker-compose.yml` file,
> Crossdock issued a test request to each for every Behavior defined.

[Publish Test Client →](publish-test-client.md)