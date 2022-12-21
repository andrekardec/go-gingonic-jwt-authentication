[![Packagist](https://img.shields.io/packagist/l/doctrine/orm.svg)](https://github.com/andrekardec/go-gingonic-jwt-authentication/blob/main/LICENSE.md)

<a href="https://www.linkedin.com/in/andrekardec/">
   <img src="http://img.shields.io/badge/author-@andrekardec-blue.svg?style=flat-square">
</a>

# Go Gin Gonic JWT Authentication API

[UNDER CONSTRUCTION]

This API provides an Identity Access Management with JWT powered by Gin Gonic.

## Table of Contents

- [Dependencies](#dependencies)
  - [Installing make](#installing-make)
  - [Air - Live reload for Go apps](#air---live-reload-for-go-apps)

* [Running](#running)
  - [Environment Variables](#environment-variables)
  - [Docker](#docker)
  - [Migrations](#migrations)
* [Sonarqube](#sonarqube)
* [API Docs](#api-docs)
* [Testing](#testing)
* [Deploy Lambda Function to AWS](#deploy-lambda-function-to-aws)
  - [GitHub Actions](#github-actions)
  - [Terraform](#terraform)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

### Dependencies

- `go`
- `docker`
- `docker-compose`
- `make`

#### Installing make

Windows

Go to your Power Shell terminal with administrative privileges, then run `choco install make`. You should have make ready to use.

Linux

Make comes in default in most of Linux Distributions, so you should verify if it is already installed before considering installing it. You can verify it typing the below mentioned command in the terminal.

```shell
make -version
```

You can install the make package by typing.

```shell
sudo apt install make
```

MacOS

#### Air - Live reload for Go apps

## Running

TL;DR: To run the application you just need to run the below mentioned command in the root folder of your project. But, to make everything workout in the expected way, make sure you have all the dependencies listed before and that the next steps of the running section were fulfilled.

### Environment Variables

The program looks for the following environment variables:

- `DB_USER`: The postgres database username that gets used within the postgres connection
  string (Default: `root`).
- `DB_PASS`: The postgres database password that gets used within the postgres connection
  string (Default: `root`).
- `DB_NAME`: The postgres database name that gets used within the postgres connection string
  (Default: `user`).
- `DB_HOST`: The postgres database host name that gets used within the postgres connection
  string (Default `db`).
- `DB_PORT`: The postgres database port that gets used within the postgres connection string
  (Default: `5432`).

If the environment variable has a supplied default and none are set within the context of the host
machine, then the default will be used.

To set any given environment variable, simply execute the following
pattern, replacing `[ENV_NAME]` with the name of the environment variable and `[ENV_VALUE]` with the
desired value of the environment variable: `export [ENV_NAME]=[ENV_VALUE]`. To unset any set environment
variable, simply execute the following pattern, replacing `[ENV_NAME]` with the name of the environment
variable: `unset [ENV_NAME]`.

### Docker

This API uses Docker to provider the PostgreSQL database. With you Docker open and running, run the below mentioned command in the terminal and you're ready to go.

```shell
make run
```

### Migrations

## Sonarqube

## API Docs

## Testing

### Unit tests

### E2E tests

## Deploy Lambda Function to AWS

### GitHub Actions

### Terraform
