package postgres

import (
	"context"

	"github.com/ademuanthony/legalpedia/legalresource"
	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/ademuanthony/legalpedia/web"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (m *PgDb) FindLegalResources(ctx context.Context, req web.PagedResultRequest) (*legalresource.PagedResponse, error) {
	var queries []qm.QueryMod

	totalCount, err := models.Articles(queries...).Count(ctx, m.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Limit(req.Limit), qm.Offset(req.Offset))

	records, err := models.ForeignLegalResources(
		queries...,
	).All(ctx, m.db)
	if err != nil {
		return nil, err
	}

	var articles = make([]legalresource.LegalResource, len(records))
	for i, art := range records {
		articles[i] = legalresource.LegalResource{
			ID:          art.ID,
			Title:       art.Title.String,
			Description: art.Description.String,
			URL:         art.URL.String,
		}
	}

	return &legalresource.PagedResponse{
		Data: articles, TotalCount: totalCount,
	}, nil
}
