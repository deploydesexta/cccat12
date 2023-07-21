package main

import (
	"context"
	"github.com/deploydesexta/cccat12/src/infrastructure/http"
	"github.com/deploydesexta/cccat12/src/infrastructure/jsonutil"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	http2 "net/http"
)

type ChiRouterAdapter struct {
	Chi chi.Router
}

type ChiRequestAdapter struct {
	w   http2.ResponseWriter
	req *http2.Request
}

func NewChiRouterAdapter() *ChiRouterAdapter {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return &ChiRouterAdapter{Chi: r}
}

func (r *ChiRouterAdapter) Router(ctrl http.Router) {
	ctrl.Bind(r)
}

func (r *ChiRouterAdapter) Get(path string, handler http.HandlerFunc) {
	r.Chi.Get(path, func(w http2.ResponseWriter, r *http2.Request) {
		err := handler(&ChiRequestAdapter{w, r})
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	})
}

func (r *ChiRouterAdapter) Post(path string, handler http.HandlerFunc) {
	r.Chi.Post(path, func(w http2.ResponseWriter, r *http2.Request) {
		err := handler(&ChiRequestAdapter{w, r})
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	})
}

func (r *ChiRouterAdapter) Start(port string) error {
	log.Printf("⇨ http server started on [::]%s\n", port)
	return http2.ListenAndServe(port, r.Chi)
}

func (r *ChiRequestAdapter) Bind(i interface{}) error {
	return jsonutil.FromReader(r.req.Body, i)
}

func (r *ChiRequestAdapter) JSON(code int, i interface{}) error {
	b, err := jsonutil.ToJsonBytes(i)
	if err != nil {
		render.Status(r.req, http2.StatusInternalServerError)
		render.PlainText(r.w, r.req, err.Error())
	}

	r.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = r.w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChiRequestAdapter) Param(name string) string {
	return chi.URLParam(r.req, name)
}

func (r *ChiRequestAdapter) QueryParam(name string) string {
	return r.req.URL.Query().Get(name)
}

func (r *ChiRequestAdapter) String(code int, s string) error {
	render.Status(r.req, code)
	render.PlainText(r.w, r.req, s)
	return nil
}

func (r *ChiRequestAdapter) Context() context.Context {
	return r.req.Context()
}
