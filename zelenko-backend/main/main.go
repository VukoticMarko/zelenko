package main

import (
	"fmt"
	"net/http"
	"zelenko-backend/router"
)

func main() {
	var (
		httpRouter router.Router = router.NewMuxRouter()
	)
	const port string = ":2023"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and runing")
	})

	httpRouter.SERVE(port)

}
