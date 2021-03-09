package profile

import (
	"database/sql"
	"fmt"
	cnf "main/config"
	ent "main/entities"
	"strings"
)

var (
	query string
	Vals  []interface{}
	Keys  []string
	Ques  []string
	resp  int8 = 1
)

func GetProfileDet(id int32, user_type int32) (d ent.ProfileDetRes) {
	if user_type == 1 || user_type == 2 {
		Vals = []interface{}{&d.ID, &d.Hes_code, &d.Branch_nm, &d.Gender, &d.Usr_title, &d.Usr_code, &d.Identity, &d.Email, &d.Phone, &d.Lnk_zoom, &d.Lnk_avatar, &d.Lastlogin}
		query = "SELECT id,hes_code,getBranch(branch_id),gender,usr_title,usr_code,usr_name,email,phone,lnk_zoom,lnk_avatar,lastlogin FROM ben_auth WHERE id=? AND (usr_type=2 OR usr_type=1)"
	} else if user_type == 3 {
		Vals = []interface{}{&d.ID, &d.Hes_code, &d.Class_nm, &d.Parent_nm, &d.Student_num, &d.Gender, &d.Usr_title, &d.Usr_code, &d.Identity, &d.Email, &d.Phone, &d.Lnk_avatar, &d.Discontinuity, &d.Lastlogin}
		query = "SELECT id,hes_code,getClass(class_id),getParent(parent_id),student_num,gender,usr_title,usr_code,usr_name,email,phone,lnk_avatar,discontinuity,lastlogin FROM ben_auth WHERE id=? AND usr_type=3"
	} else {
		Vals = []interface{}{&d.ID, &d.Hes_code, &d.Class_nm, &d.Parent_nm, &d.Student_num, &d.Gender, &d.Usr_title, &d.Usr_code, &d.Identity, &d.Email, &d.Phone, &d.Lnk_avatar, &d.Lastlogin}
		query = "SELECT id,hes_code,getClass(class_id),getParent(parent_id),student_num,gender,usr_title,usr_code,usr_name,email,phone,lnk_avatar,lastlogin FROM ben_auth WHERE id=? AND usr_type=4"
	}
	cnf.DB.QueryRow(query, id).Scan(Vals...)
	return
}

func GetProcs(user_type int32, id int32, d ent.ProfileProcReq) (count int8) {
	Keys = []string{}
	Vals = []interface{}{}
	Ques = []string{}

	if d.Hes_code != "" {
		Keys = append(Keys, "hes_code=?")
		Vals = append(Vals, d.Hes_code)
	}
	if d.Email != "" {
		Keys = append(Keys, "email=?")
		Vals = append(Vals, d.Email)
	}
	if d.Phone != "" {
		Keys = append(Keys, "phone=?")
		Vals = append(Vals, d.Phone)
	}
	if d.Lnk_zoom != "" {
		Keys = append(Keys, "lnk_zoom=?")
		Vals = append(Vals, d.Lnk_zoom)
	}

	query := fmt.Sprintf("UPDATE ben_auth SET %s WHERE id=%d AND usr_type=%d", strings.Join(Keys, ","), id, user_type)
	qUp, err := cnf.DB.Prepare(query)
	_, err = qUp.Exec(Vals...)
	if err != nil && err != sql.ErrNoRows {
		count = 2
		//panic(err.Error())
	}
	defer qUp.Close()
	if err != nil && err != sql.ErrNoRows {
		//panic(err.Error())
		count = 2
	}
	count = 1
	return
}

func GetDashboard(d ent.DashReq) (dat ent.DashRes) {
	if d.Typer == 1 {
		query = "SELECT (SELECT COUNT(id) FROM ben_auth WHERE usr_type=2), (SELECT COUNT(id) FROM ben_auth WHERE usr_type=3), (SELECT COUNT(id) FROM ben_auth WHERE usr_type=4)"
	} else if d.Typer == 2 {
		query = "SELECT (SELECT COUNT(id) FROM ben_auth WHERE usr_type=2), (SELECT COUNT(id) FROM ben_auth WHERE usr_type=3), (SELECT COUNT(id) FROM ben_auth WHERE usr_type=4)"
	} else if d.Typer == 3 {
		query = "SELECT (SELECT COUNT(id) FROM ben_auth WHERE usr_type=2), (SELECT COUNT(id) FROM ben_auth WHERE usr_type=3), (SELECT COUNT(id) FROM ben_auth WHERE usr_type=4)"
	}
	cnf.DB.QueryRow(query).Scan(&dat.TeacherCnt, &dat.StudentCnt, &dat.ParentCnt)
	return
}

func ChangePass(íd int32, pass string) (count int64) {
	update := `UPDATE ben_auth SET usr_pass=? WHERE id = ?;`
	_, err := cnf.DB.Exec(update, pass, íd)
	if err != nil {
		panic(err)
	} else {
		count = 1
	}
	return
}
