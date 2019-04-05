package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"net/http"
)

func readConnectors(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	df := db.ConnectorFunctions{}
	configs, err := df.List(ctx.Query("query"), qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
	}

	replyJson(ctx, configs)
}

func readConnector(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	df := db.ConnectorFunctions{}
	config, err := df.ById(ctx.Param("id"), qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}
	if config == nil {
		replyError(ctx, http.StatusNotFound, err)
		return
	}

	replyJson(ctx, config)
}

func createConnector(ctx *gin.Context) {
	connectorConfig := new(db.ConnectorModel)
	err := unmarshalFromRawData(ctx, connectorConfig)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := validateConnector(connectorConfig, "", qc); err != nil {
		replyError(ctx, http.StatusBadRequest, err)
		return
	}

	err = connectorConfig.Create(qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	replyJson(ctx, connectorConfig)
}

func updateConnector(ctx *gin.Context) {
	connectorConfig := new(db.ConnectorModel)
	err := unmarshalFromRawData(ctx, connectorConfig)
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
	df := db.ConnectorFunctions{}
	existingConnector, err := df.ById(id, qc)
	if err != nil {
		replyError(ctx, http.StatusNotFound, err)
		return
	}

	if err := validateConnector(connectorConfig, existingConnector.Id, qc); err != nil {
		replyError(ctx, http.StatusBadRequest, err)
		return
	}

	err = connectorConfig.Update(id, qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	replyJson(ctx, connectorConfig)
}

func deleteConnector(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	id := ctx.Param("id")
	df := db.ConnectorFunctions{}
	connectorConfig, err := df.ById(id, qc)
	if err != nil {
		replyError(ctx, http.StatusNotFound, err)
		return
	}

	err = connectorConfig.Delete(qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func validateConnector(connectorConfig *db.ConnectorModel, id string, qc db.QueryConfig) error {
	if connectorConfig.Config == nil {
		return errors.New("connectorConfig config : empty config body")
	}

	err := connectorConfig.Config.Validate()
	if err != nil {
		return err
	}

	df := db.ConnectorFunctions{}
	found, err := df.Exists(connectorConfig, id, qc)

	if err != nil {
		return err
	}

	if found {
		return errors.New(fmt.Sprintf("connectorConfig config : '%s' already exists", connectorConfig.Name))
	}

	return nil
}
