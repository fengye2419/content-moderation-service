package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"github.com/urfave/cli/v2"

	_ "github.com/fengye2419/content-moderation-service/docs" // 引入生成的 Swagger 文档
	"github.com/fengye2419/content-moderation-service/routers"
)

// @title           Content Moderation Service
// @version         1.0
// @description     This is content moderation service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app := &cli.App{
		Name:  "content-moderation-service",
		Usage: "Run the content moderation service",
		Action: func(c *cli.Context) error {
			overseer.Run(overseer.Config{
				Program: prog,
				Address: ":8080",
				Fetcher: &fetcher.File{Path: os.Args[0]},
			})
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func prog(state overseer.State) {
	// Initialize the router
	r := routers.InitializeRouter()

	// Start the server
	log.Println("Content Moderation Service is running on", state.Listener.Addr())
	http.Serve(state.Listener, r)
}
