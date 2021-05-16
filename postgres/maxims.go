package postgres

import (
	"context"
	"fmt"

	"github.com/ademuanthony/legalpedia/maxim"
	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pg *PgDb) FindMaximes(ctx context.Context, req maxim.FindRequest) (*maxim.PagedResponse, error) {
	var queries []qm.QueryMod
	if req.Title != "" {
		statement := fmt.Sprintf("(%s LIKE '%%%s%%')",
			models.MaximColumns.Title, req.Title,
		)
		queries = append(queries, qm.Where(statement))
	}

	totalCount, err := models.Maxims(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Limit(req.Limit), qm.Offset(req.Offset))

	data, err := models.Maxims(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var result = make([]maxim.Maxim, len(data))
	for i, d := range data {
		result[i] = maxim.Maxim{
			ID:        d.ID,
			UUID:      d.UUID.String,
			Title:     d.Title.String,
			Content:   d.Content.String,
			VersionNo: d.VersionNo.Int,
		}
	}

	return &maxim.PagedResponse{
		Data:       result,
		TotalCount: totalCount,
	}, nil
}

func (pg *PgDb) FindMaximByID(ctx context.Context, id string) (*maxim.Maxim, error) {
	d, err := models.Maxims(models.MaximWhere.UUID.EQ(null.StringFrom(id))).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}
	return &maxim.Maxim{
		ID:        d.ID,
		UUID:      d.UUID.String,
		Title:     d.Title.String,
		Content:   d.Content.String,
		VersionNo: d.VersionNo.Int,
	}, nil
}
