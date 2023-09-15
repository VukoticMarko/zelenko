package main

import (
	"fmt"
	"net/http"
	"zelenko-backend/controller"
	"zelenko-backend/crdt"
	"zelenko-backend/repository"
	"zelenko-backend/router"
	"zelenko-backend/service"
)

func main() {
	var (
		g_counter            crdt.GCounter                   = *crdt.NewGCounter()
		greenScoreRepository repository.GreenScoreRepository = repository.NewGreenScoreRepository()
		greenScoreService    service.GreenScoreService       = service.NewGreenScoreService(greenScoreRepository, g_counter)
		greenScoreController controller.GreenScoreController = controller.NewGreenScoreController(greenScoreService)

		httpRouter router.Router = router.NewMuxRouter()
	)
	const port string = ":2023"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and runing")
	})

	httpRouter.POST("/addScore", greenScoreController.AddOne)
	httpRouter.POST("/subScore", greenScoreController.SubOne)

	httpRouter.SERVE(port)

}
