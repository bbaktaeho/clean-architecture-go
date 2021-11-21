package controllers

import (
	"clean-architecture-go/dto"
	"clean-architecture-go/entities"
	"clean-architecture-go/helpers"
	"clean-architecture-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService,
		jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authRequest := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authRequest.(entities.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.Id, 10))
		v.Token = generatedToken
		response := helpers.BuildResponse(true, "ok", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Please check again your credential", "Invaild Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.Id, 10))
		createdUser.Token = token
		response := helpers.BuildResponse(true, "ok", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
