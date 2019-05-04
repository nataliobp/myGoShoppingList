package core

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type LoggingMiddleware struct {
	Next http.Handler
}

func (m *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if nil == m.Next {
		m.Next = http.DefaultServeMux
	}

	requestString := r.Method + " " + r.URL.RequestURI()

	if r.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		requestString += "\n" + string(bodyBytes)
	}

	log.Println(requestString)

	m.Next.ServeHTTP(w, r)
}
