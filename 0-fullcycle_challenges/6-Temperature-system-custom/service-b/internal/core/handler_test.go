package core

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/dto"
)

func TestSearchCEPHandler_Success(t *testing.T) {
	// backup and restore
	oldSearch := searchViaCep
	oldGetTemp := getTemperature
	defer func() {
		searchViaCep = oldSearch
		getTemperature = oldGetTemp
	}()

	searchViaCep = func(ctx context.Context, cep string) (dto.GetViacepApi, error) {
		return dto.GetViacepApi{Localidade: "Camaragibe"}, nil
	}
	getTemperature = func(ctx context.Context, city string) (float64, error) {
		return 28.5, nil
	}

	req := httptest.NewRequest(http.MethodGet, "/?cep=54786625", nil)
	rr := httptest.NewRecorder()
	SearchCEPHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]float64
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	if got := resp["temp_C"]; got != 28.5 {
		t.Fatalf("expected temp_C 28.5, got %v", got)
	}
	expectedF := 28.5*1.8 + 32
	if got := resp["temp_F"]; !approxEqual(got, expectedF) {
		t.Fatalf("expected temp_F %v, got %v", expectedF, got)
	}
	expectedK := 28.5 + 273
	if got := resp["temp_K"]; !approxEqual(got, expectedK) {
		t.Fatalf("expected temp_K %v, got %v", expectedK, got)
	}
}

func approxEqual(a, b float64) bool {
	const eps = 1e-6
	if a > b {
		return a-b < eps
	}
	return b-a < eps
}

func TestSearchCEPHandler_InvalidCEP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?cep=123", nil)
	rr := httptest.NewRecorder()
	SearchCEPHandler(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected status 422, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "invalid zipcode") {
		t.Fatalf("expected body to contain 'invalid zipcode', got %q", rr.Body.String())
	}
}

func TestSearchCEPHandler_CEPNotFound(t *testing.T) {
	oldSearch := searchViaCep
	defer func() { searchViaCep = oldSearch }()

	searchViaCep = func(ctx context.Context, cep string) (dto.GetViacepApi, error) {
		return dto.GetViacepApi{}, ErrCepNotFound
	}

	req := httptest.NewRequest(http.MethodGet, "/?cep=54786625", nil)
	rr := httptest.NewRecorder()
	SearchCEPHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "can not find zipcode") {
		t.Fatalf("expected body to contain 'can not find zipcode', got %q", rr.Body.String())
	}
}
