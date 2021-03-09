package messages

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

func GetDetail(user_id int32, id int64) (d ent.MessagesRes) {
	query1 := "SELECT id,message FROM messages WHERE id=? AND m_to=?"
	err1 := cnf.DB.QueryRow(query1, id, user_id).Scan(&d.Id, &d.Message)
	if err1 != nil && err1 != sql.ErrNoRows {
		panic(err1.Error())
	}

	go func() {
		query2 := "SELECT upReaded(?)"
		cnf.DB.QueryRow(query2, id)
	}()
	return
}

func GetProcs(user_type int32, user_id int32, d ent.MessagesProcReq) (count int64) {
	count = 0
	if d.Pid == 1 {
		Keys := []string{"m_type", "class_id", "m_from", "m_to", "message"}
		Vals := []interface{}{d.Typer, 1, user_id, d.Id, d.Message}
		Ques := []string{"?", "?", "?", "?", "?"}

		insert := fmt.Sprintf("INSERT INTO messages(%s) VALUES(%s)", strings.Join(Keys, ","), strings.Join(Ques, ","))
		res, err := cnf.DB.Exec(insert, Vals...)
		count, _ = res.LastInsertId()
		if err != nil {
			panic(err)
		}
	}
	if d.Pid == 2 {
		dat := ent.MessagesListData{}
		query1 := "SELECT class_id,m_from,m_to FROM messages WHERE id=?"
		err1 := cnf.DB.QueryRow(query1, d.Id).Scan(&dat.Class_id, &dat.M_from_id, &dat.M_to_id)
		if err1 != nil && err1 != sql.ErrNoRows {
			panic(err1.Error())
		}

		Keys := []string{"class_id", "m_from", "m_to", "message"}
		Vals := []interface{}{*dat.Class_id, *dat.M_to_id, *dat.M_from_id, d.Message}
		Ques := []string{"?", "?", "?", "?"}
		fmt.Printf("%d %d %d %s", Vals...)
		insert := fmt.Sprintf("INSERT INTO messages(%s) VALUES(%s)", strings.Join(Keys, ","), strings.Join(Ques, ","))
		res, err := cnf.DB.Exec(insert, Vals...)
		count, _ = res.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
		go func() {
			query2 := "SELECT upAnswered(?)"
			cnf.DB.QueryRow(query2, d.Id)
		}()
	}
	if d.Pid == 3 {
		actup := `UPDATE messages SET is_deleted=1 WHERE id=?`
		res, err := cnf.DB.Exec(actup, d.Id)
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

func GetBoxBlink(user_id int32) (count int64) {
	query1 := "SELECT COUNT(id) FROM messages WHERE m_to=? AND is_readed=0"
	err1 := cnf.DB.QueryRow(query1, user_id).Scan(&count)
	if err1 != nil && err1 != sql.ErrNoRows {
		panic(err1.Error())
	}
	return
}
