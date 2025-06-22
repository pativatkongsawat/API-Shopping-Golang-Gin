package users_test

import (
	"bytes"
	"encoding/json"
	controllerauth "go_gin/controller/ControllerAuth"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var router *gin.Engine

// Struct สำหรับอ่าน payload
type UserPayloads struct {
	ValidUser        map[string]interface{} `json:"valid_user"`
	InvalidEmail     map[string]interface{} `json:"invalid_email"`
	ShortPassword    map[string]interface{} `json:"short_password"`
	InvalidFirstname map[string]interface{} `json:"invalid_firstname"`
}

func setupRouter() {
	router = gin.Default()
	router.POST("/auth/register", controllerauth.Register)
}

func loadPayload(t *testing.T) UserPayloads {
	data, err := os.ReadFile("tests/user_payload.json")
	if err != nil {
		t.Fatalf("อ่าน testdata ไม่ได้: %v", err)
	}
	var payloads UserPayloads
	err = json.Unmarshal(data, &payloads)
	if err != nil {
		t.Fatalf("แปลง json ไม่ได้: %v", err)
	}
	return payloads
}

func TestRegisterUser(t *testing.T) {
	setupRouter()
	payloads := loadPayload(t)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
	}{
		{"Valid User", payloads.ValidUser, http.StatusCreated},
		{"Invalid Email", payloads.InvalidEmail, http.StatusBadRequest},
		{"Short Password", payloads.ShortPassword, http.StatusBadRequest},
		{"Invalid Firstname", payloads.InvalidFirstname, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/users/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
