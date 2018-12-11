# :cow: Outlaw

Outlaw is a redirector that uses [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) to generate SSL certificates automatically and redirect wildcards and specific routes.

Features

* Redirect HTTP to HTTPS when you run behind a CDN (different origins)
* Redirect multiple domains to a canonical domain
* Create special redirects schemes (e.g. for iOS apps)

## Development

[Boulder](https://github.com/letsencrypt/boulder)

```
 ./bin/outlaw --dynamodb --http-port 5002 --https-port 5001 --acme-url http://localhost:4000/directory --verbose
```

