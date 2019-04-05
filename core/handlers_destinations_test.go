package core

import (
	"encoding/json"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func TestDestinationAPI(t *testing.T) {
	testDestinationRead(t)
	config := testDestinationCreate(t)
	testDestinationReadById(t, config)
	testDestinationUpdate(t, config)
	testDestinationDelete(t, config)
}

func testDestinationRead(t *testing.T) {
	response := client.Get("/config/destination/")
	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, response.Body.String(), "[]")
}

func testDestinationReadById(t *testing.T, config *db.DestinationModel) {
	response := client.Get("/config/destination/" + config.Id)
	assert.Equal(t, response.Code, http.StatusOK)

	result := &db.DestinationModel{}
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NilError(t, err)

	assert.Equal(t, config.Id, result.Id)
	assert.Equal(t, config.Name, result.Name)
	assert.Equal(t, config.Payload, result.Payload)
}

func testDestinationCreate(t *testing.T) *db.DestinationModel {
	config := destinations.Config{
		Name:  "elastic",
		Kind:  "elastic",
		Index: "index-places",
		URL:   "-",
	}
	postModel := db.DestinationModel{
		Config: &config,
	}
	response, err := client.Post("/config/destination/", postModel)
	assert.NilError(t, err)
	assert.Equal(t, response.Code, http.StatusOK)

	err = json.Unmarshal(response.Body.Bytes(), &postModel)
	assert.NilError(t, err)
	assert.Equal(t, db.IDEmpty(postModel.Id), false)

	return &postModel
}

func testDestinationUpdate(t *testing.T, config *db.DestinationModel) {
	config.Config.Name = "db"
	response, err := client.Put("/config/destination/"+config.Id, config)
	assert.NilError(t, err)
	assert.Equal(t, response.Code, http.StatusOK)
}

func testDestinationDelete(t *testing.T, config *db.DestinationModel) {
	response := client.Delete("/config/destination/" + config.Id)
	assert.Equal(t, response.Code, http.StatusNoContent)
}
