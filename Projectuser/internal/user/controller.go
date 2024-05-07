package user

import (
	"context"
	"encoding/json"
	"fmt"
	"go-fundamental/projectGO/Projectuser/internal/domain"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		GetAll Controller
	}
	CreateReq struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllUsers(ctx, s, w)
		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var user domain.User
			if err := decode.Decode(&user); err != nil {
				MsgResponse(w, http.StatusBadRequest, "Bad Request")
				return
			}
			PostUser(ctx, s, w, user)

		default:
			MsgResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
			return
		}

	}
}

func PostUser(ctx context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(CreateReq)

	//validar si los datos son invalidos
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	s.CreateUser(ctx, req.FirstName, req.LastName, req.Email)
	MsgResponse(w, http.StatusCreated, "User Created")

}

func GetAllUsers(ctx context.Context, s Service, w http.ResponseWriter) {
	users, err := s.GetAllUsers(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	DataResponse(w, http.StatusOK, users)

}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status":%d, "data":%s}`, status, value)
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status":%d, "message":"%s"}`, status, message)
}
