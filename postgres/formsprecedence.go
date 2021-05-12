package postgres

import (
	"context"
	"fmt"

	fp "github.com/ademuanthony/legalpedia/formsandprecedents"
	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pg *PgDb) FormsPrecedenceCategories(ctx context.Context) ([]string, error) {
	records, err := models.FormsPrecedences(
		qm.Select("DISTINCT category"),
	 	qm.OrderBy(models.FormsPrecedenceColumns.Category),
	 ).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var names = make([]string, len(records))
	for i, r := range records {
		names[i] = r.Category.String
	}
	return names, nil
}

func (pg *PgDb) FindFormsPrecedences(ctx context.Context, req fp.FindRequest) (*fp.PagedResponse, error) {
	var queries []qm.QueryMod
	if req.Category != "" {
		queries = append(queries, models.FormsPrecedenceWhere.Category.EQ(null.StringFrom(req.Category)))
	}

	if req.Title != "" {
		statement := fmt.Sprintf("(%s LIKE '%%%s%%')", models.FormsPrecedenceColumns.Title, req.Title)
		queries = append(queries, qm.Where(statement))
	}

	totalCount, err := models.FormsPrecedences(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Limit(req.Limit), qm.Offset(req.Offset))

	data, err := models.FormsPrecedences(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var result = make([]fp.FormsPrecedence, len(data))
	for i, d := range data {
		result[i] = fpFromModel(d)
	}

	return &fp.PagedResponse{Data: result, TotalCount: totalCount}, nil
}

func (pg *PgDb) FindFormsPrecedenceByID(ctx context.Context, id string) (*fp.FormsPrecedence, error) {
	data, err := models.FormsPrecedences(models.FormsPrecedenceWhere.UUID.EQ(id)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	result := fpFromModel(data)
	return &result, nil
}

func fpFromModel(d *models.FormsPrecedence) fp.FormsPrecedence {
	return fp.FormsPrecedence{
		ID:        d.ID.Int,
		UUID:      d.UUID,
		Title:     d.Title.String,
		Category:  d.Category.String,
		VersionNo: d.VersionNo.Int,
		Content:   d.Content.String,
	}
}
