package students

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/config"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/students"
	"net/http"
)

var (
	final           int
	warn            ent.ReturnPropMess
	res             interface{}
	RecID           int64
	filters, srcval string
)

type (
	StudentsController struct{}
)

func NewStudentsController() *StudentsController {
	return &StudentsController{}
}

func (sc StudentsController) GetDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.OnlyIDReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res = mod.GetDetail(d.ID)
	json.NewEncoder(w).Encode(res)
}

func (sc StudentsController) GetList(w http.ResponseWriter, r *http.Request) {
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
			count = "SELECT count(id) as frequency FROM ben_auth WHERE usr_type=3 " + act + " " + srcval + ""
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

func (sc StudentsController) GetProcs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.StudentProcReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	data := mod.GetProcs(d)
	if data > 0 {
		warn = cm.GetPropMess(1, 0, 0, "????lem ba??ar??l??", nil)
	} else {
		warn = cm.GetPropMess(2, 0, 0, "????lem yap??l??rken hata olu??tu !", nil)
	}
	json.NewEncoder(w).Encode(warn)
}
