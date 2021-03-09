package teachers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/config"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/teachers"
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
	TeachersController struct{}
)

func NewTeachersController() *TeachersController {
	return &TeachersController{}
}

func (tc TeachersController) GetDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.OnlyIDReq
	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res = mod.GetDetail(d.ID)
	json.NewEncoder(w).Encode(res)
}

func (tc TeachersController) GetList(w http.ResponseWriter, r *http.Request) {
	var paging ent.PaginateDataStruct
	var count, query string
	var args []interface{}
	var act string
	var result []ent.TeacherListData
	result = []ent.TeacherListData{}

	if r.Method == "POST" {
		r.ParseForm()
		start := r.FormValue("start")
		end := r.FormValue("length")
		draw := r.FormValue("draw")
		filters = r.FormValue("filters")

		if filters != "0" {
			act = "branch_id=" + filters
		} else {
			act = "(branch_id>0 OR branch_id IS NULL)"
		}

		searchValue := r.FormValue("search[value]")
		if searchValue != "" {
			srcval = "AND usr_title LIKE " + "%" + searchValue + "%"
		} else {
			srcval = ""
		}

		if draw == "1" {
			count = "SELECT count(id) as frequency FROM ben_auth WHERE usr_type=2 AND " + act + " " + srcval + ""
			err := config.DB.QueryRow(count).Scan(&final)
			if err != nil {
				fmt.Printf("QueryRow: %v\n", err)
			}
		}

		if searchValue != "" {
			p := "%" + searchValue + "%"
			args = []interface{}{p, start, end}
			query = "SELECT id,getBranch(branch_id),usr_title,usr_code,usr_name,phone,is_act FROM ben_auth WHERE usr_type=2 AND usr_title LIKE ? ORDER BY usr_title ASC LIMIT ?,?"
			rows, err := config.DB.Query(query, args...)
			if err != nil && err.Error() != "no rows in result set" {
				panic(err.Error())
			}
			defer rows.Close()

			for rows.Next() {
				dat := ent.TeacherListData{}
				err2 := rows.Scan(&dat.Id, &dat.Branch_name, &dat.Usr_title, &dat.Usr_code, &dat.Usr_name, &dat.Phone, &dat.Is_act)
				if err2 != nil && err.Error() != "no rows in result set" {
					panic(err.Error())
				}
				result = append(result, dat)
			}
			final = len(result)
		} else {
			query = "SELECT id,getBranch(branch_id),usr_title,usr_code,usr_name,phone,is_act FROM ben_auth WHERE usr_type=2 AND " + act + " ORDER BY usr_title ASC LIMIT ?,?"
			args = []interface{}{start, end}
			rows, err := config.DB.Query(query, args...)
			if err != nil && err.Error() != "no rows in result set" {
				panic(err.Error())
			}

			for rows.Next() {
				dat := ent.TeacherListData{}
				err2 := rows.Scan(&dat.Id, &dat.Branch_name, &dat.Usr_title, &dat.Usr_code, &dat.Usr_name, &dat.Phone, &dat.Is_act)
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

func (tc TeachersController) GetProcs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.TeacherProcReq
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
