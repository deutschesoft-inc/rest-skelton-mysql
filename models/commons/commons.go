package commons

import (
	"database/sql"
	"fmt"
	cnf "main/config"
	ent "main/entities"
	"strings"
)

var (
	query, param string
)

func GetUserList(typer int) []ent.ComUserAutoList {
	if typer == 0 {
		param = " usr_type=1 OR usr_type=2 AND "
	} else {
		param = fmt.Sprintf(" usr_type=%d AND ", typer)
	}
	query := "SELECT id,getClass(class_id),usr_title FROM ben_auth WHERE " + param + " is_act=1 ORDER BY usr_title ASC"
	rows, err := cnf.DB.Query(query)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}
	defer rows.Close()
	geodata := []ent.ComUserAutoList{}
	for rows.Next() {
		dat := ent.ComUserAutoList{}
		err2 := rows.Scan(&dat.Id, &dat.Class_name, &dat.Usr_title)
		cnf.CheckErr(err2)
		geodata = append(geodata, dat)
	}
	return geodata
}

func GetLessonList() []ent.ComLessonAutoList {
	query := "SELECT id,l_code,l_name FROM lessons WHERE is_act=1 ORDER BY l_name ASC"
	rows, err := cnf.DB.Query(query)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}
	defer rows.Close()
	geodata := []ent.ComLessonAutoList{}
	for rows.Next() {
		dat := ent.ComLessonAutoList{}
		err2 := rows.Scan(&dat.Id, &dat.L_code, &dat.L_name)
		cnf.CheckErr(err2)
		geodata = append(geodata, dat)
	}
	return geodata
}

func GetClassList() []ent.ComClassAutoList {
	query := "SELECT id,CONCAT(c_num,'/',c_char) FROM classes WHERE is_act=1 ORDER BY id ASC"
	rows, err := cnf.DB.Query(query)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}
	defer rows.Close()
	geodata := []ent.ComClassAutoList{}
	for rows.Next() {
		dat := ent.ComClassAutoList{}
		err2 := rows.Scan(&dat.Id, &dat.Class_name)
		cnf.CheckErr(err2)
		geodata = append(geodata, dat)
	}
	return geodata
}

func GetStudentsList(class_id int32) []ent.ComRollCallList {
	query := "SELECT id,getClass(class_id),student_num,gender,usr_title FROM ben_auth WHERE class_id=? AND usr_type=3 AND is_act=1 ORDER BY usr_title ASC"
	rows, err := cnf.DB.Query(query, class_id)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}
	defer rows.Close()
	geodata := []ent.ComRollCallList{}
	for rows.Next() {
		dat := ent.ComRollCallList{}
		err2 := rows.Scan(&dat.Id, &dat.Class_nm, &dat.Student_num, &dat.Gender, &dat.Usr_title)
		cnf.CheckErr(err2)
		geodata = append(geodata, dat)
	}
	return geodata
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

func SyllAppoint(d ent.SyllAppointReq) (count int64) {
	Keys := []string{"student_id", "syllabus_id", "syllabus_lesson", "is_join"}
	Vals := []interface{}{d.Student_id, &d.Syllabus_id, &d.Syllabus_lesson, &d.Is_join}
	Ques := []string{"?", "?", "?", "?"}

	insert := fmt.Sprintf("INSERT INTO ben_auth(%s) VALUES(%s)", strings.Join(Keys, ","), strings.Join(Ques, ","))
	res, err := cnf.DB.Exec(insert, Vals...)
	count, _ = res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return
}
