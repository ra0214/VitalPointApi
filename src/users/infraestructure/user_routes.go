package infraestructure

import (
	"vitalPoint/src/users/application"
	"vitalPoint/src/users/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IUser) *gin.Engine {
	r := gin.Default()

	createUser := application.NewCreateUser(repo)
	createUserController := NewCreateUserController(createUser)

	viewUser := application.NewViewUser(repo)
	viewUserController := NewViewUserController(viewUser)

	editUserUseCase := application.NewEditUser(repo)
	editUserController := NewEditUserController(editUserUseCase)

	deleteUserUseCase := application.NewDeleteUser(repo)
	deleteUserController := NewDeleteUserController(deleteUserUseCase)

	loginUser := application.NewLoginUser(repo)
	loginUserController := NewLoginUserController(loginUser)

	r.POST("/user", createUserController.Execute)
	r.GET("/user", viewUserController.Execute)
	r.PUT("/user/:id", editUserController.Execute)
	r.DELETE("/user/:id", deleteUserController.Execute)
	r.POST("/login", loginUserController.Execute)

	return r
}
