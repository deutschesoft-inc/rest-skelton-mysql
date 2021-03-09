package commons

import (
	"encoding/json"
	"io/ioutil"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/commons"
	"net/http"
)

type (
	CommonsController struct{}
)

var (
	warn     ent.ReturnPropMess
	ChkPhone bool
	Identity string
	RecID    int64
	res      interface{}
)

func NewCommonsController() *CommonsController {
	return &CommonsController{}
}

func (tc CommonsController) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.OnlyTyperReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res := mod.GetUserList(d.Typer)
	json.NewEncoder(w).Encode(res)
}

func (tc CommonsController) GetLessonList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := mod.GetLessonList()
	json.NewEncoder(w).Encode(res)
}

func (tc CommonsController) GetClassList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := mod.GetClassList()
	json.NewEncoder(w).Encode(res)
}
func (tc CommonsController) GetStudentsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.OnlyClassIDReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res := mod.GetStudentsList(d.Class_id)
	json.NewEncoder(w).Encode(res)
}

func (tc CommonsController) GetDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.DashReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res := mod.GetDashboard(d)
	json.NewEncoder(w).Encode(res)
}

func (tc CommonsController) ChangePass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.ChangePass
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	RecID = mod.ChangePass(d.Id, d.Pass)
	if RecID > 0 {
		res = cm.GetPropMess(1, 0, 0, "Şifre Başarıyla Güncellenmiştir", nil)
	} else {
		res = cm.GetPropMess(2, 0, 0, "Şifre Güncellenirken Hata Oluştu", nil)
	}

	json.NewEncoder(w).Encode(res)
}

func (tc CommonsController) SyllAppoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.SyllAppointReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	RecID = mod.SyllAppoint(d)
	if RecID > 0 {
		res = cm.GetPropMess(1, 0, 0, "Şifre Başarıyla Güncellenmiştir", nil)
	} else {
		res = cm.GetPropMess(2, 0, 0, "Şifre Güncellenirken Hata Oluştu", nil)
	}

	json.NewEncoder(w).Encode(res)
}
