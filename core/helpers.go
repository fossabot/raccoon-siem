package core

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"net/http"
	"strconv"
	"strings"
	"upper.io/db.v3/lib/sqlbuilder"
)

func reply(ctx *gin.Context, err error, results ...[]byte) {
	if err != nil {
		_ = ctx.Error(err)
		ctx.String(http.StatusInternalServerError, "%v", ctx.Errors)
		ctx.Abort()
	} else if len(results) > 0 {
		ctx.String(http.StatusOK, "%s", results[0])
	}
}

func replyJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func replyError(ctx *gin.Context, code int, err error) {
	if err != nil {
		_ = ctx.Error(err)
	}
	ctx.String(code, "%s", err)
	ctx.Abort()
}

func getPageParam(ctx *gin.Context) uint {
	page := ctx.Query("page")
	result, err := strconv.Atoi(page)
	if err != nil {
		return 0
	}
	return uint(result)
}

func getOrderByParam(ctx *gin.Context) []interface{} {
	orderParam := ctx.GetString("order")
	if orderParam == "" {
		return nil
	}

	parts := strings.Split(orderParam, ",")
	orderBy := make([]interface{}, len(parts), len(parts))
	for i := range parts {
		orderBy[i] = parts[i]
	}

	return orderBy
}

func getQc(ctx *gin.Context) (db.QueryConfig, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return db.QueryConfig{}, err
	}

	return db.QueryConfig{
		Tx:      tx,
		Page:    getPageParam(ctx),
		OrderBy: getOrderByParam(ctx),
	}, nil
}

func setTx(ctx *gin.Context, tx sqlbuilder.Tx) {
	ctx.Set("tx", tx)
}

func getTx(ctx *gin.Context) (sqlbuilder.Tx, error) {
	value, found := ctx.Get("tx")
	if !found {
		return nil, errors.New("transaction not found in context")
	}

	tx, success := value.(sqlbuilder.Tx)
	if !success {
		return nil, errors.New("wrong value for transaction in context")
	}
	return tx, nil
}

func unmarshalFromRawData(ctx *gin.Context, dest interface{}) error {
	data, err := ctx.GetRawData()
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return err
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		replyError(ctx, http.StatusInternalServerError, err)
		return err
	}

	return nil
}

