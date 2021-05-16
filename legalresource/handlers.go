package legalresource

import (
	"io"
	"net/http"
	"strings"

	"github.com/ademuanthony/legalpedia/web"
)

func (m *module) indexPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	reqInput := web.GetPanitionInfo(r)
	title := strings.ToUpper(r.FormValue("q"))

	result, err := m.db.FindLegalResources(r.Context(), reqInput)
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch data", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("legalresource/index", struct {
		*web.CommonPageData
		Items           []LegalResource
		Page            web.PageInfo
		PageTitle       string
		SearchTerm      string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Items:          result.Data,
		Page: web.PaginationResponseInfo(result.TotalCount, reqInput.Page,
			reqInput.PageSize, map[string]interface{}{"q": title}),
		SearchTerm: title,
		PageTitle:  "Foreign Legal Resources",
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Foreign Legal Resources",
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
