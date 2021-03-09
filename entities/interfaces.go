package entities

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
)

type TokenClaim struct {
	ID        int64   `json:"id"`
	Identity  string  `json:"identity"`
	Title     string  `json:"title"`
	Mobile    *string `json:"mobile"`
	EMail     *string `json:"email"`
	City      *string `json:"city"`
	State     *string `json:"state"`
	Lastlogin *string `json:"ll"`
	Exp       int64   `json:"exp"`
	Iat       int64   `json:"iat"`
	*jwt.MapClaims
}

type ReturnMess struct {
	S int8   `json:"s"`
	M string `json:"m"`
}

type ReturnMessDevice struct {
	S int8   `json:"s"`
	A int8   `json:"a"`
	M string `json:"m"`
}

type ReturnPropMess struct {
	S int8        `json:"s"`
	I int64       `json:"i,omitempty"`
	T int         `json:"t,omitempty"`
	M string      `json:"m"`
	D interface{} `json:"d,omitempty"`
}

type ReturnAuthMess struct {
	S int8        `json:"s"`
	I int64       `json:"i,omitempty"`
	T int         `json:"t,omitempty"`
	C *int32      `json:"c"`
	M string      `json:"m"`
	D interface{} `json:"d,omitempty"`
}

type ReturnMessProposal struct {
	S int8   `json:"s"`
	M string `json:"m"`
	C int    `json:"c,omitempty"`
	I int64  `json:"i,omitempty"`
}

type ReqSingleSMS struct {
	Num string `json:"num"`
	Msg string `json:"msg"`
}

type SIOWarnToAllDat struct {
	NType int8   `json:"NType"`
	Data  string `json:"Data"`
}

type ReqMail struct {
	Mail string `json:"mail"`
}

type MailAgent struct {
	EMail string `json:"mail"`
}

type GetUserInfo struct {
	UserID   int64  `json:"userid,omitempty"`
	FullName string `json:"fname"`
	Identity string `json:"identity"`
}

type LoginReq struct {
	Usertype int8   `json:"t,omitempty"`
	Username string `json:"user"`
	Password string `json:"pass"`
}
type LoginDataEnduser struct {
	ID            int64   `json:"id"`
	Usr_type      int     `json:"usr_type"`
	Branch_id     *int    `json:"branch_id"`
	Teacher_id    *int    `json:"teacher_id"`
	Class_id      *int32  `json:"class_id"`
	Parent_id     *int    `json:"parent_id"`
	Identity      *string `json:"identity"`
	Title         string  `json:"title"`
	Usr_code      *string `json:"usr_code"`
	Lastlogin     *string `json:"ll"`
	Exp           int64   `json:"exp"`
	Iat           int64   `json:"iat"`
	jwt.MapClaims `json:"map,omitempty"`
}

type TeacherListData struct {
	Id          int32   `json:"id"`
	Branch_id   *int    `json:"branch_id,omitempty"`
	Branch_name *string `json:"branch_name"`
	Usr_title   *string `json:"usr_title"`
	Usr_code    *string `json:"usr_code"`
	Usr_name    *string `json:"usr_name"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone"`
	Is_act      *int    `json:"is_act"`
}

type TeacherProcReq struct {
	Pid       int    `json:"pid"`
	Id        int32  `json:"id,omitempty"`
	Usr_type  int    `json:"usr_type,omitempty"`
	Hes_code  string `json:"hes_code,omitempty"`
	Branch_id int    `json:"branch_id,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Usr_title string `json:"usr_title,omitempty"`
	Usr_code  string `json:"usr_code,omitempty"`
	Usr_name  string `json:"usr_name,omitempty"`
	Usr_pass  string `json:"usr_pass,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Lnk_zoom  string `json:"lnk_zoom,omitempty"`
	Is_act    int    `json:"is_act,omitempty"`
}

