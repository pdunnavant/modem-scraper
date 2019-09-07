#!/bin/bash

BUILD_VERSION="1.0.0"
FLAGS="-X main.BuildVersion=${BUILD_VERSION}"
go build -ldflags="${FLAGS}"
./modem-scraper -version
