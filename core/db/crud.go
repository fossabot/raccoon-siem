package db

import "upper.io/db.v3/lib/sqlbuilder"

func create(tableName string, values interface{}, qc QueryConfig) (string, error) {
	inserter := qc.Tx.InsertInto(tableName).Values(values)

	var id string
	it := inserter.Returning("id").Iterator()
	if err := it.ScanOne(&id); err != nil {
		return "", err
	}

	return id, nil
}

func update(tableName string, id string, values interface{}, qc QueryConfig) error {
	updater := qc.Tx.Update(tableName).
		Set(values).
		Where("id", id)

	_, err := updater.Exec()
	return err
}

func read(tableName string, qc QueryConfig, dest interface{}, before func(selector sqlbuilder.Selector) sqlbuilder.Selector) error {
	selector := qc.Tx.SelectFrom(tableName).OrderBy(qc.OrderBy...)

	selector = before(selector)

	var err error
	if qc.Page > 0 {
		err = selector.Paginate(defaultPageSize).Page(qc.Page - 1).All(dest)
	} else {
		err = selector.All(dest)
	}

	return err
}
