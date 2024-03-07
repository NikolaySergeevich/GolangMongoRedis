package commands

import (
	"redis-cache/database"
)

type Repository interface {
	AddCommand(command database.Command) error
	FindByCommand() (database.Command, error)
}
