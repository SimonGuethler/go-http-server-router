package router

import (
	"errors"
	"net"
	"strings"
)

type HttpParser struct {
	Buffer       []byte
	BufferString string
	BufferParts  []string
	Method       HTTPMethod
	Path         string
	PathParams   map[string]string
	Query        string
	QueryParams  map[string]string
	Headers      map[string]string
	Body         string
}

func NewHttpParser(buffer []byte) *HttpParser {
	bufferString := string(buffer)
	bufferParts := strings.Split(bufferString, "\r\n")
	for i, part := range bufferParts {
		bufferParts[i] = strings.TrimSpace(part)
	}
	return &HttpParser{
		Buffer:       buffer,
		BufferString: bufferString,
		BufferParts:  bufferParts,
	}
}

func (p *HttpParser) parseRequestLine() error {
	requestLineParts := strings.Fields(p.BufferParts[0])
	if len(requestLineParts) != 3 {
		return errors.New("invalid request line")
	}

	method, err := ParseHTTPMethod(requestLineParts[0])
	if err != nil {
		return err
	}
	p.Method = method

	rawPath := requestLineParts[1]
	if strings.Contains(rawPath, "?") {
		pathParts := strings.Split(rawPath, "?")
		path, err := SanitizePath(pathParts[0], false)
		if err != nil {
			return err
		}
		p.Path = path
		p.Query = pathParts[1]
	} else {
		path, err := SanitizePath(rawPath, false)
		if err != nil {
			return err
		}
		p.Path = path
	}

	return nil
}

func (p *HttpParser) parseBuffer() error {
	p.Headers = make(map[string]string)
	for i := 1; i < len(p.BufferParts); i++ {
		if p.BufferParts[i] == "" {
			p.Body = strings.Join(p.BufferParts[i+1:], " ")
			break
		}
		headerParts := strings.SplitN(p.BufferParts[i], ":", 2)

		if len(headerParts) != 2 {
			return errors.New("invalid header")
		}
		p.Headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
	}

	return nil
}

func (p *HttpParser) parseQuery() error {
	if p.Query == "" {
		return nil
	} else {
		p.QueryParams = make(map[string]string)
		queryParts := strings.Split(p.Query, "&")
		for _, queryPart := range queryParts {
			queryParamParts := strings.Split(queryPart, "=")
			if len(queryParamParts) != 2 {
				return errors.New("invalid query parameter")
			}
			p.QueryParams[strings.TrimSpace(queryParamParts[0])] = strings.TrimSpace(queryParamParts[1])
		}
	}

	return nil
}

func (p *HttpParser) Parse() error {
	err := p.parseRequestLine()
	if err != nil {
		return err
	}

	err = p.parseQuery()
	if err != nil {
		return err
	}

	err = p.parseBuffer()
	if err != nil {
		return err
	}

	return nil
}

func ParseHTTPMethod(method string) (HTTPMethod, error) {
	switch method {
	case "GET":
		return GET, nil
	case "POST":
		return POST, nil
	case "PUT":
		return PUT, nil
	case "PATCH":
		return PATCH, nil
	case "DELETE":
		return DELETE, nil
	case "OPTIONS":
		return OPTIONS, nil
	case "HEAD":
		return HEAD, nil
	case "CONNECT":
		return CONNECT, nil
	case "TRACE":
		return TRACE, nil
	default:
		return GET, errors.New("invalid method")
	}
}

func ReadRequest(conn net.Conn, maxRequestSize int) ([]byte, error) {
	var buffer []byte
	tmp := make([]byte, maxRequestSize)

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err.Error() == "EOF" && n == 0 {
				return buffer, nil
			}
			return nil, err
		}

		buffer = append(buffer, tmp[:n]...)

		if n < maxRequestSize {
			break
		}
	}

	return buffer, nil
}

func ExtractPathVariables(routePath string, requestPath string) map[string]string {
	routeParts := SplitPath(routePath)
	requestParts := SplitPath(requestPath)

	var pathVariables = make(map[string]string)
	for i, part := range routeParts {
		if IsPathParam(part) {
			pathVariables[part[1:]] = requestParts[i]
		}
	}
	return pathVariables
}
