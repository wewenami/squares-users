package database

import "git.squares.dev/src/database/drivers"

func New(config drivers.Config) (drivers.DataStore, error) {
	switch config.DSName {
	case "postgres":
		return nil, nil
	default:
		return nil, drivers.ErrConfigNotSet
	}
}
