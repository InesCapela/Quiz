package main
/*
import (
	"Project_2021_PSRS/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoginHE OK
func TestLoginHE(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()

	pt := map[string]interface{}{
		"email":    "desenvolvimento@ufp.edu.pt",
		"password": "badpassword",
	}
	b, _ := json.Marshal(pt)

	request, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewReader(b))
	router.ServeHTTP(w, request)

	assert.Equal(t, 401, w.Code)

	pt = map[string]interface{}{
		"email":          "desenvolvimento@ufp.edu.pt",
		"badKeyPassword": "badpassword",
	}
	b, _ = json.Marshal(pt)

	request, _ = http.NewRequest("POST", "/api/v1/auth/login", bytes.NewReader(b))
	router.ServeHTTP(w, request)

	assert.Equal(t, 401, w.Code)

	pt = map[string]interface{}{
		// no keys
	}
	b, _ = json.Marshal(pt)

	request, _ = http.NewRequest("POST", "/api/v1/auth/login", bytes.NewReader(b))
	router.ServeHTTP(w, request)

	assert.Equal(t, 401, w.Code)
}*/

