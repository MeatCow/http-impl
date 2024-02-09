package httpimpl

import (
	"errors"
	"strconv"
	"strings"
)

type Response struct {
	Status int
}

type URIScheme string

const (
	Http  URIScheme = "http"
	Https URIScheme = "https"
)

type Method string

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)

type Url struct {
	scheme URIScheme
	host string
	port uint16
	path string
}

type Request struct {
	url Url
	method Method
}

func Get(url string) (*Response, error) {
	_, err := parseUrl(url)
	if err != nil {
		return nil, err
	}

	

	return &Response{Status: 200}, nil
}

func parseUrl(url string) (*Url, error) {
	state := Url{
		port: 80,
		path: "/",
	}

	if url == "" {
		return nil, errors.New("Empty URL")
	}

	if strings.HasPrefix(url, "http://") {
		state.scheme = Http
		url, _ = strings.CutPrefix(url, "http://")
	} else if strings.HasPrefix(url, "https://") {
		return nil, errors.New("No support for https")
	} else {
		return nil, errors.New("Not a valid HTTP URI Scheme")
	}

	origin, path, found := strings.Cut(url, "/")
	if found {
		state.path += path;
	}


	host, port, found := strings.Cut(origin, ":")
	if found {
		parsedPort, err := strconv.ParseUint(port, 10, 16)
		if err != nil {
			return nil, errors.New("Unable to parse port")
		}
		state.port = uint16(parsedPort)
	}
	state.host = host

	return &state, nil
}