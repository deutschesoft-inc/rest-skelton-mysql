package sms

import (
	"fmt"
	"io/ioutil"
	"log"
	cnf "main/config"
	ent "main/entities"
	"net/http"
	"os"
	"strings"
)

var (
	keycnt, param int
	count         int64
	query         string
)

func SMSSingle(Num string, Msg string) (res int8) {
	auth := cnf.Auth
	url := "http://api.netgsm.com.tr/bulkhttppost.asp"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("usercode", auth.User)
	q.Add("password", auth.Pass)
	q.Add("gsmno", Num)
	q.Add("message", Msg)
	q.Add("msgheader", auth.Header)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	check := strings.SplitAfter(string(body), " ")
	seos := strings.TrimSpace(check[0])
	if seos == "00" {
		res = 1
	} else {
		res = 2
	}
	return
}

func SMSMulti(Nums []string, Msg string) (res int8) {
	auth := cnf.Auth
	url := "http://api.netgsm.com.tr/bulkhttppost.asp"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("usercode", auth.User)
	q.Add("password", auth.Pass)
	q.Add("gsmno", strings.Join(Nums, ","))
	q.Add("message", Msg)
	q.Add("msgheader", auth.Header)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	check := strings.SplitAfter(string(body), " ")
	seos := strings.TrimSpace(check[0])
	if seos == "00" {
		res = 1
	} else {
		res = 2
	}
	return
}

func GetNums(usr_type int) (result []string) {
	var number string
	if usr_type > 10 {
		query = "SELECT REPLACE(phone, ' ', '') FROM `ben_auth` WHERE phone IS NOT NULL AND class_id=SUBSTR(?,2)"
		param = usr_type
	} else {
		query = "SELECT REPLACE(phone, ' ', '') FROM `ben_auth` WHERE phone IS NOT NULL AND usr_type=?"
		param = usr_type
	}

	rows, err := cnf.DB.Query(query, usr_type)
	if err != nil && err.Error() != "no rows in result set" {
		panic(err.Error())
	}

	for rows.Next() {
		err2 := rows.Scan(&number)
		if err2 != nil && err2.Error() != "no rows in result set" {
			panic(err2.Error())
		}
		result = append(result, number)
	}
	defer rows.Close()

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
	return
}
