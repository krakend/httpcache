httpcache
=========

This a fork of github.com/gregjones/httpcache which includes various fixes.
Original backends are removed since are not used by KrakenD.

[![GoDoc](https://godoc.org/github.com/krakendio/httpcache?status.svg)](https://godoc.org/github.com/krakendio/httpcache)

Package httpcache provides a http.RoundTripper implementation that works as a mostly [RFC 7234](https://tools.ietf.org/html/rfc7234) compliant cache for http responses.

It is only suitable for use as a 'private' cache (i.e. for a web-browser or an API-client and not for a shared proxy).

This project isn't actively maintained; it works for what I, and seemingly others, want to do with it, and I consider it "done". That said, if you find any issues, please open a Pull Request and I will try to review it. Any changes now that change the public API won't be considered.

License
-------

-	[MIT License](LICENSE.txt)
