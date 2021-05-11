package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/ademuanthony/legalpedia/article"
	"github.com/ademuanthony/legalpedia/postgres/models"
	"github.com/ademuanthony/legalpedia/web"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (m *PgDb) FetchArticles(ctx context.Context, req web.PagedResultRequest, term string) (*article.PagedArticleResponse, error) {
	var queries []qm.QueryMod
	if term != "" {
		statement := fmt.Sprintf("(%s LIKE '%%%s%%')",
			models.ArticleColumns.Title, strings.ToUpper(term),
		)
		queries = append(queries, qm.Where(statement))
	}

	totalCount, err := models.Articles(queries...).Count(ctx, m.db)
	if err != nil {
		return nil, err
	}

	queries = append(queries, qm.Limit(req.Limit), qm.Offset(req.Offset))

	records, err := models.Articles(
		queries...,
	).All(ctx, m.db)
	if err != nil {
		return nil, err
	}

	var articles = make([]article.Article, len(records))
	for i, art := range records {
		articles[i] = article.Article{
			ID:        art.ID.Int,
			UUID:      art.UUID,
			Title:     art.Title.String,
			Content:   art.Content.String,
			VersionNo: art.VersionNo.Int,
		}
	}

	return &article.PagedArticleResponse{
		Articles: articles, TotalCount: totalCount,
	}, nil
}

func (m *PgDb) FindByID(ctx context.Context, uuid string) (*article.Article, error) {
	art, err := models.FindArticle(ctx, m.db, uuid)
	if err != nil {
		return nil, err
	} // 08131944798

	return &article.Article{
		UUID:      art.UUID,
		ID:        art.ID.Int,
		Title:     art.Title.String,
		Content:   art.Content.String,
		VersionNo: art.ID.Int,
	}, nil
}
