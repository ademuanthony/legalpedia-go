package article

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

func Activate(ctx context.Context, db store, server *web.Server) error {

	m := module{db, server}

	return m.setupHttpRoutes(ctx)
}

func (m *module) setupHttpRoutes(ctx context.Context) error {
	if err := m.server.Templates.AddTemplate("articles"); err != nil {
		return err
	}
	if err := m.server.Templates.AddTemplate("article"); err != nil {
		return err
	}

	// page routes
	m.server.AddRoute("/articles", web.GET, m.articlesPage)
	m.server.AddRoute("/articles/{id}", web.GET, m.articlePage, web.IDParamCtx)

	// api routes
	m.server.AddAPIRoute("articles/Filter", web.GET, m.articlesEndpoint)

	m.server.AddMenuItem(web.MenuItem{
		Href:      "/articles",
		HyperText: "Articles",
		Info:      "Artcles",
		Icon:      "feather icon-package",
	})

	return nil
}
