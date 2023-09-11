<section align="center">

# Log Auth

[![Go Report Card](https://goreportcard.com/badge/github.com/MarcusXavierr/log-auth)](https://goreportcard.com/report/github.com/MarcusXavierr/log-auth)
![example workflow](https://github.com/MarcusXavierr/log-auth/actions/workflows/go.yml/badge.svg)

[Installation](#installation) •
[Getting Startd](#getting-started) •
[Usage](#usage)

</section>

## Installation

### go install
you can install this project with go
```bash
go install github.com/MarcusXavierr/log-auth@latest
```
### Binaries
You can also install a compiled binary to your machine and then put it in your shell path.

Go to the [releases page](https://github.com/MarcusXavierr/log-auth/releases) and choose the option that best fits your environment.

## Getting Started
To get started, just put this configuration file into your home folder, on a file name `.log-auth.yml`

### Configuration
This is an example config file, name `.log-auth.yml`
```yml
email: "your@email.com"
password: "your_password"

# default values, just change if you know what you're doing
#
# request_method: "POST" # the http method used
# endpoint: "http://api-auth.homol.logcomex.io/api/login" # the endpoint to send the login request
# shipment_intel_dataset_id: "62c3d669-c1ed-499e-bcd5-32ff1108a814" # The dataset ID represents the dataset used (each PBI product have one dataset_id)
# freight_intel_dataset_id: "4765cb7a-b9b3-4f04-a378-a34045e31836" # same here
```

## Usage

### Get a link to log in on shipment intel
You can easilly get an link to log in on shipment intel platform
```bash
log-auth shipment
```
### Get a link to log in on freight intel
You can easilly get an link to log in on shipment intel platform
```bash
log-auth freight
```
