package students

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

func GetDetail(id int64) (d ent.StudentRes) {
	query1 := "SELECT id,usr_type,hes_code,class_id,parent_id,student_num,gender,usr_title,usr_code,usr_name,usr_pass,email,phone,is_act FROM ben_auth WHERE usr_type=3 AND id=?"
	err1 := cnf.DB.QueryRow(query1, id).Scan(&d.Id, &d.Usr_type, &d.Hes_code, &d.Class_id, &d.Parent_id, &d.Student_num, &d.Gender, &d.Usr_title, &d.Usr_code, &d.Usr_name, &d.Usr_pass, &d.Email, &d.Phone, &d.Is_act)
	if err1 != nil && err1 != sql.ErrNoRows {
		panic(err1.Error())
	}
	return
}

func GetProcs(d ent.StudentProcReq) (count int64) {
	count = 0
	if d.Pid == 1 {
		Keys := []string{"usr_type", "hes_code", "class_id", "parent_id", "student_num", "gender", "usr_title", "usr_name", "usr_pass", "email", "phone"}
		Vals := []interface{}{3, d.Hes_code, d.Class_id, d.Parent_id, d.Student_num, d.Gender, d.Usr_title, d.Usr_name, d.Usr_pass, d.Email, d.Phone}
		Ques := []string{"?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?"}

		insert := fmt.Sprintf("INSERT INTO ben_auth(%s) VALUES(%s)", strings.Join(Keys, ","), strings.Join(Ques, ","))
		res, err := cnf.DB.Exec(insert, Vals...)
		count, _ = res.LastInsertId()
		if err != nil {
			panic(err)
		}
	}
	if d.Pid == 2 {
		//update := fmt.Sprintf("UPDATE web.w_blogs SET (%s) = (%s) WHERE id=%d", strings.Join(Keys, ","), Vals, d.Id)
		update := `UPDATE ben_auth SET hes_code=?,class_id=?,parent_id=?,student_num=?,gender=?,usr_title=?,usr_name=?,usr_pass=?,email=?,phone=? WHERE usr_type=3 AND id = ?`
		_, err := cnf.DB.Exec(update, d.Hes_code, d.Class_id, d.Parent_id, d.Student_num, d.Gender, d.Usr_title, d.Usr_name, d.Usr_pass, d.Email, d.Phone, d.Id)
		if err != nil {
			panic(err)
		} else {
			count = 1
		}
	}
	if d.Pid == 3 {
		del := `DELETE FROM ben_auth WHERE usr_type=3 AND id=?`
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
		actup := `UPDATE ben_auth SET is_act=? WHERE usr_type=3 AND id = ?`
		res, err := cnf.DB.Exec(actup, d.Is_act, d.Id)
		if err != nil {
			panic(err)
		}
		count, err = res.RowsAffected()
		if err != nil {
			panic(err)
		}
	}
	if d.Pid == 5 {
		avatarup := `UPDATE ben_auth SET lnk_avatar=? WHERE usr_type=3 AND usr_name = ?`
		_, err := cnf.DB.Exec(avatarup, d.Lnk_avatar, d.Usr_name)
		if err != nil {
			panic(err)
		}
		count = 1
		if err != nil {
			panic(err)
		}
	}
	return
}
