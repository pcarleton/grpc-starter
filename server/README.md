# Running locally

Server:
```
# Create local certs
bash ./make_localhost_certs.sh

go run bin/run_server/run_server.go
```

Client:
```
go run bin/run_client/run_client.go -address localhost:5001
```
