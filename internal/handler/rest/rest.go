package rest

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/service"
	"INTERN_BCC/pkg/helper"
	"INTERN_BCC/pkg/middleware"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Timeout())

	v1 := r.router.Group("/api/v1")

	v1.GET("/health-check", healthCheck)

	v1.GET("/time-out", testTimeout)

	v1.GET("/login-user", r.middleware.AuthenticateUser, getLoginUser)

	v1.POST("/register", r.Register)
	v1.POST("/login", r.Login)

	// user.POST("/profile/upload", r.middleware.AuthenticateUser, r.UploadPhoto)

	// City
	cityGroup := v1.Group("/city")
	cityGroup.GET("/get-city/:id", r.GetCity)
	cityGroup.GET("/get-city/all-of-the-cities", r.GetAllCity)

	// Place
	placeGroup := v1.Group("/place")
	placeGroup.POST("/create-place", r.CreatePlace)
	placeGroup.GET("/get-place/id", r.GetPlaceByID)
	placeGroup.GET("/get-place/all-of-the-places", r.GetAllPlace)

	// Culinary
	CulinaryGroup := v1.Group("/culinary")
	CulinaryGroup.POST("/create-culinary", r.CreateCulinary)
	CulinaryGroup.GET("/get-culinary/:id", r.GetCulinaryByID)
	CulinaryGroup.GET("/get-culinary/all-of-the-culinaries", r.GetAllCulinary)
	// v1.GET("/culinary/search", r.SearchCulinary)

	// Ticket
	v1.POST("/ticket", r.CreateTicket)
	v1.GET("/ticket/:id", r.GetTicketByID)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    r.router.Run(fmt.Sprintf(":%s", port))
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		helper.Error(ctx, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	helper.Success(ctx, http.StatusOK, "get login user", user.(entity.User))
}
