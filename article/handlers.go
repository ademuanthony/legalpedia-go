package article

import (
	"io"
	"net/http"

	"github.com/ademuanthony/legalpedia/web"
)

// PAGE HANDLERS

func (m *module) articlesPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)

	pagedArticles, err := m.db.FetchArticles(r.Context(), pageReq, r.FormValue("q"))
	if err != nil {
		log.Errorf("fetchArticles execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch articles", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("articles", struct {
		*web.CommonPageData
		Articles        []Article
		Page            web.PageInfo
		Term            string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Articles:       pagedArticles.Articles,
		Page: web.PaginationResponseInfo(pagedArticles.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"q": r.FormValue("q")}),
		Term: r.FormValue("q"),
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Articles",
				Active:    true,
			},
		},
	})

	if err != nil {
		log.Errorf("Template execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "", web.ExpStatusError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	if _, err = io.WriteString(w, str); err != nil {
		log.Error(err)
	}
}

func (m *module) articlePage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	article, err := m.db.FindByID(r.Context(), web.GetIDParamCtx(r))
	if err != nil {
		log.Errorf("findByID execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch articles", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("article", struct {
		*web.CommonPageData
		Article         *Article
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Article:        article,
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Articles",
				Href:      "/articles",
			},
			{
				HyperText: article.Title,
				Active:    true,
			},
		},
	})

	if err != nil {
		log.Errorf("Template execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "", web.ExpStatusError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	if _, err = io.WriteString(w, str); err != nil {
		log.Error(err)
	}
}

// END PAGED HANDLES

// API ENDPOINTS
func (m *module) articlesEndpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)

	pagedArticles, err := m.db.FetchArticles(r.Context(), pageReq, r.FormValue("term"))
	if err != nil {
		log.Errorf("fetchArticles execute failure: %v", err)
		web.RenderErrorfJSON(w, "Error in fetching articles")
		return
	}

	web.RenderJSON(w, pagedArticles)
}

// END API ENDPOINT
