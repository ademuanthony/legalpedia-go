package rulesofcourt

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

func Activate(ctx context.Context, db store, server *web.Server) error {
	m := module{
		db:     db,
		server: server,
	}

	return m.setupHTTP()
}

func (m *module) setupHTTP() error {
	if err := m.server.Templates.AddTemplate("rules/state"); err != nil {
		return err
	}
	if err := m.server.Templates.AddTemplate("rules/others"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/rules/state", web.GET, m.stateRulesPage)
	// m.server.AddRoute("/articles/{id}", web.GET, m.articlePage, web.IDParamCtx)

	// api routes
	// m.server.AddAPIRoute("articles/Filter", web.GET, m.articlesEndpoint)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/rules/state",
		HyperText: "Rules of Court",
		Info:      "Rules of Court",
		Icon:      "feather icon-package",
	})

	return nil
}
