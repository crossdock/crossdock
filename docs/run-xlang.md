[← Write Test Client](write-test-client.md)

# Run Xlang

To run xlang, we'll need to configure our `docker-compose.yml`:

```yml
xlang:
    image: yarpc/xlang
    links:
        - client
    environment:
        - XLANG_CLIENTS=client
        - XLANG_DIMENSION_BEHAVIOR=dance,run

client:
    build: .
    ports:
        - 8080
```

In the configuration above, the `xlang:` entry can be read as:

> Define a container named xlang that runs the [yarpc/xlang](https://hub.docker.com/r/yarpc/xlang/) image,
> assigns our Test Client as a runtime dependency,
> and defines a custom Dimension "behavior".

And the `client:` entry can be read as:

> Define a container named client that is created by running `docker build`
> the `Dockerfile` located in the current directory, then open port 8080.

Of course, we'll need to define a `Dockerfile` in order to build our Test Client:

```Dockerfile
FROM golang:onbuild
EXPOSE 8080
```

Finally, we can call xlang:

```
$ docker-compose run xlang

Beginning matrix of tests...

  STATUS  | CLIENT |      RESPONSE       | BEHAVIOR
+---------+--------+---------------------+----------+
  PASSED  | client | ok                  | dance
  SKIPPED | client | 404                 | run

```

The above output can be read as:

> For every Test Client configured in our `docker-compose.yml` file,
> xlang issued a test request to each for every Behavior defined.

[Publish Test Client →](publish-test-client.md)
