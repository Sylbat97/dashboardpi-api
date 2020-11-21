package handler

import (
	"dashboard-pi/api/data"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
DiskHandler extends RootHandler
Return information related to Disk
*/
func DiskHandler(w http.ResponseWriter, r *http.Request) error {
	disk, err := data.GetDiskData()
	if err != nil {
		return NewHTTPError(err, 500, err.Error())
	}
	json.NewEncoder(w).Encode(disk)
	fmt.Println("Endpoint Hit: Disk")
	return nil
}
