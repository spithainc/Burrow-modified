# Burrow - Kafka Consumer Lag Checking (Forked and modified by spitha)
> You can find the original 'README' at [Burrow](https://github.com/linkedin/Burrow).

## Modifications
* Remove Notifier and Zookeeper - do not need `Notifier` and `Zookeeper` settings
* Fix SASL Bug: set `sasl.handshake-first` to `true`
* Add SASL Mechanisms: PLAIN, OAUTHBEARER, GSSAPI

## Getting Started
### Quick Start
* Our docker image is available through the Docker Hub repository. You can run a burrow in a few simple steps.
* First you need to create a new configuration file 'burrow.yaml', we recommend the yaml format.
* You can run the `Burrow` with the following command.
```
docker run -p 8000:8000 -v ./burrow.yaml:/etc/burrow/burrow.yaml spitharepo/burrow:latest
```
> [Burrow HTTP-Endpoint](https://github.com/linkedin/Burrow/wiki/HTTP-Endpoint)

### Prerequisites
Burrow is written in Go, so before you get started, you should [install and set up Go](https://golang.org/doc/install). As the dependencies
are managed using Go module, the lowest version of Go supported is 1.11, though we recommend using version 1.12 for development.

### Build and Install
```
$ Clone github.com/linkedin/Burrow to a directory outside of $GOPATH. Alternatively, you can export GO111MODULE=on to enable Go module.
$ cd to the source directory.
$ go mod tidy
$ go install
```

### Running Burrow
```
$ $GOPATH/bin/Burrow --config-dir /path/containing/config
```

### Using Docker
A Docker file is available which builds this project on top of an Alpine Linux image.
To use it, build your docker container, mount your Burrow configuration into `/etc/burrow` and run docker.

A [Docker Compose](docker-compose.yml) is also available for quick and easy development.

Install [Docker Compose](https://docs.docker.com/compose/) and then:

1. Build the docker container:
   ```
   docker-compose build
   ```

2. Run the docker compose stack which includes kafka and zookeeper:
   ```
   docker-compose down; docker-compose up
   ```

3. Some test topics have already been created by default and Burrow can be accessed on `http://localhost:8000/v3/kafka`.


### Configuration
For information on how to write your configuration file, check out the [detailed wiki](https://github.com/linkedin/Burrow/wiki)

### Additional configuration examples for the newly added SASL mechanism in yaml format
- PLAIN
```
sasl:
  my-kafka:
    mechanism: PLAIN
    username: client
    password: client-secret
```

- SCRAM
```
  sasl:
    my-kafka:
      mechanism: SCRAM-SHA-512
      username: client
      password: client-secret
```

- OAUTHBEARER
```
sasl:
  my-kafka:
    mechanism: OAUTHBEARER
    clientId: client
    clientSecret: client-secret
    tokenEndpoint:  http://.../realms/my-kafka/protocol/openid-connect/token
```

- GSSAPI
```
sasl:
  my-kafka:
    mechanism: GSSAPI
    servicename: kafka
    realm: mydomain.com
    username: client
    keytabPath: keytab/client.keytab
    kerberosConfigPath: /etc/krb5.conf
    disablePAFXFAST: fasle
```

## License
Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied.