type TeachersRes struct {
	Id        int32   `json:"id"`
	Usr_type  *int    `json:"usr_type"`
	Hes_code  *string `json:"hes_code"`
	Branch_id *int    `json:"branch_id"`
	Gender    *string `json:"gender"`
	Usr_title *string `json:"usr_title"`
	Usr_code  *string `json:"usr_code"`
	Usr_name  *string `json:"usr_name"`
	Usr_pass  *string `json:"usr_pass"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone"`
	Lnk_zoom  *string `json:"lnk_zoom"`
	Is_act    *int    `json:"is_act"`
}

type StudentListData struct {
	Id          int32   `json:"id"`
	Teacher_id  *int    `json:"teacher_id"`
	Class_id    *int    `json:"class_id"`
	Class_nm    *string `json:"class_nm"`
	Parent_id   *int    `json:"parent_id"`
	Student_num *int    `json:"student_num"`
	Gender      *string `json:"gender"`
	Parent_nm   *string `json:"parent_nm"`
	Usr_title   *string `json:"usr_title"`
	Usr_name    *string `json:"usr_name"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone"`
	Is_act      *int    `json:"is_act"`
}

type StudentProcReq struct {
	Pid         int     `json:"pid"`
	Id          int32   `json:"id,omitempty"`
	Hes_code    string  `json:"hes_code,omitempty"`
	Class_id    *int    `json:"class_id,omitempty"`
	Parent_id   *int    `json:"parent_id,omitempty"`
	Student_num *int    `json:"student_num"`
	Gender      *string `json:"gender"`
	Usr_title   string  `json:"usr_title,omitempty"`
	Usr_name    string  `json:"usr_name,omitempty"`
	Usr_pass    string  `json:"usr_pass,omitempty"`
	Email       string  `json:"email,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Lnk_avatar  string  `json:"lnk_avatar,omitempty"`
	Is_act      int     `json:"is_act,omitempty"`
}

type StudentRes struct {
	Id          int32   `json:"id"`
	Usr_type    *int    `json:"usr_type"`
	Hes_code    *string `json:"hes_code"`
	Class_id    *int    `json:"class_id"`
	Parent_id   *int    `json:"parent_id"`
	Student_num *int    `json:"student_num"`
	Gender      *string `json:"gender"`
	Usr_title   *string `json:"usr_title"`
	Usr_code    *string `json:"usr_code"`
	Usr_name    *string `json:"usr_name"`
	Usr_pass    *string `json:"usr_pass"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone"`
	Is_act      *int    `json:"is_act"`
}

type ParentListData struct {
	Id        int32   `json:"id"`
	Usr_title *string `json:"usr_title"`
	Usr_name  *string `json:"usr_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Is_act    *int    `json:"is_act"`
}

type ParentProcReq struct {
	Pid        int    `json:"pid"`
	Id         int32  `json:"id,omitempty"`
	Hes_code   string `json:"hes_code,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Usr_title  string `json:"usr_title,omitempty"`
	Usr_name   string `json:"usr_name,omitempty"`
	Usr_pass   string `json:"usr_pass,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Is_act     int    `json:"is_act,omitempty"`
	Student_id int    `json:"student_id,omitempty"`
}

type ParentRes struct {
	Id        int32   `json:"id"`
	Hes_code  *string `json:"hes_code"`
	Gender    *string `json:"gender"`
	Usr_title *string `json:"usr_title"`
	Usr_code  *string `json:"usr_code"`
	Usr_name  *string `json:"usr_name"`
	Usr_pass  *string `json:"usr_pass"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone"`
	Is_act    *int    `json:"is_act"`
}

type SyllabusListData struct {
	Id           int32   `json:"id"`
	Class_id     *int    `json:"class_id"`
	Syllabus_num *int    `json:"syllabus_num"`
	Class_name   *string `json:"class_name"`
	On_date      *string `json:"on_date"`
	Is_act       *int    `json:"is_act"`
}

