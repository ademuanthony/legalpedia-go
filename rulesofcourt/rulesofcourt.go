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
	if err := m.server.Templates.AddTemplate("rules/rules"); err != nil {
		return err
	}
	if err := m.server.Templates.AddTemplate("rules/rule"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/rules/state", web.GET, m.stateRulesPage)
	m.server.AddRoute("/rules/others", web.GET, m.otherRulesPage)
	m.server.AddRoute("/rules/read/{id}", web.GET, m.readRule, web.IDParamCtx)

	// api routes
	// m.server.AddAPIRoute("articles/Filter", web.GET, m.articlesEndpoint)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/rules/state",
		HyperText: "Rules of Court",
		Info:      "Rules of Court",
		Icon:      "feather icon-package",
	})
	m.server.AddMenuItem(web.MenuItem{
		Href:      "/rules/others",
		HyperText: "Other Rules of Court",
		Info:      "Other Rules of Court",
		Icon:      "feather icon-package",
	})

	return nil
}
