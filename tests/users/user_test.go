package users_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_gin/database"
	"go_gin/routes"
	"log"

	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func TestMain(m *testing.M) {

	err := godotenv.Load("../../.env.dev") 
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectMySQL()
	code := m.Run()
	os.Exit(code)
}

func setupRouter() {
	router = gin.Default()
	routes.InitRoutes(router)
}

func loadPayload(t *testing.T) []map[string]interface{} {
	data, err := os.ReadFile("../user_payload.json")
	if err != nil {
		t.Fatalf("อ่าน testdata ไม่ได้: %v", err)
	}

	var payloads []map[string]interface{}
	err = json.Unmarshal(data, &payloads)
	if err != nil {
		t.Fatalf("แปลง json ไม่ได้: %v", err)
	}
	return payloads
}

func TestRegisterUser(t *testing.T) {
	setupRouter()
	payloads := loadPayload(t)

	wantStatuses := []int{
		http.StatusCreated,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
	}

	if len(payloads) != len(wantStatuses) {
		t.Fatalf("จำนวน payload และ expected status ไม่ตรงกัน")
	}

	for i, payload := range payloads {
		t.Run(fmt.Sprintf("TestCase #%d", i+1), func(t *testing.T) {
			body, _ := json.Marshal([]map[string]interface{}{payload})

			req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, wantStatuses[i], w.Code)
		})
	}
}
