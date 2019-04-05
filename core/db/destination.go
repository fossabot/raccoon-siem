package db

import (
	"encoding/json"
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
)

type DestinationModel struct {
	Id      string               `json:"id,omitempty" db:"id,omitempty"`
	Name    string               `json:"-" db:"name,omitempty"`
	Config  *destinations.Config `json:"config,omitempty" db:"-"`
	Payload string               `json:"-" db:"payload,omitempty"`
}

type DestinationFunctions struct{}

func (r DestinationFunctions) List(query string, qc QueryConfig) ([]DestinationModel, error) {
	destinationEntries := make([]DestinationModel, 0, 0)
	selector := qc.Tx.SelectFrom(destinationTable).OrderBy("name")

	if query != "" {
		selector = selector.Where("name ilike", query+"%")
	}

	var err error
	if qc.Page > 0 {
		err = selector.Paginate(defaultPageSize).Page(qc.Page - 1).All(&destinationEntries)
	} else {
		err = selector.All(&destinationEntries)
	}

	for i := range destinationEntries {
		destinationModel := &destinationEntries[i]
		if err := destinationModel.loadConfig(); err != nil {
			return nil, err
		}
	}

	return destinationEntries, err
}

func (r *DestinationFunctions) ById(id string, qc QueryConfig) (*DestinationModel, error) {
	configs := make([]DestinationModel, 0, 1)
	selector := qc.Tx.SelectFrom(destinationTable).
		Where("id", id)

	err := selector.All(&configs)

	if err != nil {
		return nil, err
	}

	if len(configs) == 0 {
		return nil, nil
	}

	config := &configs[0]
	if err := config.loadConfig(); err != nil {
		return nil, err
	}

	return config, nil
}

func (r DestinationFunctions) Exists(config *DestinationModel, id string, qc QueryConfig) (bool, error) {
	configs := make([]DestinationModel, 0, 1)
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

func (r *DestinationModel) Create(qc QueryConfig) error {
	if err := r.dumpConfig(); err != nil {
		return err
	}

	inserter := qc.Tx.InsertInto(destinationTable).Values(r)

	var id string
	it := inserter.Returning("id").Iterator()
	if err := it.ScanOne(&id); err != nil {
		return err
	}
	r.Id = id
	return nil
}

func (r *DestinationModel) Update(id string, qc QueryConfig) error {
	if err := r.dumpConfig(); err != nil {
		return err
	}

	updater := qc.Tx.Update(destinationTable).
		Set(r).
		Where("id", id)

	_, err := updater.Exec()
	return err
}

func (r *DestinationModel) Delete(qc QueryConfig) error {
	deleter := qc.Tx.DeleteFrom(destinationTable).Where("id", r.Id)
	_, err := deleter.Exec()
	return err
}

func (r *DestinationModel) dumpConfig() error {
	bytes, err := json.Marshal(r.Config)

	if err != nil {
		return err
	}

	r.Name = r.Config.Name
	r.Payload = string(bytes)

	return nil
}

func (r *DestinationModel) loadConfig() error {
	r.Config = new(destinations.Config)
	return json.Unmarshal([]byte(r.Payload), r.Config)
}
