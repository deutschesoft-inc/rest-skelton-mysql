package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/config"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/sms"
	"net/http"
)

var (
	final           int
	warn            ent.ReturnPropMess
	res             []string
	RecID           int64
	filters, srcval string
	Res             int8
)

type (
	SmsController struct{}
)

func NewSmsController() *SmsController {
	return &SmsController{}
}

func (sc SmsController) RouterIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.SendSMSReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res = mod.GetNums(d.Usr_type)
	Res = mod.SMSMulti(res, d.Message)

	if Res == 1 {
		warn = cm.GetPropMess(1, 0, 0, "SMSler Başarıyla Gönderildi", nil)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadGateway)
		warn = cm.GetPropMess(1, 0, 0, "SMSler Gönderilirken Hata Oluştu", nil)
	}
	json.NewEncoder(w).Encode(warn)
}

func (sc SmsController) GetNums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.PhoneReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res = mod.GetNums(d.Usr_type)
	json.NewEncoder(w).Encode(res)
}

func (tc SmsController) SingleSend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var par ent.ReqSingleSMS
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(s, &par)
	if err != nil {
		panic(err)
	}

	Res = mod.SMSSingle(par.Num, par.Msg)

	if Res == 1 {
		warn = cm.GetPropMess(1, 0, 0, "SMS Başarıyla Gönderildi", nil)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadGateway)
		warn = cm.GetPropMess(1, 0, 0, "SMS Gönderilirken Hata Oluştu", nil)
	}

	json.NewEncoder(w).Encode(warn)
}

func (sc SmsController) GetList(w http.ResponseWriter, r *http.Request) {
	var paging ent.PaginateDataStruct
	var count, query string
	var args []interface{}
	var act string
	var result []ent.StudentListData
	result = []ent.StudentListData{}

	if r.Method == "POST" {
		r.ParseForm()
		start := r.FormValue("start")
		end := r.FormValue("length")
		draw := r.FormValue("draw")
		filters = r.FormValue("filters")

		if filters != "0" {
			act = "AND class_id=" + filters
		} else {
			act = "AND (class_id>0 OR class_id IS NULL)"
		}

		searchValue := r.FormValue("search[value]")
		if searchValue != "" {
			srcval = "AND student_num LIKE " + "%" + searchValue + "%"
		} else {
			srcval = ""
		}

		if draw == "1" {
			count = "SELECT count(id) as frequency FROM ben_auth WHERE usr_type=3 AND " + act + " " + srcval + ""
			err := config.DB.QueryRow(count).Scan(&final)
			if err != nil {
				fmt.Printf("QueryRow: %v\n", err)
			}
		}

		if searchValue != "" {
			p := "%" + searchValue + "%"
			args = []interface{}{p, start, end}
			query = "SELECT id,class_id,getClass(class_id),parent_id,getParent(parent_id),student_num,gender,usr_title,usr_name,phone,is_act FROM ben_auth WHERE usr_type=3 " + act + " AND student_num LIKE ? ORDER BY id DESC LIMIT ?,?"
			rows, err := config.DB.Query(query, args...)
			if err != nil && err.Error() != "no rows in result set" {
				panic(err.Error())
			}
			defer rows.Close()

			for rows.Next() {
				dat := ent.StudentListData{}
				err2 := rows.Scan(&dat.Id, &dat.Class_id, &dat.Class_nm, &dat.Parent_id, &dat.Parent_nm, &dat.Student_num, &dat.Gender, &dat.Usr_title, &dat.Usr_name, &dat.Phone, &dat.Is_act)
				if err2 != nil && err.Error() != "no rows in result set" {
					panic(err.Error())
				}
				result = append(result, dat)
			}
			final = len(result)
		} else {
			query = "SELECT id,class_id,getClass(class_id),parent_id,getParent(parent_id),student_num,gender,usr_title,usr_name,phone,is_act FROM ben_auth WHERE usr_type=3 " + act + " ORDER BY id DESC LIMIT ?,?"
			args = []interface{}{start, end}
			rows, err := config.DB.Query(query, args...)
			if err != nil && err.Error() != "no rows in result set" {
				panic(err.Error())
			}

			for rows.Next() {
				dat := ent.StudentListData{}
				err2 := rows.Scan(&dat.Id, &dat.Class_id, &dat.Class_nm, &dat.Parent_id, &dat.Parent_nm, &dat.Student_num, &dat.Gender, &dat.Usr_title, &dat.Usr_name, &dat.Phone, &dat.Is_act)
				if err2 != nil && err2.Error() != "no rows in result set" {
					panic(err2.Error())
				}
				result = append(result, dat)
			}
			defer rows.Close()
		}

		paging.DataList = result
		paging.Draw = draw
		paging.RecordsFiltered = final
		paging.RecordsFiltered = final
		e, err := json.Marshal(paging)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(e)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (sc SmsController) GetProcs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.StudentProcReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	data := mod.GetProcs(d)
	if data > 0 {
		warn = cm.GetPropMess(1, 0, 0, "İşlem başarılı", nil)
	} else {
		warn = cm.GetPropMess(2, 0, 0, "İşlem yapılırken hata oluştu !", nil)
	}
	json.NewEncoder(w).Encode(warn)
}
