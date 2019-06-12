simpleproxy provides a basic example of using Go's [httputil.ReverseProxy][] to
proxy all HTTP requests to a target URL.

[httputil.ReverseProxy]: https://golang.org/pkg/net/http/httputil/#ReverseProxy

Set the environment variable `TARGET_URL` to the URL requests should be proxied
to.  For example:

    TARGET_URL="http://example.com/foo" go run .

The URL path of the request will be appended to the target URL path.  With the
above configuration, the URL `http://localhost:8080/bar` will be proxied to
`http://example.com/foo/bar`.

Setting the environment variable `FORCE_SSL=1` will redirect all non-secure
requests to use https.  Simpleproxy relies on the `X-Forwarded-Proto: http`
request header being set to identify non-secure requests.

---

imageproxy is copyright Google, but is not an official Google product.  It is
available under the [Apache 2.0 License](./LICENSE).
