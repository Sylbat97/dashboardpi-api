package handler

import (
	"dashboard-pi/api/data"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
CPUHandler extends RootHandler
Return information related to CPU
Ameliorations : Monitor CPU in background and cache data to reduce response time
*/
func CPUHandler(w http.ResponseWriter, r *http.Request) error {
	cpu, err := data.GetCPUData()
	if err != nil {
		return NewHTTPError(err, 500, err.Error())
	}
	json.NewEncoder(w).Encode(cpu)
	fmt.Println("Endpoint Hit: CPU")
	return nil
}
