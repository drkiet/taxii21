# Progress report
Implementing TAXII 2.1 Spec located here 
`https://docs.oasis-open.org/cti/taxii/v2.1/cs01/taxii-v2.1-cs01.html`

## 12/5/2020
- Completed the discovery query "/taxii2/"
- Used self signed certificate via `openssl` tool
- Tested with Postman using these requests:

Request:
```
curl --location --request GET 'https://localhost:9443/taxii2/' \
--header 'Accept: application/taxii+json;version=2.1' \
--header 'Cookie: JSESSIONID=00D768F9C14B2A7528E236364F0C98CA'
```

Response:

```
{
    "title": "TAXII 2.1 Server",
    "description": "TAXII 2.1 Server in GoLang",
    "contact": "Kiet T. Tran, Ph.D.",
    "default": "https://student-VirtualBox/api1/",
    "api_roots": [
        "https://student-VirtualBox/api1/",
        "https://student-VirtualBox/trustgroup1/",
        "https://student-VirtualBox/trustgroup2/"
    ]
}
```