/*package main

import (
	"Project_2021_PSRS/routes"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

)

// TestLogin OK
func TestLogin(t *testing.T) {
	router := routes.GenerateToken

	w := httptest.NewRecorder()

	pt := map[string]interface{}{
		"username": "test1",
		"password": "...",
	}
	b, _ := json.Marshal(pt)

	request, _ := http.NewRequest("POST", "/login", bytes.NewReader(b))
	router.ServeHTTP(w, request)

	assert.Equal(t, 401, w.Code)

	pt = map[string]interface{}{
		"email":          "desenvolvimento@ufp.edu.pt",
		"badKeyPassword": "badpassword",
	}
	b, _ = json.Marshal(pt)

	request, _ = http.NewRequest("POST", "/login", bytes.NewReader(b))
	router.ServeHTTP(w, request)


	assert.Equal(t, 401, w.Code)

	pt = map[string]interface{}{
		// no keys
	}
	b, _ = json.Marshal(pt)

	request, _ = http.NewRequest("POST", "/login", bytes.NewReader(b))
	router.ServeHTTP(w, request)

	assert.Equal(t, 401, w.Code)
}

*/