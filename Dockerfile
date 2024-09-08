ARG DOCKER_BUILD_IMAGE=golang:1.22
ARG DOCKER_BASE_IMAGE=alpine:3.18

FROM ${DOCKER_BUILD_IMAGE} AS build

RUN apt-get update -yq

WORKDIR /github.com/rakib-09/golang-clean-architecture

COPY ./ ./

RUN make build

FROM ${DOCKER_BASE_IMAGE} AS final

RUN apk update && apk add ca-certificates
RUN apk add --no-cache tzdata

COPY --from=build /github.com/rakib-09/golang-clean-architecture/binary/app /bin/app

WORKDIR /github.com/rakib-09/golang-clean-architecture

EXPOSE 8080
ENTRYPOINT ["app"]