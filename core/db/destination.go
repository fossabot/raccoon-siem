package db

import (
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
	"upper.io/db.v3/lib/sqlbuilder"
)

type DestinationFunctions struct {}

func (r DestinationFunctions) List(query string, qc QueryConfig) ([]destinations.Config, error) {
	destinationEntries := make([]destinations.Config, 0, 0)
	selector := qc.Tx.SelectFrom(destinationTable).OrderBy("name")

	if query != "" {
		selector = selector.Where("name ilike", query + "%")
	}

	var err error
	if qc.Page > 0 {
		err = selector.Paginate(defaultPageSize).Page(qc.Page - 1).All(&destinationEntries)
	} else {
		err = selector.All(&destinationEntries)
	}

	return destinationEntries, err
}

func (r DestinationFunctions) ById(id string, qc QueryConfig) (*destinations.Config, error) {
	configs := make([]destinations.Config, 0, 1)
	selector := qc.Tx.SelectFrom(destinationTable).
		Where("id", id)

	err := selector.All(&configs)

	if err != nil {
		return nil, err
	}

	if len(configs) == 0 {
		return nil, nil
	}

	return &configs[0], nil
}

func (r DestinationFunctions) Exists(config *destinations.Config, id string, qc QueryConfig) (bool, error) {
	configs := make([]destinations.Config, 0, 1)
	selector := qc.Tx.SelectFrom(destinationTable).
		Where("name", config.Name)

	if !IDEmpty(id) {
		selector = selector.And("id <>", id)
	}

	err := selector.All(&configs)

	if err != nil {
		return false, err
	}

	if len(configs) == 0 {
		return false, nil
	}

	return true, nil
}

func (r DestinationFunctions) Create(config *destinations.Config, qc QueryConfig) error {
	inserter := qc.Tx.InsertInto(destinationTable).Values(config)

	var id string
	it := inserter.Returning("id").Iterator()
	if err := it.ScanOne(&id); err != nil {
		return err
	}
	config.Id = id
	return nil
}

func (r DestinationFunctions) Update(id string, config *destinations.Config, qc QueryConfig) error {
	updater := qc.Tx.Update(destinationTable).
		Set(config).
		Where("id", id)

	_, err := updater.Exec()
	return err
}

func (r DestinationFunctions) Delete(name string, tx sqlbuilder.Tx) error {
	return nil
}
