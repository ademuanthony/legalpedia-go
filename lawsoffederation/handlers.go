package lawsoffederation

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/ademuanthony/legalpedia/web"
)

func (m *module) indexPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)
	title := strings.ToUpper(r.FormValue("q"))
	year, _ := strconv.Atoi(r.FormValue("y"))

	reqInput := FindRequest{
		PagedResultRequest: pageReq,
		Title:              title,
		Year:               year,
	}

	result, err := m.db.GetAllLawsOfFederation(r.Context(), reqInput)
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch data", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("lawsoffederation/index", struct {
		*web.CommonPageData
		Items           []LawsOfFederation
		Page            web.PageInfo
		PageTitle       string
		SearchTerm      string
		Year            string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Items:          result.Data,
		Page: web.PaginationResponseInfo(result.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"q": title}),
		SearchTerm: title,
		Year:       r.FormValue("y"),
		PageTitle:  "Laws of Federation",
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Laws of Federation",
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

func (m *module) detailPage(w http.ResponseWriter, r *http.Request) {
}
