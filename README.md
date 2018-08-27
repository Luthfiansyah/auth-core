# Auth Core


## How to Install ?

Under `$GOPATH`, do the following command
```
mkdir -p src/github.com/auth-core/
cd src/github.com/auth-core/
git clone <url>
```

And you should now find the source code under `auth-core` directory.

## Directory structure

```
  + your_gopath/
  |
  +--+ src/
  |  |
  |  +--+ github.com/
  |     |
  |     +--+ auth-core/
  |         |
  |         +--+ main.go
  |            + routing.go
  |            + app/
  |            |
  |            +--+ + controllers/
  |                 + models/
  |                 + types/
  |                 + repositories/
  |            + config/
  |            + database/
  |            + documents/
  |            + export/
  |            + html/
  |            + logging/
  |            + public/
  |            + vendor/
  |            + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

Get into the `auth-core` folder by

```
cd auth-core/
```


Under that folder you need to run below command to install required library dependency file glide.yaml

```
glide install
```
## How to compile and run the code ?
And then following this command :
```
go build && ./auth-core -c config.toml
```

### OR compile Install and run the code ?
And then following this command :
```
go install
$GOPATH/bin/auth-core -c config.toml
```

The configuration is described in the config.toml file.
In particular, in the [server] section, you will find:
- mode: set it to development, production, localserver
- port: the port the server will bind to
- debug: the debug flag keeps greater amount of logging enabled, including gin logging

## How do we know this application server is running?
You can try calling the root context of url
for the example, the base url of this server is

http://localhost:8888

Just call it in normal browser and you will get message something like

```
{
    "message": "Auth Core Core Run on development mode",
    "start_time": "[06 July 2018] 02:45:51 WIB"
}
```

## How to test a code ?

To run all the test in current folder (we have `dk_test.go` and `trd_test.go`)
```
go test
```

To run specific test file (for the example only `trd_test.go` file)

```
go test trd_test.go
```

To run specific test method (for the example only `Inquiry` method in `dk_test.go`)

```
go test -run TestDKInquiry
```



# How to test with Postman ?

When the server is up, you can use a Postman apps (chrome plugins) to do testing for every API provided. Import the postman collection API to your Postman from the `postman` folder.


## How to add new service API in app ?
```
Style
1. At `app/controllers/modulename_controller.go` add new controller
2. At `app/models/modelname.go` add new model
3. At `app/types/typename.go` add new type
4. At `app/repositories/repositoryname.go` add new repository

```
## How to commit and push the change into repository with glide (optional step )?

We are using glide for dependency management, so first we need to get glide
```
go get github.com/Masterminds/glide
```

Now every time we add new dependency, (still under `auth-core` directory) we need to update the glide by calling
```
$GOPATH/src/github.com/auth-core/
$glide get "repo_name"
```
glide will create and copy the required dependency to vendor folder. This included libs vendor need to push to repository


## How to build a Docker Image ?

Just run this command
```
PENDING
```

## Ready API

```
GET   /v1
```

## General rules of json parameter

```
PENDING
```

## Sequence diagram

```
PENDING
```

## Response From API

```

PING GET http://localhost:8888
=====
{
    "message": "auth-core Core Run on development mode",
    "start_time": "[06 July 2018] 02:45:51 WIB"
}


```

## Login From API

```

POST http://localhost:8888/v1/auth/token
=====
Request : 
{
	"username" : "admin",
	"password" : "admin"
}

Response :
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjEsImV4cGlyZWRfYXQiOjE1NjY4ODE0MjIsImlzcyI6ImFkbWluIn0.iPfnJPR0Aal3XOouND1WfXOhCn31Lro-zQx2crYD0no",
        "expired_at": 1566881422
    },
    "general_response": {
        "response_status": true,
        "response_code": "000",
        "response_message": "Success",
        "response_timestamp": "2018-08-27T04:50:22.350994013Z"
    }
}

```

## Redis Get Key Example

```

$redis-cli -h localhost -p 6379 
localhost:6379>get client_1
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOjEsImV4cGlyZWRfYXQiOjE1NjY4ODE0MjIsImlzcyI6ImFkbWluIn0.iPfnJPR0Aal3XOouND1WfXOhCn31Lro-zQx2crYD0no"

```

====================================================================================

## auth-core Core deployment

The server is at: auth-core

```
DEVELOPMENT:
   host   : http://localhost:8888
   branch : develop
   mode   : development. (local)
   folder : auth-core

STAGING   :
   host   : http://new-domain
   branch : staging
   mode   : staging. (staging)
   folder : auth-core

PRODUCTION:
   host   : http:///new-domain
   branch : release
   mode   : production. (production)
   folder : auth-core
```

Author Moh Reza Luthfiansyah
