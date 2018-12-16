# :cow: Outlaw

Outlaw is a redirector that uses [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) to generate SSL certificates automatically and redirect wildcards and specific routes.

Features

* Redirect parked Domains to a URL
* Redirect HTTP to HTTPS when you run behind a CDN (different origins)
* Redirect multiple domains to a canonical domain
* Create special redirects schemes (e.g. for iOS apps)

Databases

* [MongoDB](https://www.mongodb.com)
* DynamoDB (coming soon)
* more coming soon...

## Client

Outlaw includes a [gRPC](https://grpc.io) client to control the server.

## Docker

```
docker run andersnormal/outlaw:1.0.0-beta.0 --mongo --mongo-endpoint mongo --mongo-username root --mongo-password example --mongo-auth-database admin --verbose
```

## Staging

Before moving to production it is recommended to test in the [Staging Environment](https://acme-staging.api.letsencrypt.org/directory) of [Let's Encrypt](https://letsencrypt.org). The URL for the Staging ACME V1 Environment can be set via `--acme-url https://letsencrypt.org`.

## Development

[Boulder](https://github.com/letsencrypt/boulder) is supported to test in a local development environment.

```
 ./bin/outlaw --dynamodb --http-port 5002 --https-port 5001 --acme-url http://localhost:4000/directory --verbose
```

