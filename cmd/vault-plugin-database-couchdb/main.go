package main

import (
	"log"
	"os"

	couchdb "github.com/cosminolariu/couchdb"
	dbplugin "github.com/hashicorp/vault/sdk/database/dbplugin/v5"
)

func main() {
	if err := Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// Run starts serving the plugin
func Run() error {
	db, err := couchdb.New()
	if err != nil {
		return err
	}
	dbplugin.Serve(db.(dbplugin.Database))
	return nil
}
