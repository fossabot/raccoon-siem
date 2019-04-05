package api

import (
	"encoding/json"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func TestDictionaryAPI(t *testing.T) {
	testDictionaryRead(t)
	config := testDictionaryCreate(t)
	testDictionaryReadById(t, config)
	testDictionaryUpdate(t, config)
	testDictionaryDelete(t, config)
}

func testDictionaryRead(t *testing.T) {
	response := client.Get("/config/dictionary/")
	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, response.Body.String(), "[]")
}

func testDictionaryReadById(t *testing.T, config *db.DictionaryModel) {
	response := client.Get("/config/dictionary/" + config.Id)
	assert.Equal(t, response.Code, http.StatusOK)

	result := &db.DictionaryModel{}
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NilError(t, err)

	assert.Equal(t, config.Id, result.Id)
	assert.Equal(t, config.Name, result.Name)
	assert.Equal(t, config.Payload, result.Payload)
}

func testDictionaryCreate(t *testing.T) *db.DictionaryModel {
	data := make(map[string]string)

	data["error"] = "ERROR"
	data["warning"] = "WARNING"

	config := dictionaries.Config{
		Name: "error-codes",
		Data: data,
	}
	postModel := db.DictionaryModel{
		Config: &config,
	}
	response, err := client.Post("/config/dictionary/", postModel)
	assert.NilError(t, err)
	assert.Equal(t, response.Code, http.StatusOK)

	err = json.Unmarshal(response.Body.Bytes(), &postModel)
	assert.NilError(t, err)
	assert.Equal(t, db.IDEmpty(postModel.Id), false)

	return &postModel
}

func testDictionaryUpdate(t *testing.T, config *db.DictionaryModel) {
	config.Config.Name = "errors"
	response, err := client.Put("/config/dictionary/"+config.Id, config)
	assert.NilError(t, err)
	assert.Equal(t, response.Code, http.StatusOK)
}

func testDictionaryDelete(t *testing.T, config *db.DictionaryModel) {
	response := client.Delete("/config/dictionary/" + config.Id)
	assert.Equal(t, response.Code, http.StatusNoContent)
}
