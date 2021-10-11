package data

import (
	"jembi/rpro-api/config"
)

type FlowsFlow struct {
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	FlowType     string `json:"flow_type"`
	BaseLanguage string `json:"base_language"`
}

func GetFlowsFlowList() ([]FlowsFlow, error) {

	flowsFlows := []FlowsFlow{}

	db := config.GetConnection(false)

	rows, err := db.Query("SELECT name, uuid, flow_type, base_language  FROM flows_flow")

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var ff FlowsFlow

		rows.Scan(&ff.Name, &ff.UUID, &ff.FlowType, &ff.BaseLanguage)

		flowsFlows = append(flowsFlows, ff)

	}

	defer db.Close()

	return flowsFlows, nil

}
