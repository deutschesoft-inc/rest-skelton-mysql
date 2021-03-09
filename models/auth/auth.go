package auth

import (
	"database/sql"
	"main/config"
	ent "main/entities"
)

var (
	Custype int8
)

func ChkLoginEndUser(dat ent.LoginReq) (res ent.LoginDataEnduser, status int8) {
	query := "SELECT id,usr_type,class_id,usr_title,usr_code, usr_name, lastlogin, is_act FROM ben_auth WHERE usr_name=? AND usr_pass=?"
	err := config.DB.QueryRow(query, dat.Username, dat.Password).Scan(&res.ID, &res.Usr_type, &res.Class_id, &res.Title, &res.Usr_code, &res.Identity, &res.Lastlogin, &status)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}
	return
}