type SyllabusProcReq struct {
	Pid          int    `json:"pid"`
	Id           int64  `json:"id,omitempty"`
	Class_id     int64  `json:"class_id,omitempty"`
	Syllabus_num int64  `json:"syllabus_num,omitempty"`
	Bdate        string `json:"bdate,omitempty"`
	Edate        string `json:"edate,omitempty"`
	Days         int    `json:"days,omitempty"`
	L1           string `json:"l_1,omitempty"`
	L1s          int    `json:"l_1_s,omitempty"`
	L1u          string `json:"l_1_u,omitempty"`
	L2           string `json:"l_2,omitempty"`
	L2s          int    `json:"l_2_s,omitempty"`
	L2u          string `json:"l_2_u,omitempty"`
	L3           string `json:"l_3,omitempty"`
	L3s          int    `json:"l_3_s,omitempty"`
	L3u          string `json:"l_3_u,omitempty"`
	L4           string `json:"l_4,omitempty"`
	L4s          int    `json:"l_4_s,omitempty"`
	L4u          string `json:"l_4_u,omitempty"`
	L5           string `json:"l_5,omitempty"`
	L5s          int    `json:"l_5_s,omitempty"`
	L5u          string `json:"l_5_u,omitempty"`
	L6           string `json:"l_6,omitempty"`
	L6s          int    `json:"l_6_s,omitempty"`
	L6u          string `json:"l_6_u,omitempty"`
	L7           string `json:"l_7,omitempty"`
	L7s          int    `json:"l_7_s,omitempty"`
	L7u          string `json:"l_7_u,omitempty"`
	L8           string `json:"l_8,omitempty"`
	L8s          int    `json:"l_8_s,omitempty"`
	L8u          string `json:"l_8_u,omitempty"`
	L9           string `json:"l_9,omitempty"`
	L9s          int    `json:"l_9_s,omitempty"`
	L9u          string `json:"l_9_u,omitempty"`
	L10          string `json:"l_10,omitempty"`
	L10s         int    `json:"l_10_s,omitempty"`
	L10u         string `json:"l_10_u,omitempty"`
	L11          string `json:"l_11,omitempty"`
	L11s         int    `json:"l_11_s,omitempty"`
	L11u         string `json:"l_11_u,omitempty"`
	L12          string `json:"l_12,omitempty"`
	L12s         int    `json:"l_12_s,omitempty"`
	L12u         string `json:"l_12_u,omitempty"`
	On_date      string `json:"on_date,omitempty"`
	Uniq_id      string `json:"uniq_id,omitempty"`
	Is_act       int    `json:"is_act,omitempty"`
}

type SyllabusLogReq struct {
	Id           int64  `json:"id,omitempty"`
	Class_id     int64  `json:"class_id,omitempty"`
	Syllabus_num int64  `json:"syllabus_num,omitempty"`
	Bdate        string `json:"bdate,omitempty"`
	Edate        string `json:"edate,omitempty"`
	Days         int    `json:"days,omitempty"`
	L1           string `json:"l_1,omitempty"`
	L1s          int    `json:"l_1_s,omitempty"`
	L1u          string `json:"l_1_u,omitempty"`
	L2           string `json:"l_2,omitempty"`
	L2s          int    `json:"l_2_s,omitempty"`
	L2u          string `json:"l_2_u,omitempty"`
	L3           string `json:"l_3,omitempty"`
	L3s          int    `json:"l_3_s,omitempty"`
	L3u          string `json:"l_3_u,omitempty"`
	L4           string `json:"l_4,omitempty"`
	L4s          int    `json:"l_4_s,omitempty"`
	L4u          string `json:"l_4_u,omitempty"`
	L5           string `json:"l_5,omitempty"`
	L5s          int    `json:"l_5_s,omitempty"`
	L5u          string `json:"l_5_u,omitempty"`
	L6           string `json:"l_6,omitempty"`
	L6s          int    `json:"l_6_s,omitempty"`
	L6u          string `json:"l_6_u,omitempty"`
	L7           string `json:"l_7,omitempty"`
	L7s          int    `json:"l_7_s,omitempty"`
	L7u          string `json:"l_7_u,omitempty"`
	L8           string `json:"l_8,omitempty"`
	L8s          int    `json:"l_8_s,omitempty"`
	L8u          string `json:"l_8_u,omitempty"`
	L9           string `json:"l_9,omitempty"`
	L9s          int    `json:"l_9_s,omitempty"`
	L9u          string `json:"l_9_u,omitempty"`
	L10          string `json:"l_10,omitempty"`
	L10s         int    `json:"l_10_s,omitempty"`
	L10u         string `json:"l_10_u,omitempty"`
	L11          string `json:"l_11,omitempty"`
	L11s         int    `json:"l_11_s,omitempty"`
	L11u         string `json:"l_11_u,omitempty"`
	L12          string `json:"l_12,omitempty"`
	L12s         int    `json:"l_12_s,omitempty"`
	L12u         string `json:"l_12_u,omitempty"`
	On_date      string `json:"on_date,omitempty"`
	Uniq_id      string `json:"uniq_id,omitempty"`
	Is_act       int    `json:"is_act,omitempty"`
}

