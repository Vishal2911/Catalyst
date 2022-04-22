package store

import (
	"magantas/catalyst/model"
	"time"
)

func (store Store) DashBoard() (model.DashBoard) {
	dashboard := model.DashBoard{}
	
	for _ , value := range store.Token {
		currentday := time.Now()
		diff := currentday.Sub(value.CreatedAT)
		if int(diff.Hours()/24) > 7 {
			dashboard.ExpiredToken = append(dashboard.ExpiredToken, value)
		}else if value.Active == false {
			dashboard.RevokedToken = append(dashboard.RevokedToken, value)
		}else if value.Used == true {
			dashboard.UsedToken = append(dashboard.UsedToken, value)
		}else {
			dashboard.PendingToken = append(dashboard.PendingToken, value)
		}
	}
	return dashboard

}
