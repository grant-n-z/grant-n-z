# Authentication and Authorization micro service
[![Build Status](http://www.concourse.developer-tm.com:8080/api/v1/teams/main/pipelines/grant-n-z-pipeline/jobs/test/badge)](https://www.concourse.developer-tm.com/teams/main/pipelines/springboot-bestpractice-pipeline)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/tomoyane/grant-n-z/blob/master/LICENSE.txt)

----

Grant-n-z is simple and fast authentication and authorization server.
When i develop some micro service application and some micro service tool, i need to shared authentication system.
Then i can use grant-n-z application.

If you can use container infrastructure, you can use docker image.

----

## To start using grant-n-z

### Environment variable
```bash
GRANT_N_Z_PRIVATE_KEY
GRANT_N_Z_MYSQL_HOST
GRANT_N_Z_MYSQL_USER
GRANT_N_Z_MYSQL_PASS
GRANT_N_Z_MYSQL_PORT
GRANT_N_Z_MYSQL_DB
```

* GRANT_N_Z_PRIVATE_KEY
  * This is base64 private key for signed token. 

* GRANT_N_Z_MYSQL_HOST
  * This is grant-n-z database host on mysql.
  
* GRANT_N_Z_MYSQL_USER
  * This is grant-n-z database user on mysql.

* GRANT_N_Z_MYSQL_PASS
  * This is grant-n-z database pass on mysql.

* GRANT_N_Z_MYSQL_PORT
  * This is grant-n-z database port on mysql.

* GRANT_N_Z_MYSQL_DB
  * This is `grant-n-z` db on mysql.

### Run on Container
Base docker image
```bash
$ docker pull tomohito/grant-n-z
```

If you run Container Orchestration tool, you can write base image this one.
```dockerfile
FROM tomohito/grant-n-z:latest

ENV GRANT_N_Z_MYSQL_HOST="{HOST NAME}"
ENV GRANT_N_Z_MYSQL_USER="{HOST USER NAME}"
ENV GRANT_N_Z_MYSQL_PASS="{HOST PASSWORD}"
ENV GRANT_N_Z_MYSQL_PORT="{HOST PORT}"
ENV GRANT_N_Z_MYSQL_DB="grant_n_z"
ENV GRANT_N_Z_PRIVATE_KEY="{PRIVATE KEy}"
```

### Run on server

1. Go get command.
   ```bash
   $ go get github.com/tomoyane/grant-n-z
   ```
    
    * If you not install go, you can use binary.

       * [Binary file]()

2. You have to create `grant_n_z` database on mysql. Grant-n-z supports only mysql.

3. Run application.
   ```bash
   $ ./grant-n-z
   ```

## API Reference

### POST /v1/users
#### Header
| Parameter | Value |
|:---|---|
| Content-Type | application/json |


#### Request Body
| Parameter | Value |
|:---|---|
| username | string |
| email | string |
| password | string |


#### Response
##### 201 Created
Header 

| Parameter | Value |
|:---|---|
| Location | {hostname}/v1/users/{uuid} |

Body

| Parameter | Value |
|:---|---|
| message | string |

##### 400 BadRequest
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

##### 422 UnProcessableEntity
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

##### 500 InternalServerError
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

### POST /v1/tokens
#### Header
| Parameter | Value |
|:---|---|
| Content-Type | application/json |


#### Request Body
| Parameter | Value |
|:---|---|
| email | string |
| password | string |


#### Response
##### 200 OK
Body

| Parameter | Value |
|:---|---|
| token | string |

##### 400 BadRequest
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

##### 422 UnProcessableEntity
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

##### 500 InternalServerError
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

### POST /v1/grants
#### Header
| Parameter | Value |
|:---|---|
| Content-Type | application/json |
| Authorization | {TOKEN} |

#### Response
##### 200 OK
Body

| Parameter | Value |
|:---|---|
| authority | string |

##### 401 Unauthorized
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

##### 403 Forbidden
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |

##### 500 InternalServerError
| Parameter | Value |
|:---|---|
| code | int |
| message | string |
| detail | string |


## License
[MIT](https://github.com/tomoyane/grant-n-z/blob/master/LICENSE)

