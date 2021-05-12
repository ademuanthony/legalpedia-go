package formsandprecedents

import (
	"io"
	"net/http"
	"strings"

	"github.com/ademuanthony/legalpedia/web"
)

func (m *module) listPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)
	category := r.FormValue("category")
	if strings.ToLower(category) == "all" {
		category = ""
	}
	title := r.FormValue("q")

	reqInput := FindRequest{
		PagedResultRequest: pageReq,
		Title:              title,
		Category:           category,
	}

	categories, err := m.db.FormsPrecedenceCategories(r.Context())
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch data", web.ExpStatusError)
		return
	}

	pagedRules, err := m.db.FindFormsPrecedences(r.Context(), reqInput)
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch data", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("formspresedents/index", struct {
		*web.CommonPageData
		List            []FormsPrecedence
		Page            web.PageInfo
		PageTitle       string
		Categories      []string
		Category        string
		SearchTerm      string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		List:           pagedRules.Data,
		Page: web.PaginationResponseInfo(pagedRules.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"category": category, "q": title}),
		Categories: categories,
		Category:   category,
		SearchTerm: title,
		PageTitle:  "Forms and Precedents",
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Forms and Precedents",
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

func (m *module) viewPage(w http.ResponseWriter, r *http.Request) {
	id := web.GetIDParamCtx(r)
	record, err := m.db.FindFormsPrecedenceByID(r.Context(), id)
	if err != nil {
		log.Errorf("Detail execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch data", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("formspresedents/view", struct {
		*web.CommonPageData
		Item            FormsPrecedence
		Page            web.PageInfo
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Item:           *record,
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Forms and Precedents",
				Href:      "/formspresedents",
			},
			{
				HyperText: record.Title,
				Href:      "!#",
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
