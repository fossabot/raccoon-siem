package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"net/http"
)

func readDictionaries(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	df := db.DictionaryFunctions{}
	configs, err := df.List(ctx.Query("query"), qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
	}

	replyJson(ctx, configs)
}

func readDictionary(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	df := db.DictionaryFunctions{}
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

func createDictionary(ctx *gin.Context) {
	dictionaryConfig := new(db.DictionaryModel)
	err := unmarshalFromRawData(ctx, dictionaryConfig)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := validateDictionary(dictionaryConfig, "", qc); err != nil {
		replyError(ctx, http.StatusBadRequest, err)
		return
	}

	err = dictionaryConfig.Create(qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	replyJson(ctx, dictionaryConfig)
}

func updateDictionary(ctx *gin.Context) {
	dictionaryConfig := new(db.DictionaryModel)
	err := unmarshalFromRawData(ctx, dictionaryConfig)
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
	df := db.DictionaryFunctions{}
	existingDictionary, err := df.ById(id, qc)
	if err != nil {
		replyError(ctx, http.StatusNotFound, err)
		return
	}

	if err := validateDictionary(dictionaryConfig, existingDictionary.Id, qc); err != nil {
		replyError(ctx, http.StatusBadRequest, err)
		return
	}

	err = dictionaryConfig.Update(id, qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	replyJson(ctx, dictionaryConfig)
}

func deleteDictionary(ctx *gin.Context) {
	qc, err := getQc(ctx)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	id := ctx.Param("id")
	df := db.DictionaryFunctions{}
	dictionaryConfig, err := df.ById(id, qc)
	if err != nil {
		replyError(ctx, http.StatusNotFound, err)
		return
	}

	err = dictionaryConfig.Delete(qc)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func validateDictionary(dictionaryConfig *db.DictionaryModel, id string, qc db.QueryConfig) error {
	if dictionaryConfig.Config == nil {
		return errors.New("dictionaryConfig config : empty config body")
	}

	err := dictionaryConfig.Config.Validate()
	if err != nil {
		return err
	}

	df := db.DictionaryFunctions{}
	found, err := df.Exists(dictionaryConfig, id, qc)

	if err != nil {
		return err
	}

	if found {
		return errors.New(fmt.Sprintf("dictionaryConfig config : '%s' already exists", dictionaryConfig.Name))
	}

	return nil
}
