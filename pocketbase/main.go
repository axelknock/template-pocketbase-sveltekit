package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/your-org/template-pocketbase-sveltekit/pocketbase/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	// PocketBase docs reference:
	// https://pocketbase.io/docs/go-migrations/
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:          "migrations",
		Automigrate:  isGoRun,
		TemplateLang: migratecmd.TemplateLangGo,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
