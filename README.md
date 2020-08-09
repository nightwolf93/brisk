![logo](https://github.com/nightwolf93/brisk/blob/master/logo.png?raw=true)

# Brisk

[![CircleCI](https://circleci.com/gh/nightwolf93/brisk.svg?style=svg)](https://github.com/nightwolf93/brisk)

Brisk is a simple url shortener API.  
It will allow you to create short url for temporary url like for a 2FA system.  
It also provide geoip detection for analytics

## Documentation

https://nico-style931.gitbook.io/brisk/

## 3rd Party packages

| Package          | URL                                        |
| ---------------- | ------------------------------------------ |
| BriskDotNet (C#) | https://github.com/nightwolf93/BriskDotNet |
| php-brisk (PHP)  | https://github.com/nightwolf93/php-brisk   |

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

## Build

```
make build
```

```
make docker
```

## Test

```
make test
```

## Authors

Nightwolf93
