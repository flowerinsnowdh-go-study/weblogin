package controller

import (
	"github.com/flowerinsnowdh/weblogin/service"
	"net/http"
)

type Controller struct {
	Mux     *http.ServeMux
	Service *service.Service
}
