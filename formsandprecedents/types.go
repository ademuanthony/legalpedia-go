package formsandprecedents

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

type module struct {
	db     store
	server *web.Server
}

type store interface {
	FormsPrecedenceCategories(ctx context.Context) ([]string, error)
	FindFormsPrecedences(ctx context.Context, req FindRequest) (*PagedResponse, error)
	FindFormsPrecedenceByID(ctx context.Context, id string) (*FormsPrecedence, error)
}

type FormsPrecedence struct {
	UUID      string `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	Category  string `boil:"category" json:"category,omitempty" toml:"category" yaml:"category,omitempty"`
	Title     string `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Content   string `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	VersionNo int    `boil:"version_no" json:"version_no,omitempty" toml:"version_no" yaml:"version_no,omitempty"`
	ID        int    `boil:"Id" json:"Id,omitempty" toml:"Id" yaml:"Id,omitempty"`
}

type FindRequest struct {
	web.PagedResultRequest
	Title    string
	Category string
}

type PagedResponse struct {
	Data       []FormsPrecedence
	TotalCount int64
}
