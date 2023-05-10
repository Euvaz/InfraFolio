# InfraFolio

This portfolio website project is a Go-based implementation utilizing the [gin](https://github.com/gin-gonic/gin) library as the web framework, with built-in HTML templating support via the [multitemplate](https://github.com/gin-contrib/multitemplate) library for efficient and reliable website rendering.

To maintain the website's currency, the project incorporates a [Python script](https://github.com/Euvaz/InfraFolio/blob/main/scripts/scrape.py), which interfaces with GitHub's API to retrieve updated data on my project repositories and is executed automatically through a Cron Job via the provided [crontab](https://github.com/Euvaz/InfraFolio/blob/main/cron/crontab.txt) file. This data is stored in an SQLite database and automatically propagated to the website's pages.

This project is optimized for deployment as a microservice using a Docker images via the provided [Dockerfile](https://github.com/Euvaz/InfraFolio/blob/main/Dockerfile). The images are subjected to robust linting, testing, and building processes through GitHub Actions, ensuring the highest quality and reliability.

With its technical implementation and advanced functionalities, this project offers an effective way to exhibit my web development skills and showcase my professional experience, while ensuring visitors are always up-to-date on the my latest endeavors.

## Overview

- [Repository structure](https://github.com/euvaz/infrafolio#-repository-structure)
- [Usage](https://github.com/euvaz/infrafolio#usage)
  - [Docker (Recommended)](https://github.com/euvaz/infrafolio#docker-recommended)
    - [Build](https://github.com/euvaz/infrafolio#build)
    - [Run](https://github.com/euvaz/infrafolio#run)
  - [Manual](https://github.com/euvaz/infrafolio#manual)
    - [Verify Prereqs](https://github.com/euvaz/infrafolio#verify-prereqs)
    - [Download Dependencies](https://github.com/euvaz/infrafolio#download-dependencies)
    - [Compile Binary](https://github.com/euvaz/infrafolio#compile-binary)


## 📂 Repository structure

```
cron        # Crontab files
scripts     # Python helper scripts
website/    # Website files
├── static  # Static website resources
└── tmpl    # HTML templates
```

## Usage

### Docker (Recommended)

A [Dockerfile](https://github.com/Euvaz/InfraFolio/blob/main/Dockerfile) is provided and kept regularly updated for convenient deployment.

#### Build

Firstly, the Docker image will need to be built:

```
$ docker build -t infrafolio .
```

#### Run

Afterwards, the Docker image can be ran:

```
$ doker run -itdp 8080:8080 infrafolio:latest
```

### Manual

The service can be ran manually after downloading dependencies and compiling the binary. Note that this will require a Systemd/OpenRC unit file be configured manually, as this is intended to be ran as a microservice.

Prereqs:
    - [Go 1.19]()
    - [Python3]()

#### Verify Prereqs

Ensure [Go](https://go.dev/) and [Python](https://www.python.org/) are installed:

```
$ go version
go version go1.19.8 linux/amd64
$ python --version
Python 3.10.11
```

#### Download Dependencies

Go:

```
$ pip install -r scripts/requirements.txt
```

Python:

```
$ go mod downlod
```

#### Compile Binary

```
$ go build
```

