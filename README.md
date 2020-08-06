![logo](https://github.com/nightwolf93/brisk/blob/master/logo.png?raw=true)

# Brisk

[![CircleCI](https://circleci.com/gh/nightwolf93/brisk.svg?style=svg)](https://github.com/nightwolf93/brisk)

Brisk is a simple url shortener API.  
It will allow you to create short url for temporary url like for a 2FA system.

## Documentation

[Click here](Chttps://petstore.swagger.io/?url=https://raw.githubusercontent.com/nightwolf93/brisk/master/brisk_openapi.yaml)

## Docker

Brisk have a docker image here : https://hub.docker.com/r/nightwolf931/brisk

#### Environment variables

| Name                 | Description                 |
| -------------------- | --------------------------- |
| MASTER_CLIENT_ID     | ClientID of the master      |
| MASTER_CLIENT_SECRET | Client secret of the master |
| MAX_LINK_TTL         | Max TTL for a link          |
| BASE_URL             | The base URL for links      |

#### Exposed port

The exposed http port by Brisk is the port **3000**

## Authors

Nightwolf93
