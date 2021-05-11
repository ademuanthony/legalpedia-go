package rulesofcourt

import (
	"io"
	"net/http"
	"strconv"
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

	str, err := m.server.Templates.ExecTemplateToString("rules/rules", struct {
		*web.CommonPageData
		Rules           []Rule
		Page            web.PageInfo
		States          []string
		State           string
		Sections        []string
		Section         string
		Title           string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Rules:          pagedRules.Rules,
		Page: web.PaginationResponseInfo(pagedRules.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"state": state, "section": section}),
		States:   states,
		State:    state,
		Sections: sections,
		Section:  section,
		Title:    "Rules of Court",
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

func (m *module) otherRulesPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pageReq := web.GetPanitionInfo(r)
	state := r.FormValue("state")
	section := r.FormValue("section")

	states, err := m.db.GetNames(r.Context(), "Other")
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

	pagedRules, err := m.db.Other(r.Context(), reqInput)
	if err != nil {
		log.Errorf("State execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch rules", web.ExpStatusError)
		return
	}

	str, err := m.server.Templates.ExecTemplateToString("rules/rules", struct {
		*web.CommonPageData
		Rules           []Rule
		Page            web.PageInfo
		States          []string
		State           string
		Sections        []string
		Section         string
		Title           string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Rules:          pagedRules.Rules,
		Page: web.PaginationResponseInfo(pagedRules.TotalCount, pageReq.Page,
			pageReq.PageSize, map[string]interface{}{"state": state, "section": section}),
		States:   states,
		State:    state,
		Sections: sections,
		Section:  section,
		Title:    "Other Rules of Court",
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

func (m *module) readRule(w http.ResponseWriter, r *http.Request) {
	idStr := web.GetIDParamCtx(r)
	id, _ := strconv.Atoi(idStr)
	rule, err := m.db.Detail(r.Context(), id)
	if err != nil {
		log.Errorf("Detail execute failure: %v", err)
		m.server.StatusPage(w, r, web.DefaultErrorCode, web.DefaultErrorMessage, "Unable to fetch rule", web.ExpStatusError)
		return
	}

	var parentHref string
	if strings.ToUpper(rule.Type) == "STATE" {
		parentHref = "/rules/state"
	} else {
		parentHref = "/rules/others"
	}

	str, err := m.server.Templates.ExecTemplateToString("rules/rule", struct {
		*web.CommonPageData
		Rule            Rule
		Page            web.PageInfo
		States          []string
		State           string
		Sections        []string
		Section         string
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: m.server.CommonData(r),
		Rule:           *rule,
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Rules Of Court",
				Href:      parentHref,
			},
			{
				HyperText: "Rule",
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
