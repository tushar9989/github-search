#!/bin/bash
docker build -t githubsearch .
docker run -it --name githubsearch --rm githubsearch