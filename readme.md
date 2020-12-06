# TAXII 2.1 Server project

## getting started

```shell script
mkdir taxii21
go mod init github.com/drkiet/taxii21
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
go install
```

Website:
`https://github.com/denji/golang-tls`:

```shell script
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.cert -days 3650
curl -k -sL -X GET https://student-Virtualbox:9443/taxii2/
```