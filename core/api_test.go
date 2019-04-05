package core

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/core/migrator/assets"
	"github.com/tephrocactus/raccoon-siem/core/migrator/migration"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var client *RestClient

type RestClient struct {
	router *gin.Engine
}

// Function creates new test database. If it already exists we drop it before creation
// Every time all migrations are applied
func prepareTestDatabase(dbHost, dbPort string) error {
	if err := NewUdbConnection(dbHost, dbPort, ""); err != nil {
		return err
	}

	_, err := UDBConn.Exec("drop database if exists raccoon_test; create database raccoon_test;")

	cockroachMigration := migration.CockroachMigration{}
	cockroachMigrationFiles := assets.GetMigrationFiles()
	_, err = cockroachMigration.Run(dbHost, dbPort, "raccoon_test", cockroachMigrationFiles)
	if err != nil {
		return err
	}

	if err := NewUdbConnection(dbHost, dbPort, "raccoon_test"); err != nil {
		return err
	}

	return nil
}

func (r RestClient) performRequest(method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.router.ServeHTTP(w, req)
	return w
}

func (r RestClient) Get(url string) *httptest.ResponseRecorder {
	return r.performRequest("GET", url, nil)
}

func (r RestClient) Post(url string, model interface{}) (*httptest.ResponseRecorder, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return r.performRequest("POST", url, bytes.NewReader(data)), nil
}

func (r RestClient) Put(url string, model interface{}) (*httptest.ResponseRecorder, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return r.performRequest("PUT", url, bytes.NewReader(data)), nil
}

func (r RestClient) Delete(url string) *httptest.ResponseRecorder {
	return r.performRequest("DELETE", url, nil)
}

func TestMain(m *testing.M) {
	err := prepareTestDatabase("127.0.0.1", "26257")
	if err != nil {
		os.Exit(1)
		return
	}

	client = &RestClient{
		getRouter(),
	}

	os.Exit(m.Run())
}
