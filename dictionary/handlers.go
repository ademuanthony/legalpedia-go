package dictionary

import (
	"io"
	"net/http"

	"github.com/ademuanthony/legalpedia/web"
)

func (m *module) indexPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)
	title := r.FormValue("q")

	reqInput := FindRequest{
		PagedResultRequest: pageReq,
		Title:              title,
	}

	result, err := m.db.FindDictionaries(r.Context(), reqInput)
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch data", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("dictionary_index", struct {
		*web.CommonPageData
		Items           []Dictionay
		Page            web.PageInfo
		PageTitle       string
		SearchTerm      string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Items:          result.Data,
		Page: web.PaginationResponseInfo(result.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"q": title}),
		SearchTerm: title,
		PageTitle:  "Law Dictionary",
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Law Dictionary",
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
