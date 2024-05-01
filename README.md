# GoRepo
Lightweight Go server for hosting Maven repositories, setup instructions provided below.

## Table of Contents
- [Requirements](#requirements)
- [Installation](#installation)
- [How to use](#)

## Requirements
- Server (Linux)
- Go

## Installation
```docker
docker build --pull --rm -f "Dockerfile" -t gorepo:latest "."
```

## How to use
```kts
maven("localhost:8080/mvn")
```

```kts
implementation("groupid:artifactId:version)
```

To publish an artifact create a post request to with the artifact included in the body.
```xml
https://localhost:8080/mvn
```

