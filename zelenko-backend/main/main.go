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

		greenObjectRepository repository.GreenObjectRepository = repository.NewGreenObjectRepository()
		greenObjectService    service.GreenObjectService       = service.NewGreenObjectService(greenObjectRepository)
		greenObjectController controller.GreenObjectController = controller.NewGreenObjectController(greenObjectService)

		userRepository repository.UserRepository = repository.NewUserRepository()
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService)

		httpRouter router.Router = router.NewMuxRouter()
	)
	const port string = ":2023"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and runing")
	})

	httpRouter.POST("/addScore", greenScoreController.AddOne)
	httpRouter.POST("/subScore", greenScoreController.SubOne)
	httpRouter.POST("/addObject", greenObjectController.AddObject)
	httpRouter.POST("/addUser", userController.AddUser)
	httpRouter.POST("/getUser", userController.GetUser)

	httpRouter.SERVE(port)

}
