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
```

- Missing headers ...
`curl -v -k -sL -X GET https://student-Virtualbox:9443/taxii2/`
`curl -v -k -sL -H "Accept: application/taxii+json" -X GET https://student-Virtualbox:9443/taxii2/`

Need to store user1:hash(Password1234!@#$) into etcd server.

- With good user user1:Password1234!@#$

`curl -v -k -sL -H "Accept: application/taxii+json" -H "Authorization: Basic dXNlcjE6UGFzc3dvcmQxMjM0IUAjJA==" -X GET https://student-Virtualbox:9443/taxii2/`

- With unmatched user user1:Password1234!@#$
`curl -v -k -sL -H "Accept: application/taxii+json" -H "Authorization: Basic dXNlcjE6UGFzc3dvcmQxMjM0IUAjNA==" -X GET https://student-Virtualbox:9443/taxii2/`

