package db

import (
	"encoding/json"
	"github.com/tephrocactus/raccoon-siem/sdk/connectors"
)

type ConnectorModel struct {
	Id      string             `json:"id,omitempty" db:"id,omitempty"`
	Name    string             `json:"-" db:"name,omitempty"`
	Kind    string             `json:"-" db:"kind,omitempty"`
	Config  *connectors.Config `json:"config,omitempty" db:"-"`
	Payload string             `json:"-" db:"payload,omitempty"`
}

type ConnectorFunctions struct{}

func (r ConnectorFunctions) List(query string, qc QueryConfig) ([]ConnectorModel, error) {
	entries := make([]ConnectorModel, 0, 0)
	selector := qc.Tx.SelectFrom(connectorConfigTable).OrderBy("name")

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

func (r *ConnectorFunctions) ById(id string, qc QueryConfig) (*ConnectorModel, error) {
	entries := make([]ConnectorModel, 0, 1)
	selector := qc.Tx.SelectFrom(connectorConfigTable).
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

func (r ConnectorFunctions) Exists(config *ConnectorModel, id string, qc QueryConfig) (bool, error) {
	configs := make([]ConnectorModel, 0, 1)
	selector := qc.Tx.SelectFrom(connectorConfigTable).
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

func (r *ConnectorModel) Create(qc QueryConfig) error {
	if err := r.dumpConfig(); err != nil {
		return err
	}

	inserter := qc.Tx.InsertInto(connectorConfigTable).Values(r)

	var id string
	it := inserter.Returning("id").Iterator()
	if err := it.ScanOne(&id); err != nil {
		return err
	}
	r.Id = id
	return nil
}

func (r *ConnectorModel) Update(id string, qc QueryConfig) error {
	if err := r.dumpConfig(); err != nil {
		return err
	}

	updater := qc.Tx.Update(connectorConfigTable).
		Set(r).
		Where("id", id)

	_, err := updater.Exec()
	return err
}

func (r *ConnectorModel) Delete(qc QueryConfig) error {
	deleter := qc.Tx.DeleteFrom(connectorConfigTable).Where("id", r.Id)
	_, err := deleter.Exec()
	return err
}

func (r *ConnectorModel) dumpConfig() error {
	bytes, err := json.Marshal(r.Config)

	if err != nil {
		return err
	}

	r.Name = r.Config.Name
	r.Kind = r.Config.Kind
	r.Payload = string(bytes)

	return nil
}

func (r *ConnectorModel) loadConfig() error {
	r.Config = new(connectors.Config)
	return json.Unmarshal([]byte(r.Payload), r.Config)
}
