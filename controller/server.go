package controller

import (
	"magantas/catalyst/store"
)

type Server struct {
	Store *store.Store
}

func NewServer() Server {

	server := Server{
		Store: store.NewDB(),
	}

	return server

}
