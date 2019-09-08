#!/bin/bash

BUILD_VERSION="$(date +'%Y%m%d%H%M%S')-$(git log --format=%h -1)"

echo "Building binary with version [${BUILD_VERSION}]..."

BINARY_NAME="modem-scraper"
FLAGS="-X main.BuildVersion=${BUILD_VERSION}"
go build -ldflags="${FLAGS}" -o ${BINARY_NAME}
echo "Built version:"
./modem-scraper -version

ARCHIVE_NAME="${BINARY_NAME}_${BUILD_VERSION}_linux_amd64.zip"
zip ${ARCHIVE_NAME} ${BINARY_NAME}

if [ "${TRAVIS_BRANCH}" = "master" ] && [ "${TRAVIS_PULL_REQUEST_BRANCH}" = "" ] ; then
  echo "Detected master branch; pushing version tag to repository..."

  git config --local user.name "${GIT_USER_NAME}"
  git config --local user.email "${GIT_USER_EMAIL}"
  git tag ${BUILD_VERSION}
  git remote add tag-origin https://${GITHUB_TOKEN}@github.com/pdunnavant/modem-scraper.git
  git push tag-origin ${BUILD_VERSION}
fi