//func readNormalizersByIDs(ids []string, root bool, uniqueIDs map[string]bool, tx *bolt.Tx) ([]normalizers.Config, error) {
//	//if uniqueIDs == nil {
//	//	uniqueIDs = make(map[string]bool)
//	//}
//	//
//	//result := make([]normalizers.Config, 0)
//	//b := tx.Bucket(dbBucketNormalizer)
//	//
//	//for _, id := range ids {
//	//	if _, ok := uniqueIDs[id]; ok {
//	//		continue
//	//	}
//	//
//	//	uniqueIDs[id] = true
//	//
//	//	rawSettings := b.Get([]byte(id))
//	//
//	//	if rawSettings == nil {
//	//		return nil, fmt.Errorf("parser '%s' does not exist", id)
//	//	}
//	//
//	//	settings := normalizers.Config{}
//	//
//	//	if err := yaml.Unmarshal(rawSettings, &settings); err != nil {
//	//		return nil, err
//	//	}
//	//
//	//	settings.Root = root
//	//
//	//	result = append(result, settings)
//	//
//	//	var extraNormalizersIDs []string
//	//	for _, m := range settings.Mapping {
//	//		if m.Extra != nil {
//	//			extraNormalizersIDs = append(extraNormalizersIDs, m.Extra.NormalizerName)
//	//		}
//	//	}
//	//
//	//	if len(extraNormalizersIDs) > 0 {
//	//		subSettings, err := readNormalizersByIDs(extraNormalizersIDs, false, uniqueIDs, tx)
//	//		if err != nil {
//	//			return nil, err
//	//		}
//	//		result = append(result, subSettings...)
//	//	}
//	//}
//	//
//	//return result, nil
//}
//
//func readConnectorsByIDs(ids []string, tx *bolt.Tx) ([]connectors.Config, error) {
//	result := make([]connectors.Config, 0)
//	b := tx.Bucket(dbBucketConnector)
//
//	for _, id := range ids {
//		rawSettings := b.Get([]byte(id))
//
//		if rawSettings == nil {
//			return nil, fmt.Errorf("source '%s' does not exist", id)
//		}
//
//		settings := connectors.Config{}
//
//		if err := yaml.Unmarshal(rawSettings, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//	}
//
//	return result, nil
//}
//
//func readDestinationsByIDs(ids []string, tx *bolt.Tx) ([]sdk.DestinationSettings, error) {
//	result := make([]sdk.DestinationSettings, 0)
//	b := tx.Bucket(dbBucketDestination)
//
//	for _, id := range ids {
//		rawSettings := b.Get([]byte(id))
//
//		if rawSettings == nil {
//			return nil, fmt.Errorf("destination '%s' does not exist", id)
//		}
//
//		settings := sdk.DestinationSettings{}
//
//		if err := yaml.Unmarshal(rawSettings, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//	}
//
//	return result, nil
//}
//
//func readCorrelationRulesByIDs(ids []string, tx *bolt.Tx) ([]sdk.CorrelationRuleSettings, error) {
//	result := make([]sdk.CorrelationRuleSettings, 0)
//	b := tx.Bucket(dbBucketCorrelationRule)
//
//	for _, id := range ids {
//		rawSettings := b.Get([]byte(id))
//
//		if rawSettings == nil {
//			return nil, fmt.Errorf("correlation rule '%s' does not exist", id)
//		}
//
//		settings := sdk.CorrelationRuleSettings{}
//
//		if err := yaml.Unmarshal(rawSettings, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//	}
//
//	return result, nil
//}
//
//func readAggregationRulesByIDs(ids []string, tx *bolt.Tx) ([]sdk.AggregationRuleSettings, error) {
//	result := make([]sdk.AggregationRuleSettings, 0)
//	b := tx.Bucket(dbBucketAggregationRule)
//
//	for _, id := range ids {
//		rawSettings := b.Get([]byte(id))
//
//		if rawSettings == nil {
//			return nil, fmt.Errorf("aggregation rule '%s' does not exist", id)
//		}
//
//		settings := sdk.AggregationRuleSettings{}
//
//		if err := yaml.Unmarshal(rawSettings, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//	}
//
//	return result, nil
//}
//
//func readFiltersByIDs(ids []string, uniqueIDs map[string]bool, tx *bolt.Tx) ([]sdk.FilterSettings, error) {
//	if uniqueIDs == nil {
//		uniqueIDs = make(map[string]bool)
//	}
//
//	result := make([]sdk.FilterSettings, 0)
//	b := tx.Bucket(dbBucketFilter)
//
//	for _, id := range ids {
//		if _, ok := uniqueIDs[id]; ok {
//			continue
//		}
//
//		uniqueIDs[id] = true
//
//		rawSettings := b.Get([]byte(id))
//
//		if rawSettings == nil {
//			return nil, fmt.Errorf("filter '%s' does not exist", id)
//		}
//
//		settings := sdk.FilterSettings{}
//
//		if err := yaml.Unmarshal(rawSettings, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//
//		incFilterIDs := make([]string, 0)
//
//		for _, section := range settings.Sections {
//			for _, expr := range section.Expressions {
//				if strings.HasPrefix(expr, sdk.FilterIncludeSymbol) {
//					incFilterIDs = append(incFilterIDs, sdk.GetIncludedFilterName(expr))
//				}
//			}
//		}
//
//		if len(incFilterIDs) > 0 {
//			subSettings, err := readFiltersByIDs(incFilterIDs, uniqueIDs, tx)
//
//			if err != nil {
//				return nil, err
//			}
//
//			result = append(result, subSettings...)
//		}
//	}
//
//	return result, nil
//}
//
//func readAllActiveLists(tx *bolt.Tx) ([]sdk.ActiveListSettings, error) {
//	result := make([]sdk.ActiveListSettings, 0)
//	b := tx.Bucket(dbBucketActiveList)
//	c := b.Cursor()
//
//	for k, v := c.First(); k != nil; k, v = c.Next() {
//		settings := sdk.ActiveListSettings{}
//
//		if err := yaml.Unmarshal(v, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//	}
//
//	return result, nil
//}
//
//func readAllDictionaries(tx *bolt.Tx) ([]sdk.DictionarySettings, error) {
//	result := make([]sdk.DictionarySettings, 0)
//	b := tx.Bucket(dbBucketDictionary)
//	c := b.Cursor()
//
//	for k, v := c.First(); k != nil; k, v = c.Next() {
//		settings := sdk.DictionarySettings{}
//
//		if err := yaml.Unmarshal(v, &settings); err != nil {
//			return nil, err
//		}
//
//		result = append(result, settings)
//	}
//
//	return result, nil
//}
//
//func unmarshalAndGetID(dst sdk.IBaseResource, data []byte) (string, error) {
//	err := yaml.Unmarshal(data, dst)
//
//	if err != nil {
//		return "", err
//	}
//
//	if dst.ID() == "" {
//		return "", errors.New("resource ID can not be empty")
//	}
//
//	return dst.ID(), nil
//}
