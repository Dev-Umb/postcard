package server

import (
	"github.com/gin-gonic/gin"
	"postcard/apiserver/internal/option"
	"postcard/apiserver/internal/route"
	"postcard/apiserver/internal/store"
)

type Server struct {
	Web *gin.Engine
}

func NewServer(opt *option.Options) (*Server, error) {
	sro := store.DataStore{}
	g := route.ApiServer(opt.WebOptions.Engine, sro)
	return &Server{
		Web: g,
	}, nil
}
