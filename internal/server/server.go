package server

import "github.com/nimsaysm/go-react-library-management/internal/router"

func StartServer() {
	router := router.SetRouter()

	router.Run()
}