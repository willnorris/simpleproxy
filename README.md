simpleproxy provides a basic example of using Go's [httputil.ReverseProxy][] to
proxy all HTTP requests to a target URL.

[httputil.ReverseProxy]: https://golang.org/pkg/net/http/httputil/#ReverseProxy

Set the environment variable `TARGET_URL` to the URL requests should be proxied
to.  For example:

    TARGET_URL="http://example.com/foo" go run .

The URL path of the request will be appended to the target URL path.  With the
above configuration, the URL `http://localhost:8080/bar` will be proxied to
`http://example.com/foo/bar`.

---

imageproxy is copyright Google, but is not an official Google product.  It is
available under the [Apache 2.0 License](./LICENSE).
