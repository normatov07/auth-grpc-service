package api

import (
	db "auth/pkg/db/sqlc"
	"auth/pkg/pb"
	"auth/util"
	"context"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("request keldi aut servicega register!!!")
	_, err := s.Store.GetUser(context.Background(), req.Email)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email is already exists",
		}, nil
	}
	para := db.CreateUserParams{}
	para.Email = req.Email
	para.PasswordHash = util.HashedPassword(req.Password)
	para.Active = true
	para.RoleID = 1

	us, err := s.Store.CreateUser(context.Background(), para)
	if err != nil || us == (db.User{}) {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  fmt.Sprintf("sorry error on writing data to database: %v", err),
		}, nil
	}
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Error:  "",
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("request keldi aut servicega login!!!")
	usr, err := s.Store.GetUser(context.Background(), req.Email)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "email or password went wrong",
			Token:  "",
		}, nil
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

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, err := s.tokenMaker.VerifyToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}
