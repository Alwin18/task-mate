package main

import (
	"log"
	"os"
	taskmate "task-mate"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	flags "github.com/jessevdk/go-flags"
	"github.com/spf13/viper"

	"task-mate/api/handler"
	swaggerapisrv "task-mate/gen/api"
	"task-mate/gen/api/operations"
	"task-mate/infrastructure/rest"
)

var mainFlags = struct {
	AppConfig string `long:"config" description:"Main application configuration YAML path"`
}{}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	cfg := viper.New()
	cfg.SetConfigFile(".env")
	err := cfg.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed load config : %f", err)
	}

	swaggerSpec, err := loads.Embedded(swaggerapisrv.SwaggerJSON, swaggerapisrv.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTaskMateAPI(swaggerSpec)
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "App Flags",
			LongDescription:  "",
			Options:          &mainFlags,
		},
	}

	server := swaggerapisrv.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Task Mate API"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	rt := taskmate.NewRuntime(cfg)
	h := handler.NewHandler()
	rest.Route(rt, api, h)

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
