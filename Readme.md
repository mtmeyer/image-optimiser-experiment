# Basic test of different image optimisation libraries

## Run image optimisation servers

```sh 
docker compose build
docker compose up

```

## Load tester 

Not 100% convinced this is working as expected or is overly reliable 🤷


```sh
cd load-test
go run . --url https://some-site.com --requests 5 --concurrent 5
```

**Image test from unsplash comparing response times with and without image optimisation**

```sh
go run . --url http://localhost:3010/img/https://images.unsplash.com/photo-1690200371608-1b8b6a43d973\?ixlib\=rb-4.0.3\&ixid\=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D\&auto\=format\&fit\=crop\&w\=987\&q\=80 --requests 2 --concurrent 10

go run . --url https://images.unsplash.com/photo-1690200371608-1b8b6a43d973\?ixlib\=rb-4.0.3\&ixid\=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D\&auto\=format\&fit\=crop\&w\=987\&q\=80 --requests 2 --concurrent 10
```

### Flags

`url`: url of the site to test response times

`requests`: number of requests in a batch (actual requests is `requests * concurrent`)

`concurrent`: number of concurrent request batches to run
