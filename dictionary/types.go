package dictionary

import (
	"context"

	"github.com/ademuanthony/legalpedia/web"
)

type store interface {
	FindDictionaries(ctx context.Context, req FindRequest) (*PagedResponse, error)
	FindDictionaryByID(ctx context.Context, id string) (*Dictionay, error)
}

type module struct {
	db     store
	server *web.Server
}

type Dictionay struct {
	UUID      string `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
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
	Data       []Dictionay
	TotalCount int64
}
