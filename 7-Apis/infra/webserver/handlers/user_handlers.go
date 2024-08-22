package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BielPinto/curso_go/7-Apis/infra/database"
	"github.com/BielPinto/curso_go/7-Apis/internal/dto"
	"github.com/BielPinto/curso_go/7-Apis/internal/entity"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInteface
}

func NewUserHandler(userDB database.UserInteface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExperiesIn := r.Context().Value("JwtExperiesIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub":  u.ID.String(),
		"exp":  time.Now().Add(time.Second * time.Duration(jwtExperiesIn)).Unix(),
		"name": u.Name,
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}