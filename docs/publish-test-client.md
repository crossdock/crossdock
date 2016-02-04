[← Run Crossdock](run-crossdock.md)

# Publish the Test Client

Now that we have a Test Client working locally, we need to publish it
so that it's available for other repos to run.

We suggest pushing the Test Client to the [Docker Hub](https://hub.docker.com/)
every time your build succeeds. If you're using Travis, the following guide will assist
in pushing a container for every git commit that passes your tests:

[Using Travis.ci to build Docker images](https://sebest.github.io/post/using-travis-ci-to-build-docker-images/)

Don't forget to Enable Docker Support in your `.travis.yml`:

```yml
sudo: required
language: go
go: 1.5
services: docker
after_success:
    - export REPO=myorg/client
    # ...
```

You can test that the container was published by pulling it down locally like so:

```
$ docker pull myorg/client
Using default tag: latest
latest: Pulling from myorg/client

2704b7bed4fd: Pull complete
5baf16abbb79: Pull complete
Digest: sha256:6925c040d80bd102a8c8d3de1031641384cc372c303a6be04a8e67ae932b6e82
Status: Downloaded newer image for myorg/client:latest
```

[Integrate Other Repos →](integrate-other-repos.md)