type Syllabus struct {
	Header SyllabusResHead `json:"header,omitempty"`
	Data   []SyllabusRes   `json:"data"`
}

type SyllabusResHead struct {
	Id           *int32  `json:"id"`
	Class_id     *int32  `json:"class_id"`
	Syllabus_num *int32  `json:"syllabus_num"`
	Bdate        *string `json:"bdate"`
	Edate        *string `json:"edate"`
	Uniq_id      *string `json:"uniq_id"`
}

type SyllabusRes2 struct {
	L1 *json.RawMessage `json:"l_1"`
	L2 *json.RawMessage `json:"l_2"`
	L3 *json.RawMessage `json:"l_3"`
	L4 *json.RawMessage `json:"l_4"`
	L5 *json.RawMessage `json:"l_5"`
	L6 *json.RawMessage `json:"l_6"`
	L7 *json.RawMessage `json:"l_7"`
	L8 *json.RawMessage `json:"l_8"`
}
type SyllabusRes struct {
	Id          *int64  `json:"id"`
	Syllabus_id *int    `json:"syllabus_id,omitempty"`
	Class_id    *int    `json:"class_id,omitempty"`
	Class_nm    *string `json:"class_nm,omitempty"`
	Days        *int    `json:"days"`
	L1          *string `json:"l_1"`
	L1s         *int    `json:"l_1_s,omitempty"`
	L1u         *string `json:"l_1_u"`
	L2          *string `json:"l_2"`
	L2s         *int    `json:"l_2_s,omitempty"`
	L2u         *string `json:"l_2_u"`
	L3          *string `json:"l_3"`
	L3s         *int    `json:"l_3_s,omitempty"`
	L3u         *string `json:"l_3_u"`
	L4          *string `json:"l_4"`
	L4s         *int    `json:"l_4_s,omitempty"`
	L4u         *string `json:"l_4_u"`
	L5          *string `json:"l_5"`
	L5s         *int    `json:"l_5_s,omitempty"`
	L5u         *string `json:"l_5_u"`
	L6          *string `json:"l_6"`
	L6s         *int    `json:"l_6_s,omitempty"`
	L6u         *string `json:"l_6_u"`
	L7          *string `json:"l_7"`
	L7s         *int    `json:"l_7_s,omitempty"`
	L7u         *string `json:"l_7_u"`
	L8          *string `json:"l_8"`
	L8s         *int    `json:"l_8_s,omitempty"`
	L8u         *string `json:"l_8_u"`
	L9          *string `json:"l_9"`
	L9s         *int    `json:"l_9_s,omitempty"`
	L9u         *string `json:"l_9_u"`
	L10         *string `json:"l_10"`
	L10s        *int    `json:"l_10_s,omitempty"`
	L10u        *string `json:"l_10_u"`
	L11         *string `json:"l_11"`
	L11s        *int    `json:"l_11_s,omitempty"`
	L11u        *string `json:"l_11_u"`
	L12         *string `json:"l_12"`
	L12s        *int    `json:"l_12_s,omitempty"`
	L12u        *string `json:"l_12_u"`
	On_date     *string `json:"on_date"`
	Uniq_id     *string `json:"uniq_id"`
	Is_act      *int    `json:"is_act"`
}

type BranchListData struct {
	Id     *int32  `json:"id"`
	Branch *string `json:"branch"`
	Is_act *int    `json:"is_act"`
}

