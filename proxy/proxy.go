package chproxy

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"io/ioutil"
	"math/rand"
	"net/http"
)

type Proxy struct {
	ApiKey    string
	Timeout   int
	Addresses []string
}

// Wrap a connection and its existing reader
type connWithReader struct {
	net.Conn
	io.Reader
}

func (c connWithReader) Read(p []byte) (int, error) {
	return c.Reader.Read(p)
}

// Set up a TCP listener on given address and process incoming requests using ProcessRequest
func (p Proxy) ServeAndHandle(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("error listening on address: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error accepting connection: %v", err)
		}
		go p.ProcessRequest(conn)
	}
}

// Process an incoming TCP request:
// - Check that correct headers are present
// - If yes forward the connection
// - Otherwise respond with HTTP response
func (p Proxy) ProcessRequest(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(time.Duration(p.Timeout) * time.Second))

	reader, errCode, errMessage := p.checkHeaders(conn)
	if errCode != -1 {
		res := p.createErrorResponse(errCode, errMessage)
		conn.Write(res.Bytes())
		return
	}

	c := connWithReader{
		Conn:   conn,
		Reader: reader,
	}
	p.forwardConn(c)
}

// Check that HTTP request read from byte stream contains correct api key, return a copy of the
// incoming byte stream or an HTTP error code and error message.
func (p Proxy) checkHeaders(conn net.Conn) (*bytes.Buffer, int, string) {
	var buf bytes.Buffer
	tee := io.TeeReader(conn, &buf)
	reader := bufio.NewReader(tee)

	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Printf("error parsing request: %v", err)
		return nil, http.StatusBadRequest, "cannot parse incoming HTTP request"
	}
	request_key := req.Header.Get("Api-Key")

	if len(request_key) == 0 {
		return nil, http.StatusUnauthorized, "missing api key"
	} else if request_key != p.ApiKey {
		return nil, http.StatusUnauthorized, "wrong api key"
	} else {
		return &buf, -1, ""
	}
}

// Open connection to random upstream address and do a bi-directionnal copy of byte streams.
func (p Proxy) forwardConn(conn net.Conn) {
	defer conn.Close()

	rand.Seed(time.Now().Unix())
	address := p.Addresses[rand.Intn(len(p.Addresses))]

	upstream, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("error dialing upstream: %v", err)
		res := p.createErrorResponse(http.StatusInternalServerError, "")
		conn.Write(res.Bytes())
		return
	}
	defer upstream.Close()
	upstream.SetReadDeadline(time.Now().Add(time.Duration(p.Timeout) * time.Second))

	go io.Copy(upstream, conn)
	io.Copy(conn, upstream)
}

// Wrap given code and message in a http.Response and write it to a bytes.Buffer.
func (p Proxy) createErrorResponse(code int, message string) *bytes.Buffer {
	r := &http.Response{
		Status:        fmt.Sprintf("%d %s", code, http.StatusText(code)),
		StatusCode:    code,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(message)),
		ContentLength: int64(len(message)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
	var buf bytes.Buffer
	r.Write(&buf)
	return &buf
}
