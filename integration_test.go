//go:build integration
// +build integration

package main

import (
	//pg "github.com/go-pg/pg/v10"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shankar524/company-management-service/app"
	"github.com/shankar524/company-management-service/app/configs"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	//"strings"
	"testing"
)

var (
	server     *gin.Engine
	testServer *httptest.Server
)

func setUp() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server = app.SetupEngine(configs.GetConfig())
	testServer = httptest.NewServer(app.SetupEngine(configs.GetConfig()))
}

func cleanUp() {
	//server.Close()
	testServer.Close()
}

func TestMain(m *testing.M) {
	setUp()
	// exec test and this returns an exit code to pass to os
	returnCode := m.Run()

	cleanUp()
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(returnCode)
}

func TestHealthCheckEndpoint(t *testing.T) {

	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		log.Println("Running server")

		resp, err := http.Get(fmt.Sprintf("%s/health", testServer.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, `{"message":"ok"}`, string(body))
	})
}
