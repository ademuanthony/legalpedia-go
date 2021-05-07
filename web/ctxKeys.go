package web

type contextKey int

const (
	CtxIDParam contextKey = iota
	CtxChartDataType
	CtxTimestamp
	CtxProposalRefID
	CtxProposalToken
	CtxAgendaId
)
