package rulesofcourt

import (
	"io"
	"net/http"
	"strings"

	"github.com/ademuanthony/legalpedia/web"
)

var (
	sections = []string{
		"Orders", "Parts", "Schedules", "Forms", "Portrate Forms", "Civil Forms", "Appendix",
	}
)

func (m *module) stateRulesPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)
	state := r.FormValue("state")
	section := strings.Title(r.FormValue("section"))

	states, err := m.db.GetNames(r.Context(), "State")
	if err != nil {
		log.Errorf("GetNames execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch rules", web.ExpStatusError)
		return
	}

	if state == "" && len(states) > 0 {
		state = states[0]
	}

	if section == "" {
		section = sections[0]
	}
	
	reqInput := GetAllInput{
		PagedResultRequest: pageReq,
		State:              strings.ToUpper(state),
		Section:            strings.ToUpper(section),
	}

	pagedRules, err := m.db.State(r.Context(), reqInput)
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch rules", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("rules/state", struct {
		*web.CommonPageData
		Rules           []Rule
		Page            web.PageInfo
		States          []string
		State           string
		Sections        []string
		Section         string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Rules:          pagedRules.Rules,
		Page: web.PaginationResponseInfo(pagedRules.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"state": state, "section": section}),
		States:  states,
		State:   state,
		Sections: sections,
		Section: section,
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Rules Of Court",
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