type RollCallDetailReq struct {
	LineID int32  `json:"line_id"`
	Lesson string `json:"lesson"`
}
type RollCallCheckList struct {
	Id           *int32  `json:"id"`
	Syllabus_id  *int    `json:"syllabus_id"`
	Line_id      *int    `json:"line_id"`
	Student_id   *int    `json:"student_id"`
	Student_name *string `json:"student_name"`
	Class_nm     *string `json:"class_nm"`
	Lesson_val   *int8   `json:"lesson_val"`
	Lesson_nm    *string `json:"lesson_nm"`
}

type RollCallCheckReq struct {
	Pid         int      `json:"pid"`
	Syllabus_id int      `json:"syllabus_id"`
	Line_id     int      `json:"line_id"`
	Lesson      string   `json:"lesson"`
	Students    []string `json:"students"`
}

type RollCallCheckStudents struct {
	Pid         int    `json:"pid"`
	Syllabus_id int    `json:"syllabus_id"`
	Line_id     int    `json:"line_id"`
	Class_id    int    `json:"class_id"`
	Days        int    `json:"´days"`
	Students    string `json:"students"`
}

type BranchProcReq struct {
	Pid    int     `json:"pid"`
	Id     int32   `json:"id,omitempty"`
	Branch *string `json:"branch,omitempty"`
	Is_act int     `json:"is_act,omitempty"`
}

type ClassListData struct {
	Id        *int32  `json:"id"`
	Lang      *string `json:"lang"`
	C_num     *int    `json:"c_num"`
	C_char    *string `json:"c_char"`
	C_name    *string `json:"c_name"`
	C_total_m *int    `json:"c_total_m"`
	C_total_f *int    `json:"c_total_f"`
	C_total   *int    `json:"c_total"`
	Is_act    *int    `json:"is_act"`
}

type ClassProcReq struct {
	Pid       int     `json:"pid"`
	Id        *int32  `json:"id,omitempty"`
	Lang      *string `json:"lang,omitempty"`
	C_num     *int    `json:"c_num,omitempty"`
	C_char    *string `json:"c_char,omitempty"`
	C_total_m *int    `json:"c_total_m,omitempty"`
	C_total_f *int    `json:"c_total_f,omitempty"`
	Is_act    *int    `json:"is_act,omitempty"`
}

type MessagesRes struct {
	Id      int32  `json:"id"`
	Message string `json:"message"`
}

type MessagesListData struct {
	Id          *int32  `json:"id"`
	Class_id    *int32  `json:"class_id"`
	Class_name  *string `json:"class_name"`
	M_from_id   *int32  `json:"m_from_id"`
	M_from_name *string `json:"m_from_name"`
	M_to_id     *int32  `json:"m_to_id,omitempty"`
	M_to_name   *string `json:"m_to_name,omitempty"`
	Message     *string `json:"message"`
	On_date     *string `json:"on_date"`
	Is_readed   int     `json:"is_readed"`
	Is_answered int     `json:"is_answered"`
	Is_deleted  int     `json:"is_deleted"`
}

type MessagesProcReq struct {
	Pid     int    `json:"pid"`
	Typer   int    `json:"typer"`
	Id      int32  `json:"id"`
	Message string `json:"message"`
}

