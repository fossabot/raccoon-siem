package core

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"net/http"
)

var (
	df = db.DestinationFunctions{}
)

func readDestinations(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	configs, err := df.List(ctx.Query("query"), qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
	}

	replyJson(ctx, configs)
}

func readDestination(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	config, err := df.ById(ctx.Param("id"), qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
	}
	if config == nil {
		replyError(ctx, http.StatusNotFound, err)
	}

	replyJson(ctx, config)
}

func createDestination(ctx *gin.Context) {
	destination := new(destinations.Config)
	err := unmarshalFromRawData(ctx, destination)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := validateDestination(destination, "", qc); err != nil {
		replyError(ctx, http.StatusBadRequest, err)
		return
	}

	err = df.Create(destination, qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	replyJson(ctx, destination)
}

func updateDestination(ctx *gin.Context) {
	destination := new(destinations.Config)
	err := unmarshalFromRawData(ctx, destination)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	id := ctx.Param("id")
	existingDestination, err := df.ById(id, qc)
	if err != nil {
		replyError(ctx, http.StatusNotFound, err)
		return
	}

	if err := validateDestination(destination, existingDestination.Id, qc); err != nil {
		replyError(ctx, http.StatusBadRequest, err)
		return
	}

	err = df.Update(id, destination, qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	replyJson(ctx, destination)
}

func deleteDestination(ctx *gin.Context) {

}

func validateDestination(destination *destinations.Config, id string, qc db.QueryConfig) error {
	err := destination.Validate()
	if err != nil {
		return err
	}

	df := db.DestinationFunctions{}
	found, err := df.Exists(destination, id, qc)

	if err != nil {
		return err
	}

	if found {
		return errors.New(fmt.Sprintf("destination '%s' already exists", destination.Name))
	}

	return nil
}
