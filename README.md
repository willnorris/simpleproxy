simpleproxy provides a single host reverse proxy that proxies HTTP requests to a
target URL.

## Usage

Set the environment variable `TARGET_URL` to the URL requests should be proxied
to.  The URL path of the request will be appended to the target URL path.

For example, when running:

    TARGET_URL="http://example.com/foo" go run .

The URL `http://localhost:8080/bar` will be proxied to
`http://example.com/foo/bar`.

Setting the environment variable `FORCE_SSL=1` will redirect all non-secure
requests to use https.  Simpleproxy relies on the `X-Forwarded-Proto: http`
request header being set to identify non-secure requests.

## What's this for?

Reverse proxies have assorted uses.  This project was designed to provide simple
SSL support for a custom domain on a [Google Cloud Storage][] bucket.

GCS [does not support][] HTTPS serving with custom domains.  For various
reasons, I didn't want to use any of the recommended solutions, so I built this
to run on [Google Cloud Run][] (which does support SSL for custom domains) and
proxy requests to my GCS bucket.

[Google Cloud Storage]: http://cloud.google.com/storage
[does not support]: https://cloud.google.com/storage/docs/troubleshooting#https
[Google Cloud Run]: https://cloud.google.com/run/

While I built it for using Cloud Run with GCS, there is nothing Google-specific
about it.  It's just a simple reverse proxy running in a docker container.

---

imageproxy is copyright Google, but is not an official Google product.  It is
available under the [Apache 2.0 License](./LICENSE).