type PaginateDataStruct struct {
	Draw            string      `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	DataList        interface{} `json:"data,omitempty"`
}

type ComUserAutoList struct {
	Id         *int32  `json:"id"`
	Class_name *string `json:"class_name"`
	Usr_title  *string `json:"usr_title"`
}
type ComLessonAutoList struct {
	Id     *int32  `json:"id"`
	L_code *string `json:"l_code"`
	L_name *string `json:"l_name"`
}
type ComClassAutoList struct {
	Id         *int32  `json:"id"`
	Class_name *string `json:"class_name"`
}

type ComRollCallList struct {
	Id          *int32  `json:"id"`
	Class_nm    *string `json:"class_nm"`
	Student_num *int    `json:"student_num"`
	Gender      *string `json:"gender"`
	Usr_title   *string `json:"usr_title"`
}

type CorporateListReq struct {
	ID int64 `json:"id"`
}

type OnlyIDReq struct {
	ID int64 `json:"id"`
}
type OnlyIDDateReq struct {
	ID   int64  `json:"id"`
	Date string `json:"date"`
}

type OnlyIDClassIDReq struct {
	ID       int32 `json:"id,omitempty"`
	Class_id int32 `json:"class_id,omitempty"`
}

type OnlyLevelClassIDReq struct {
	Level    int   `json:"level,omitempty"`
	Class_id int32 `json:"class_id,omitempty"`
}

type OnlyClassIDReq struct {
	Class_id int32 `json:"class_id"`
}

type OnlyUniqIDReq struct {
	UniqID string `json:"uniq_id"`
}

type SyllabusNumReq struct {
	Syllabus_num int `json:"syllabus_num"`
}
type OnlyTyperReq struct {
	Typer int `json:"typer"`
}

type GeoReq struct {
	Id int `json:"id"`
}
type GeoCitiesRes struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
type GeoStatesRes struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type CatsReq struct {
	Id int `json:"id"`
}

type CatsRes struct {
	Id     int     `json:"id"`
	PName  string  `json:"pname,omitempty"`
	Name   string  `json:"name"`
	Color  *string `json:"color,omitempty"`
	Active string  `json:"is_act,omitempty"`
}

type GetPid struct {
	Pid int `json:"pid"`
}

/*type DashRes struct {
	Active  int `json:"active"`
	Passive int `json:"passive"`
	Waited  int `json:"waited"`
}*/

type DashReq struct {
	Typer  int `json:"typer"`
	AuthID int `json:"auth_id,omitempty"`
}
type DashRes struct {
	TeacherCnt *int `json:"teacher_cnt"`
	StudentCnt *int `json:"student_cnt"`
	ParentCnt  *int `json:"parent_cnt"`
}

type ChangePass struct {
	Id   int32  `json:"id"`
	Pass string `json:"pass"`
}

type PhoneReq struct {
	Usr_type int `json:"usr_type"`
}

type PhoneRes struct {
	Phone string `json:"phone"`
}

type SendSMSReq struct {
	Usr_type int    `json:"usr_type"`
	Message  string `json:"message"`
}

type GetStudents struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

//AUTH SMS
type NetGsmAuth struct {
	User   string `json:"usr"`
	Pass   string `json:"pas"`
	Header string `json:"hdr"`
	Num    string `json:"num"`
}

type ProfileProcReq struct {
	Hes_code string `json:"hes_code,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Lnk_zoom string `json:"lnk_zoom,omitempty"`
}

type ProfileDetRes struct {
	ID            int64   `json:"id"`
	Hes_code      *string `json:"hes_code"`
	Branch_id     *int    `json:"branch_id,omitempty"`
	Branch_nm     *string `json:"branch_name,omitempty"`
	Teacher_id    *int    `json:"teacher_id,omitempty"`
	Teacher_nm    *string `json:"teacher_name,omitempty"`
	Class_id      *int    `json:"class_id,omitempty"`
	Class_nm      *string `json:"class_nm,omitempty"`
	Parent_id     *int    `json:"parent_id,omitempty"`
	Parent_nm     *string `json:"parent_name,omitempty"`
	Student_num   *int    `json:"student_num,omitempty"`
	Gender        *string `json:"gender"`
	Usr_title     *string `json:"usr_title"`
	Usr_code      *string `json:"usr_code"`
	Identity      *string `json:"identity"`
	Email         *string `json:"email"`
	Phone         *string `json:"phone"`
	Lnk_zoom      *string `json:"lnk_zoom,omitempty"`
	Lnk_avatar    *string `json:"lnk_avatar"`
	Discontinuity *int    `json:"discontinuity"`
	Lastlogin     *string `json:"lastlogin"`
}

type SyllAppointReq struct {
	Student_id      int32 `json:"student_id"`
	Syllabus_id     int32 `json:"syllabus_id"`
	Syllabus_lesson int32 `json:"syllabus_lesson"`
	Is_join         int8  `json:"is_join"`
}
