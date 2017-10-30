package urlx

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func Join(p ...string) string {
	n := len(p)
	if n == 0 {
		return ""
	}
	p1 := make([]string, n, n)
	for i, pi := range p {
		p1[i] = strings.TrimPrefix(strings.TrimSuffix(pi, "/"), "/")
	}
	return strings.Join(p1, "/")
}

func Abs(baseURL, path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	scheme, host := "", ""
	if url1, err := url.Parse(baseURL); err != nil {
		return ""
	} else {
		scheme, host = url1.Scheme, url1.Host
	}
	newURL := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	return newURL.String()
}

func EncodeQuery(q map[string]string) string {
	if len(q) == 0 {
		return ""
	}
	vals := url.Values{}
	for k, v := range q {
		vals.Set(k, v)
	}
	return vals.Encode()
}

func ToValues(q map[string]string) url.Values {
	vals := url.Values{}
	for k, v := range q {
		vals.Set(k, v)
	}
	return vals
}

func SplitPort(hostAndPort string) (string, int) {
	if !strings.Contains(hostAndPort, ":") {
		return hostAndPort, 0
	}
	arr := strings.Split(hostAndPort, ":")
	host := strings.Join(arr[0:len(arr)-1], ":")
	port, err := strconv.ParseInt(arr[len(arr)-1], 10, 32)
	if err != nil {
		return host, 0
	}
	return host, int(port)
}

func JoinHostAndPort(host string, port int) string {
	if port <= 0 {
		return host
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func TopLevelDomainOfHost(hostAndPort string, withPort bool) string {
	host, port := SplitPort(hostAndPort)
	arr := strings.Split(host, ".")
	n := len(arr)
	if n <= 2 {
		if withPort {
			return JoinHostAndPort(host, port)
		} else {
			return host
		}
	}
	st := 0
	for i := n - 1; i >= 0; i-- {
		seg := arr[i]
		switch st {
		case 0:
			if !TopLevelDomainsSet.Has(seg) {
				st = 1
			}
		case 1:
			arr[i] = ""
		}
	}
	buff := bytes.NewBufferString("")
	for i, seg := range arr {
		if seg != "" {
			buff.WriteString(seg)
			if i < n-1 {
				buff.WriteString(".")
			}
		}
	}
	if withPort {
		return JoinHostAndPort(buff.String(), port)
	} else {
		return buff.String()
	}
}

func HostOf(urlStr string, withPort, topLevel bool) string {
	url1, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	if topLevel {
		return TopLevelDomainOfHost(url1.Host, withPort)
	} else {
		if withPort {
			return url1.Host
		} else {
			host, _ := SplitPort(url1.Host)
			return host
		}
	}
}
