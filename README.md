# The Go implementation of the TecPoster Server

## Get started

### install golang

```
sudo pacman -S go
```

~/.zprofile

```
export GOPATH="$HOME/space/go"
export PATH="$PATH:$HOME/bin:$GOPATH/bin"
```

### dep

<https://golang.github.io/dep>

```
go get -u github.com/golang/dep/cmd/dep
dep init
dep status
dep ensure
	dep ensure -add github.com/pkg/errors
```


### gcc and other tools in base-devel

<https://linuxhint.com/compile_c_program_linux_gcc/>

```
sudo pacman -S base-devel
gcc --version
```

### ref pkgs

* github.com/gorilla/websocket
* github.com/google/uuid
* github.com/btcsuite/btcutil/base58
* github.com/mattn/go-sqlite3


## APIs

* post.commit
* post.fetch
* post.list
* post.search
* post.create
* post.edit
* draft.save
* draft.fetch
* draft.list
* delayed 
	* user.reg
	* user.login
	* user.logout
	* (deprecate) user.refresh-token

## Data Schema

* [uid]
	* post
		* id
		* commitID
		* created
	* commit
		* id
		* postID
		* contentID
		* created
	* content
		* id
		* type
		* content
	* draft
		* id
		* content
		* changed

* [uid]post
	* pid
	* (d) uid
	* changed ?
	* title
	* pcid (post commit id): [pid]-[index]
* [uid]postCommit
	* cid (content id)
	* pid
	* content
* [uid]draft
	* pid
	* (d) uid
	* changed
	* cnt
* [uid]/txn
	* timestamp (int)
	* pcid
* content
	* typ: markdown / text / html
	* body: ""


## Request & Response Data Structure

Request

```
{
	"cmd": "post.commit",
	"token": "[token]", // deprecate
	"params: {
	}
}
```

Response

```
// default
{
	"cmd": "post.commit",
	"status": "ok",
	"data": {
	}
}

// error
{
	"cmd": "post.commit",
	"status": "error",
	"data": {
		"error": "[error message]"
	}
}
```

## User

### user.reg

request

```
{
	"cmd": "user.reg",
	"token": "",
	"params": {
		"email": "zhanjh@126.com",
		"username": "zhanjh",
		"password": "123456789"
	}
}
```

Error

* Error email format
* Username too short - minimum length is 7
* Password too short - minimum length is 7
* Email already exists
* Username already exists

Success

```
{
	"cmd": "user.reg",
	"status": "ok",
	"data": {
		"email": "zhanjh@126.com",
    "uid": "3HrupZrJFPJnhqvhDXDmEb",
    "username": "zhanjh"
	}
}
```

### user.login

request

```
{
	"cmd": "user.login",
	"params": {
		"email": "zhanjh@126.com",
		"password": "xxx"
	}
}
```

Error

* Email not found
* Incorrect password

Success

```
{
	"cmd": "user.login",
	"status": "ok",
}
```

### user.logout

request

```
{
	"cmd": "user.logout",
	"params": {}
}
```

Success

```
{
	"cmd": "user.logout",
	"status": "ok"
}
```

## post

### post.create

request

```
{
	"cmd": "post.create",
	"params": {}
}
```

Error

* Not login

Success

```
{
	"cmd": "post.create",
	"status": "ok",
	"data": {
		"pid": "***"
	}
}
```

## dependencies

```
go get -u -v github.com/btcsuite/btcutil/base58
go get -u -v github.com/dgraph-io/badger
```

## compile

```
env GOOS=windows GOARCH=amd64 go build -v -o ws-server.amd64.exe cmd/server.go
```

## commands

```
go run cmd/server.go --datadir="/tmp/tec"

go build -o ws-server cmd/server.go
```

## test

<https://blog.golang.org/cover>

```
go test ./db/sqlite3 -coverprofile=coverage.out fmt && go tool cover -func=coverage.out && cat coverage.out && rm coverage.out
```

## examples

```
{"cmd": "post.create"}

{"cmd": "draft.list"}

```
