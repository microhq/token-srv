# Token Server

Token server is a microservice to generate and store random tokens. Useful for invites.

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Start a mysql database

4. Download and start the service

	```shell
	go get github.com/micro/token-srv
	token-srv --database_url="root:root@tcp(192.168.99.100:3306)/token"
	```

	OR as a docker container

	```shell
	docker run microhq/token-srv --database_url="root:root@tcp(192.168.99.100:3306)/token" --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The API
Token server implements the following RPC Methods

Record
- Create
- Read
- Update
- Delete
- Search
- Generate


### Record.Generate
```shell
micro query go.micro.srv.token Record.Generate
{
	"token": {
		"created": 1.453406213e+09,
		"id": "1dffcde9-c079-11e5-827c-aaa11963c131",
		"name": "BDlMvievTXqgU1kQ",
		"namespace": "default",
		"updated": 1.453406213e+09
	}
}
```

### Record.Search 
```shell
micro query go.micro.srv.token Record.Search '{"namespace": "default", "limit": 10}'
{
	"tokens": [
		{
			"created": 1.453406213e+09,
			"id": "1dffcde9-c079-11e5-827c-aaa11963c131",
			"name": "BDlMvievTXqgU1kQ",
			"namespace": "default",
			"updated": 1.453406213e+09
		}
	]
}
```

