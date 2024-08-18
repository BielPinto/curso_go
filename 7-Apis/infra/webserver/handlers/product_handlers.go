package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BielPinto/curso_go/7-Apis/infra/database"
	"github.com/BielPinto/curso_go/7-Apis/internal/dto"
	"github.com/BielPinto/curso_go/7-Apis/internal/entity"
)

type ProductHandler struct {
	ProductDB database.ProductIterface
}

func NewProductHandler(db database.ProductIterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product dto.CreatProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
