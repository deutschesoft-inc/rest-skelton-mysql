package profile

import (
	"encoding/json"
	"io/ioutil"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/profile"
	"net/http"
)

type (
	ProfileController struct{}
)

var (
	warn     ent.ReturnPropMess
	ChkPhone bool
	Identity string
	RecID    int64
	res      interface{}
)

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

func (pc ProfileController) GetProfileDet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := cm.ParseToken(r)
	if t == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "authentication error"})
		return
	}
	UserID := int32(t["id"].(float64))
	UsrType := int32(t["usr_type"].(float64))

	res := mod.GetProfileDet(UserID, UsrType)
	json.NewEncoder(w).Encode(res)
}

func (pc ProfileController) GetProcs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := cm.ParseToken(r)
	if t == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "authentication error"})
		return
	}
	UserID := int32(t["id"].(float64))
	UsrType := int32(t["usr_type"].(float64))

	var d ent.ProfileProcReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	data := mod.GetProcs(UsrType, UserID, d)
	if data > 0 {
		warn = cm.GetPropMess(1, 0, 0, "İşlem başarılı", nil)
	} else {
		warn = cm.GetPropMess(2, 0, 0, "İşlem yapılırken hata oluştu !", nil)
	}
	json.NewEncoder(w).Encode(warn)
}

func (pc ProfileController) GetDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.DashReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res := mod.GetDashboard(d)
	json.NewEncoder(w).Encode(res)
}

func (pc ProfileController) ChangePass(w http.ResponseWriter, r *http.Request) {
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
