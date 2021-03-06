package zerver

import (
	"io"
	"strings"
)

const (
	// Http Header
	HEADER_CONTENTTYPE     = "Content-Type"
	HEADER_CONTENTLENGTH   = "Content-Length"
	HEADER_SETCOOKIE       = "Set-Cookie"
	HEADER_REFER           = "Referer"
	HEADER_CONTENTENCODING = "Content-Encoding"
	HEADER_USERAGENT       = "User-Agent"
	HEADER_ACCEPTENCODING  = "Accept-Encoding"
	HEADER_CACHECONTROL    = "Cache-Control"
	HEADER_EXPIRES         = "Expires"
	HEADER_AUTHRIZATION    = "Authorization"
	HEADER_METHODOVERRIDE  = "X-HTTP-Method-Override"

	// ContentEncoding
	ENCODING_GZIP    = "gzip"
	ENCODING_DEFLATE = "deflate"

	// Request Method
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
	PATCH  = "PATCH"

	// Content Type
	CONTNTTYPE_PLAIN = "text/plain;charset=utf-8"
	CONTENTTYPE_HTML = "text/html;charset=utf-8"
	CONTENTTYPE_XML  = "application/xml;charset=utf-8"
	CONTENTTYPE_JSON = "application/json;charset=utf-8"
)

// parseRequestMethod convert a string to request method, default use GET
// if string is empty
func parseRequestMethod(s string) string {
	if s == "" {
		return GET
	}
	return strings.ToUpper(s)
}

// parseContentType parse content type
func parseContentType(str string) string {
	if str == "" {
		return CONTENTTYPE_HTML
	}
	return strings.ToLower(strings.TrimSpace(str))
}

func ErrorWrite(err error, w io.Writer, data []byte) (int, error) {
	if err != nil {
		return 0, err
	}
	return w.Write(data)
}

func ErrorRead(err error, r io.Reader, data []byte) (int, error) {
	if err != nil {
		return 0, err
	}
	return r.Read(data)
}
