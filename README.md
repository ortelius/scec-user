# Ortelius v11 User Microservice
RestAPI for the User Object
![Release](https://img.shields.io/github/v/release/ortelius/scec-user?sort=semver)
![license](https://img.shields.io/github/license/ortelius/scec-user)

![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-user/build-push-chart.yml)
[![MegaLinter](https://github.com/ortelius/scec-user/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-user/actions?query=workflow%3AMegaLinter+branch%3Amain)
![CodeQL](https://github.com/ortelius/scec-user/workflows/CodeQL/badge.svg)
[![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-user/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-user)

![Discord](https://img.shields.io/discord/722468819091849316)

## Version: 11.0.0

### Terms of service
<http://swagger.io/terms/>

**Contact information:**
Ortelius Google Group
ortelius-dev@googlegroups.com

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

---
### /msapi/user

#### GET
##### Summary

Get a List of Users

##### Description

Get a list ofthe user.

##### Responses

| Code | Description |
|------|-------------|
| 200  | OK          |

#### POST
##### Summary

Create a User

##### Description

Create a new User and persist it

##### Responses

| Code | Description |
|------|-------------|
| 200  | OK          |

### /msapi/user/:key

#### GET
##### Summary

Get a User

##### Description

Get a user based on the _key or name.

##### Responses

| Code | Description |
|------|-------------|
| 200  | OK          |
