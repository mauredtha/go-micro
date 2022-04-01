package usecases

import (
	"database/sql"
	"errors"
	"log"
	"microservices/libraries/api"
	"microservices/payloads/request"
	"microservices/payloads/response"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// UserUsecase struct
type UserUsecase struct {
	Log *log.Logger
	Db  *sql.DB
}

// Create new user
func (u *UserUsecase) Create(r *http.Request) (response.UserResponse, error) {
	var userRequest request.NewUserRequest
	var res response.UserResponse

	err := api.Decode(r, &userRequest)
	if err != nil {
		u.Log.Printf("error decode user: %s", err)
		return res, err
	}

	if userRequest.Password != userRequest.RePassword {
		err = api.ErrBadRequest(errors.New("Password not match"), "")
		u.Log.Printf("error : %s", err)
		return res, err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("error generate password: %s", err)
		return res, err
	}

	userRequest.Password = string(pass)

	user := userRequest.Transform()

	err = user.Create(u.Db)
	if err != nil {
		u.Log.Printf("error call create user: %s", err)
		return res, err
	}

	res.Transform(*user)
	return res, nil
}
