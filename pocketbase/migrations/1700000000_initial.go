package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		// Add your initial schema/data setup here.
		return nil
	}, func(app core.App) error {
		// Roll back your initial schema/data setup here.
		return nil
	})
}
