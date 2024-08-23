package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/BielPinto/curso_go/7-Apis/docs"
	"github.com/BielPinto/curso_go/7-Apis/infra/database"
	"github.com/BielPinto/curso_go/7-Apis/internal/dto"
	"github.com/BielPinto/curso_go/7-Apis/internal/entity"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInteface
}

func NewUserHandler(userDB database.UserInteface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// GetJWT godoc
//
//	@Summary		Get a user JWT
//	@Description	Get a user JWT
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.GetJWTInput	true	"user credentials"
//	@Success		200		{object}	dto.GetJWTOutput
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/users/generate_token [post]
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
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
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

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateUserInput	true	"user request"
//	@Success		201
//	@Failure		500	{object}	Error
//	@Router			/users [post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
