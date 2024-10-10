package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ServerTestingHanldeHello(test *testing.T) {
	server:=New(NewConfig())
	rec:=httptest.NewRecorder()
	request,_:=http.NewRequest(http.MethodGet,"/пошел нахуй",nil)
	server.handleHello().ServeHTTP(rec,request)
	assert.Equal(test,rec.Body.String(),"пошел нахуй")
}