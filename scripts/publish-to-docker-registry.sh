#!/bin/bash

echo "$TRAVIS_BRANCH"
echo "$TRAVIS_BUILD_DIR"

TAG=$(if [ "$TRAVIS_BRANCH" == "master" ]; then echo "latest"; else echo "$TRAVIS_BRANCH"; fi)

echo "tag is $TAG"

# docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASS

