package services

import "gokanban/internal/flow"

func GetFlowBySlug(slug string) flow.FlowInterface {
	switch slug {
	case "standard":
		return flow.NewStandardFlow()
	case "demo":
		return flow.NewDemoFlow()
	case "simple":
		return flow.NewSimpleFlow()
	case "employment_search":
		return flow.NewEmploymentSearch()
	default:
		return flow.NewStandardFlow()
	}
}
