package api

import (
	db "auth/pkg/db/sqlc"
	"auth/pkg/pb"
	"auth/token"
	"auth/util"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
)

// Server to HTTP request to our banking service
type Server struct {
	*pb.UnimplementedAuthServceServer
	Config     util.Config
	Store      db.SQLStore
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates new HTTH server and setup routing.
func NewServer(config util.Config, store db.SQLStore) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can not cretae token: %w", err)
	}
	server := &Server{
		Config:     config,
		Store:      store,
		tokenMaker: tokenMaker,
	}
	setUpRouter(server)

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }
	return server, nil
}

func setUpRouter(server *Server) {
	router := gin.Default()

	router.POST("/login", server.login)
	router.POST("/register", server.register)
	router.POST("/validate", server.validate)

	server.router = router
}

// Start run HTTP server on a specific address
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

// // Render error response
// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }
