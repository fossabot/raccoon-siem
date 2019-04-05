package db

import (
	"encoding/json"
	"github.com/tephrocactus/raccoon-siem/sdk/dictionaries"
)

type DictionaryModel struct {
	Id      string               `json:"id,omitempty" db:"id,omitempty"`
	Name    string               `json:"-" db:"name,omitempty"`
	Config  *dictionaries.Config `json:"config,omitempty" db:"-"`
	Payload string               `json:"-" db:"payload,omitempty"`
}

type DictionaryFunctions struct{}

func (r DictionaryFunctions) List(query string, qc QueryConfig) ([]DictionaryModel, error) {
	entries := make([]DictionaryModel, 0, 0)
	selector := qc.Tx.SelectFrom(dictionaryConfigTable).OrderBy("name")

	if query != "" {
		selector = selector.Where("name ilike", query+"%")
	}

	var err error
	if qc.Page > 0 {
		err = selector.Paginate(defaultPageSize).Page(qc.Page - 1).All(&entries)
	} else {
		err = selector.All(&entries)
	}

	for i := range entries {
		model := &entries[i]
		if err := model.loadConfig(); err != nil {
			return nil, err
		}
	}

	return entries, err
}

func (r *DictionaryFunctions) ById(id string, qc QueryConfig) (*DictionaryModel, error) {
	entries := make([]DictionaryModel, 0, 1)
	selector := qc.Tx.SelectFrom(dictionaryConfigTable).
		Where("id", id)

	err := selector.All(&entries)

	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, nil
	}

	config := &entries[0]
	if err := config.loadConfig(); err != nil {
		return nil, err
	}

	return config, nil
}

func (r DictionaryFunctions) Exists(config *DictionaryModel, id string, qc QueryConfig) (bool, error) {
	configs := make([]DictionaryModel, 0, 1)
	selector := qc.Tx.SelectFrom(dictionaryConfigTable).
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

func (r *DictionaryModel) Create(qc QueryConfig) error {
	if err := r.dumpConfig(); err != nil {
		return err
	}

	inserter := qc.Tx.InsertInto(dictionaryConfigTable).Values(r)

	var id string
	it := inserter.Returning("id").Iterator()
	if err := it.ScanOne(&id); err != nil {
		return err
	}
	r.Id = id
	return nil
}

func (r *DictionaryModel) Update(id string, qc QueryConfig) error {
	if err := r.dumpConfig(); err != nil {
		return err
	}

	updater := qc.Tx.Update(dictionaryConfigTable).
		Set(r).
		Where("id", id)

	_, err := updater.Exec()
	return err
}

func (r *DictionaryModel) Delete(qc QueryConfig) error {
	deleter := qc.Tx.DeleteFrom(dictionaryConfigTable).Where("id", r.Id)
	_, err := deleter.Exec()
	return err
}

func (r *DictionaryModel) dumpConfig() error {
	bytes, err := json.Marshal(r.Config)

	if err != nil {
		return err
	}

	r.Name = r.Config.Name
	r.Payload = string(bytes)

	return nil
}

func (r *DictionaryModel) loadConfig() error {
	r.Config = new(dictionaries.Config)
	return json.Unmarshal([]byte(r.Payload), r.Config)
}
