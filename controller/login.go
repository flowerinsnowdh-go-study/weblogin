package controller

import "net/http"

func (c *Controller) ControlLogin(w http.ResponseWriter, r *http.Request) {
	c.Mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
}
