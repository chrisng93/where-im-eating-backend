package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisng93/where-im-eating-backend/api"
	"github.com/chrisng93/where-im-eating-backend/db"
	"github.com/gorilla/handlers"
	flags "github.com/jessevdk/go-flags"
)

type flagOptions struct {
	Port       string `long:"port" description:"The port for the server to run on." default:"8080" required:"false"`
	DBHost     string `long:"db_host" description:"The database host." default:"localhost:27017" required:"false"`
	DBUsername string `long:"db_username" description:"The database user." default:"" required:"false"`
	DBPassword string `long:"db_password" description:"The database user's password." default:"" required:"false"`
	DBName     string `long:"db_name" description:"The database name." default:"" required:"false"`
}

var options flagOptions

func main() {
	options = flagOptions{}
	_, err := flags.Parse(&options)
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	dbOps := db.Init(options.DBUsername, options.DBPassword, options.DBHost, options.DBName)
	router := api.Init(dbOps)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", options.Port), handlers.CORS()(router)))
}
