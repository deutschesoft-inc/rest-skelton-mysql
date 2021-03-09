package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	cnf "main/config"
	ent "main/entities"
	cm "main/libs/shared"
	mod "main/models/messages"
	"net/http"
)

var (
	final                          int
	warn                           ent.ReturnPropMess
	res                            interface{}
	RecID                          int64
	filters, srcval, fltr, usrtype string
)

type (
	MessagesController struct{}
)

func NewMessagesController() *MessagesController {
	return &MessagesController{}
}

func (mc MessagesController) GetDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.OnlyIDReq
	t := cm.ParseToken(r)
	if t == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "authentication error"})
		return
	}
	UserID := int32(t["id"].(float64))

	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	res = mod.GetDetail(UserID, d.ID)
	json.NewEncoder(w).Encode(res)
}

func (mc MessagesController) GetList(w http.ResponseWriter, r *http.Request) {
	var paging ent.PaginateDataStruct
	var count, query string
	var args []interface{}
	var result []ent.MessagesListData
	result = []ent.MessagesListData{}

	t := cm.ParseToken(r)
	if t == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "authentication error"})
		return
	}
	UserID := int32(t["id"].(float64))
	UserType := int32(t["usr_type"].(float64))

	if r.Method == "POST" {
		r.ParseForm()
		start := r.FormValue("start")
		end := r.FormValue("length")
		draw := r.FormValue("draw")
		filters = r.FormValue("filters")
		searchValue := r.FormValue("search[value]")

		if UserType == 3 || UserType == 4 {
			ClassID := int32(t["class_id"].(float64))
			usrtype = fmt.Sprintf("OR (m_type=1 AND class_id=%d AND m_to=0) AND", ClassID)
		} else {
			usrtype = "AND"
		}

		if filters != "0" {
			fltr = "AND class_id=" + filters + " " + usrtype
		} else {
			fltr = usrtype
		}

		if searchValue != "" {
			srcval = fmt.Sprintf("WHERE m_to=%d AND is_deleted=0 %s getUserTitle(m_from) LIKE '%%%s%%'", UserID, fltr, searchValue)
		} else {
			srcval = fmt.Sprintf("WHERE m_to=%d %s is_deleted=0", UserID, fltr)
		}

		if draw == "1" {
			count = "SELECT count(id) as frequency FROM messages " + srcval + ""
			err := cnf.DB.QueryRow(count).Scan(&final)
			if err != nil {
				fmt.Printf("QueryRow: %v\n", err)
			}
		}

		if searchValue != "" {
			p := "%" + searchValue + "%"
			args = []interface{}{UserID, p, start, end}
			query = "SELECT id,class_id,getClass(class_id),m_from,getUserTitle(m_from),message,on_date,is_readed,is_answered FROM messages WHERE m_to=? " + fltr + " is_deleted=0 AND getUserTitle(m_from) LIKE ? ORDER BY on_date DESC LIMIT ?,?"
			rows, err := cnf.DB.Query(query, args...)
			if err != nil {
				panic(err.Error())
			}
			defer rows.Close()

			for rows.Next() {
				dat := ent.MessagesListData{}
				err2 := rows.Scan(&dat.Id, &dat.Class_id, &dat.Class_name, &dat.M_from_id, &dat.M_from_name, &dat.Message, &dat.On_date, &dat.Is_readed, &dat.Is_answered)
				if err2 != nil && err.Error() != "no rows in result set" {
					panic(err.Error())
				}
				result = append(result, dat)
			}
			final = len(result)
		} else {
			query = "SELECT id,class_id,getClass(class_id),m_from,getUserTitle(m_from),message,on_date,is_readed,is_answered FROM messages WHERE m_to=? " + fltr + " is_deleted=0 ORDER BY on_date DESC LIMIT ?,?"
			println(query)
			args = []interface{}{UserID, start, end}
			rows, err := cnf.DB.Query(query, args...)
			if err != nil && err.Error() != "no rows in result set" {
				panic(err.Error())
			}

			for rows.Next() {
				dat := ent.MessagesListData{}
				err2 := rows.Scan(&dat.Id, &dat.Class_id, &dat.Class_name, &dat.M_from_id, &dat.M_from_name, &dat.Message, &dat.On_date, &dat.Is_readed, &dat.Is_answered)
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

func (mc MessagesController) GetProcs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d ent.MessagesProcReq
	t := cm.ParseToken(r)
	if t == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "authentication error"})
		return
	}
	UserID := int32(t["id"].(float64))
	UserType := int32(t["usr_type"].(float64))

	req, err := ioutil.ReadAll(r.Body)
	cm.CheckErr(err)
	err = json.Unmarshal(req, &d)
	cm.CheckErr(err)

	data := mod.GetProcs(UserType, UserID, d)
	if data > 0 {
		warn = cm.GetPropMess(1, 0, 0, "İşlem başarılı", nil)
	} else {
		warn = cm.GetPropMess(2, 0, 0, "İşlem yapılırken hata oluştu !", nil)
	}
	json.NewEncoder(w).Encode(warn)
}

func (mc MessagesController) GetBoxBlink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := cm.ParseToken(r)
	if t == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "authentication error"})
		return
	}
	UserID := int32(t["id"].(float64))

	resp := mod.GetBoxBlink(UserID)
	if resp > 0 {
		warn = cm.GetPropMess(1, 0, int(UserID), "Okunmamış Mesaj Var", nil)
	} else {
		warn = cm.GetPropMess(2, 0, -1, "Yeni Mesaj Yok", nil)
	}

	json.NewEncoder(w).Encode(warn)
}
