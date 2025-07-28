package infraestructure

import (
	"net/http"
	"vitalPoint/src/users/application"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	useCase *application.CreateUser
}

func NewCreateUserController(useCase *application.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

type RequestBody struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (cu_c *CreateUserController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON"})
		return
	}

	// Validar el rol
	if body.Role != "admin" && body.Role != "normal" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rol inválido. Debe ser 'admin' o 'normal'"})
		return
	}

	err := cu_c.useCase.Execute(body.UserName, body.Email, body.Password, body.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado correctamente"})
}
