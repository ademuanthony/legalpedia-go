package article

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

type module struct {
	db     store
	server *web.Server
}

type store interface {
	FetchArticles(ctx context.Context, req web.PagedResultRequest, term string) (*PagedArticleResponse, error)
	FindByID(ctx context.Context, uuid string) (*Article, error)
}

type Article struct {
	UUID      string `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	Title     string `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Content   string `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	VersionNo int    `boil:"version_no" json:"version_no,omitempty" toml:"version_no" yaml:"version_no,omitempty"`
	ID        int    `boil:"Id" json:"Id,omitempty" toml:"Id" yaml:"Id,omitempty"`
}

type FindRequest struct {
	web.PagedResultRequest
	Title string
}

type PagedArticleResponse struct {
	Articles   []Article
	TotalCount int64
}
