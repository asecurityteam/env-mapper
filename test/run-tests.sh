#!/bin/sh
docker build -t local/test/env-mapper .
docker build -t local/test/env-mapper-tests test
docker run -i local/test/env-mapper-tests