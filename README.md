# Ortelius v11 User Microservice

> Version 11.0.0

RestAPI for the User Object
![Release](https://img.shields.io/github/v/release/ortelius/scec-user?sort=semver)
![license](https://img.shields.io/github/license/ortelius/scec-user)

![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-user/build-push-chart.yml)
[![MegaLinter](https://github.com/ortelius/scec-user/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-user/actions?query=workflow%3AMegaLinter+branch%3Amain)
![CodeQL](https://github.com/ortelius/scec-user/workflows/CodeQL/badge.svg)
[![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-user/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-user)

![Discord](https://img.shields.io/discord/722468819091849316)

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| GET | [/msapi/user](#getmsapiuser) | Get a List of Users |
| POST | [/msapi/user](#postmsapiuser) | Create a User |
| GET | [/msapi/user/:key](#getmsapiuserkey) | Get a User |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |

## Path Details

***

### [GET]/msapi/user

- Summary  
Get a List of Users

- Description  
Get a list ofthe user.

#### Responses

- 200 OK

***

### [POST]/msapi/user

- Summary  
Create a User

- Description  
Create a new User and persist it

#### Responses

- 200 OK

***

### [GET]/msapi/user/:key

- Summary  
Get a User

- Description  
Get a user based on the _key or name.

#### Responses

- 200 OK

## References
