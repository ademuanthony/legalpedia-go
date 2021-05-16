package main

import (
	"context"

	"github.com/ademuanthony/legalpedia/article"
	"github.com/ademuanthony/legalpedia/dictionary"
	"github.com/ademuanthony/legalpedia/formsandprecedents"
	"github.com/ademuanthony/legalpedia/homepage"
	"github.com/ademuanthony/legalpedia/legalresource"
	"github.com/ademuanthony/legalpedia/maxim"
	"github.com/ademuanthony/legalpedia/postgres"
	"github.com/ademuanthony/legalpedia/rulesofcourt"
	"github.com/ademuanthony/legalpedia/web"
)

func setupModules(ctx context.Context, cfg *config, server *web.Server) error {
	var err error

	pgDb, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DebugLevel == "debug")
	if err != nil {
		return err
	}
	if err = pgDb.CreateTables(ctx); err != nil {
		log.Error("Error creating tables: ", err)
		return err
	}

	_, err = homepage.New(server)
	if err != nil {
		log.Error(err)
		return err
	}

	if err = article.Activate(ctx, pgDb, server); err != nil {
		log.Error(err)
		return err
	}

	if err = rulesofcourt.Activate(ctx, pgDb, server); err != nil {
		log.Error(err)
		return err
	}

	if err = formsandprecedents.Activate(ctx, pgDb, server); err != nil {
		log.Error(err)
		return err
	}

	if err = dictionary.Activate(ctx, pgDb, server); err != nil {
		log.Error(err)
		return err
	}

	if err = maxim.Activate(ctx, pgDb, server); err != nil {
		log.Error(err)
		return err
	}

	if err = legalresource.Activate(ctx, pgDb, server); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
