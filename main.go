package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB()

	// バリデーションのインジェクション
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	// reposirotyのインジェクション
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	// usecasのインジェクション
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	// Controllerのインジェクション
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	// Controllerの実体を渡す
	e := router.NewRouter(userController,taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
