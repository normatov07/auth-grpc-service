package api

import (
	db "auth/pkg/db/sqlc"
	"auth/pkg/pb"
	"auth/util"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) login(ctx *gin.Context) {
	fmt.Println("request keldi aut servicega login!!!")
	usr, err := s.Store.GetUser(ctx, req.Email)
	if err != nil {
		http.
	}

	if err := util.CheckHashPassword(req.Password, usr.PasswordHash); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "email or password went wrong",
			Token:  "",
		}, nil
	}
	token, err := s.tokenMaker.CreateToken(usr.Email, s.Config.AccessTokenDuration)
	if err != nil {
		fmt.Errorf("failed creating token: %v", err)
	}
	usr.LoginAt = time.Now()
	para := db.UpdateUserParams{
		ID:      usr.ID,
		LoginAt: usr.LoginAt,
	}
	err = s.Store.UpdateUser(context.Background(), para)

	if err != nil {
		fmt.Errorf("error update login at data: %v", err)
	}
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Error:  "",
		Token:  token,
		Data: &pb.Data{
			Id:        usr.ID,
			Email:     usr.Email,
			RoleId:    usr.RoleID,
			AccountId: usr.AccountID.Int64,
			Active:    usr.Active,
			LoginAt:   timestamppb.New(usr.LoginAt),
			CreatedAt: timestamppb.New(usr.CreatedAt),
		},
	}, nil
}

func (s *Server) register(ctx *gin.Context) {
	fmt.Println()
}

func (s *Server) validate(ctx *gin.Context) {
	fmt.Println()
}
