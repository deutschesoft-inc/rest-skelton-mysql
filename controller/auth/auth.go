package auth

import (
	"encoding/json"
	"io/ioutil"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/auth"
	"net/http"
)

type (
	AuthController struct{}
)

var (
	warn      ent.ReturnAuthMess
	RecID     int64
	status    int8
	LDEndUser ent.LoginDataEnduser
	Token     string
)

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (tc AuthController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.LoginReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	LDEndUser, status = mod.ChkLoginEndUser(d)
	RecID = LDEndUser.ID
	Token, err = cm.TokenEndUser(LDEndUser)

	if RecID > 0 && status != 0 {
		w.WriteHeader(http.StatusOK)
		warn = cm.GenAuthMess(1, RecID, LDEndUser.Usr_type, LDEndUser.Class_id, Token, nil)
	} else if RecID > 0 && status == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		warn = cm.GenAuthMess(2, 0, 0, nil, "Hesabınız bloke edilmiştir, Lütfen müşteri hizmetlerimize ulaşın", nil)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		warn = cm.GenAuthMess(2, 0, 0, nil, "Kullanıcı Adı veya Şifreniz yanlış", nil)
	}
	json.NewEncoder(w).Encode(warn)
}
