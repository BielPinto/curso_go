package main

import (
	"log"
	"net/http"

	"github.com/BielPinto/curso_go/7-Apis/configs"
	"github.com/BielPinto/curso_go/7-Apis/infra/database"
	"github.com/BielPinto/curso_go/7-Apis/infra/webserver/handlers"
	"github.com/BielPinto/curso_go/7-Apis/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title		Go Expert Api Example
// @version		1.0
// @decription Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name  Gabriel Pinto
// @contact.url 	htto://www.fullcycle.com.br
// @contact.email gabrielp.roch@hotail.com

// @license.name	Full Cycle License
// @license.url		http://www.fullcycle.com.br

// @host	localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Autorization

func main() {

	configs, err := configs.LoadConfig(".")
	println(configs.DBDriver)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperiesIn", configs.JwtExperiesIn))
	// r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)

	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
