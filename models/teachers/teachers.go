package teachers

import (
	"database/sql"
	"fmt"
	cnf "main/config"
	ent "main/entities"
	"strings"
)

var (
	keycnt int
	count  int64
)

func GetDetail(id int64) (d ent.TeachersRes) {
	query1 := "SELECT id,usr_type,hes_code,branch_id,gender,usr_title,usr_code,usr_name,usr_pass,email,phone,lnk_zoom,is_act FROM ben_auth WHERE usr_type=2 AND id=?"
	err1 := cnf.DB.QueryRow(query1, id).Scan(&d.Id, &d.Usr_type, &d.Hes_code, &d.Branch_id, &d.Gender, &d.Usr_title, &d.Usr_code, &d.Usr_name, &d.Usr_pass, &d.Email, &d.Phone, &d.Lnk_zoom, &d.Is_act)
	if err1 != nil && err1 != sql.ErrNoRows {
		panic(err1.Error())
	}
	return
}

func GetProcs(d ent.TeacherProcReq) (count int64) {
	count = 0
	if d.Pid == 1 {
		Keys := []string{"usr_type", "hes_code", "branch_id", "gender", "usr_title", "usr_code", "usr_name", "usr_pass", "email", "phone", "lnk_zoom"}
		Vals := []interface{}{2, d.Hes_code, d.Branch_id, d.Gender, &d.Usr_title, &d.Usr_code, &d.Usr_name, &d.Usr_pass, &d.Email, &d.Phone, &d.Lnk_zoom}
		Ques := []string{"?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?"}

		insert := fmt.Sprintf("INSERT INTO ben_auth(%s) VALUES(%s)", strings.Join(Keys, ","), strings.Join(Ques, ","))
		res, err := cnf.DB.Exec(insert, Vals...)
		count, _ = res.LastInsertId()
		if err != nil {
			panic(err)
		}
	}
	if d.Pid == 2 {
		update := `UPDATE ben_auth SET usr_type=?,hes_code=?,branch_id=?,gender=?,usr_title=?,usr_code=?,usr_name=?,usr_pass=?,email=?,phone=?,lnk_zoom=? WHERE usr_type=2 AND id = ?;`
		_, err := cnf.DB.Exec(update, d.Usr_type, d.Hes_code, d.Branch_id, d.Gender, d.Usr_title, d.Usr_code, d.Usr_name, d.Usr_pass, d.Email, d.Phone, d.Lnk_zoom, d.Id)
		if err != nil {
			panic(err)
		} else {
			count = 1
		}
	}
	if d.Pid == 3 {
		del := `DELETE FROM ben_auth WHERE usr_type=2 AND id=?;`
		res, err := cnf.DB.Exec(del, d.Id)
		if err != nil {
			panic(err)
		}
		count, err = res.RowsAffected()
		if err != nil {
			panic(err)
		}
	}
	if d.Pid == 4 {
		actup := `UPDATE ben_auth SET is_act=? WHERE usr_type=2 AND id = ?;`
		res, err := cnf.DB.Exec(actup, d.Is_act, d.Id)
		if err != nil {
			panic(err)
		}
		count, err = res.RowsAffected()
		if err != nil {
			panic(err)
		}
	}
	return
}
