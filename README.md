# petstore

[![Build Status](https://cloud.drone.io/api/badges/prologic/petstore/status.svg)](https://cloud.drone.io/prologic/petstore)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/petstore)](https://goreportcard.com/report/github.com/prologic/petstore)
[![Docker Version](https://images.microbadger.com/badges/version/prologic/petstore.svg)](https://microbadger.com/images/prologic/petstore)
[![Image Info](https://images.microbadger.com/badges/image/prologic/petstore.svg)](https://microbadger.com/images/prologic/petstore)

petstore is a simple implementation of a subset of the
[Swagger Petstore](https://petstore.swagger.io/#/pet/getPetById) sample API.

**Note:** Only the "Pet" service and endpoints are fully implemented. This is
intentional to keep scope small and to facilitate a Micro-services design where
each service is responsible for only a subset of Data or Business Domain. It is
fully expected that a similar design and implementation such as the one presented
here is duplicated for other services.

## Quick-start

### From Source

```#!sh
$ go get github.com/prologic/petstore
$ petstore
```

### Using Docker

```#!bash
$ docker run -d -p 8000:8000 prologic/petstore
```

## Demo

There is also a public demo instance available at: https://petstore.mills.io/

## Usage

Run the petstore:

```#!sh
$ petstore
```
Then visit: http://localhost:8000/

## Configuration

By default petstore stores Pet(s) in `petstore.db` in the local directory.

This can be configured with the `-dbpath /path/to/petstore.db` option.

See `petstore --help` for other options.

## License

petstore is licensed under the terms of the [MIT License](https://opensource.org/licenses/MIT)
