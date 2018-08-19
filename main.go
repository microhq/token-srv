package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/microhq/token-srv/db"
	"github.com/microhq/token-srv/handler"
	"github.com/microhq/token-srv/proto/record"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.token"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/token",
			},
		),
		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()

	db.Init()

	record.RegisterRecordHandler(service.Server(), new(handler.Record))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
