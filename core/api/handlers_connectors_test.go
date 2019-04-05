package api

import (
	"encoding/json"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func TestConnectorAPI(t *testing.T) {
	testConnectorRead(t)
	config := testConnectorCreate(t)
	testConnectorReadById(t, config)
	testConnectorUpdate(t, config)
	testConnectorDelete(t, config)
}

func testConnectorRead(t *testing.T) {
	response := client.Get("/config/connector/")
	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, response.Body.String(), "[]")
}

func testConnectorReadById(t *testing.T, config *db.ConnectorModel) {
	response := client.Get("/config/connector/" + config.Id)
	assert.Equal(t, response.Code, http.StatusOK)

	result := &db.ConnectorModel{}
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NilError(t, err)

	assert.Equal(t, config.Id, result.Id)
	assert.Equal(t, config.Name, result.Name)
	assert.Equal(t, config.Payload, result.Payload)
}

func testConnectorCreate(t *testing.T) *db.ConnectorModel {
	config := connectors.Config{
		Name:      "errors",
		Kind:      "nats",
		URL:       "127.0.0.1",
		Delimiter: ";",
		Subject:   "errors",
	}
	postModel := db.ConnectorModel{
		Config: &config,
	}
	response, err := client.Post("/config/connector/", postModel)
	assert.NilError(t, err)
	assert.Equal(t, response.Code, http.StatusOK)

	err = json.Unmarshal(response.Body.Bytes(), &postModel)
	assert.NilError(t, err)
	assert.Equal(t, db.IDEmpty(postModel.Id), false)

	return &postModel
}

func testConnectorUpdate(t *testing.T, config *db.ConnectorModel) {
	config.Config.Name = "errors"
	response, err := client.Put("/config/connector/"+config.Id, config)
	assert.NilError(t, err)
	assert.Equal(t, response.Code, http.StatusOK)
}

func testConnectorDelete(t *testing.T, config *db.ConnectorModel) {
	response := client.Delete("/config/connector/" + config.Id)
	assert.Equal(t, response.Code, http.StatusNoContent)
}
