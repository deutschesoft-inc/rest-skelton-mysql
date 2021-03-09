package shared

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	cnf "main/config"
	ent "main/entities"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

func GetMess(s int8, m string) ent.ReturnMess {
	mess := ent.ReturnMess{
		S: s,
		M: m,
	}
	return mess
}

func GetMessDevice(s int8, a int8, m string) ent.ReturnMessDevice {
	mess := ent.ReturnMessDevice{
		S: s,
		A: a,
		M: m,
	}
	return mess
}

func GetPropMess(s int8, i int64, t int, m string, d interface{}) ent.ReturnPropMess {
	mess := ent.ReturnPropMess{
		S: s,
		I: i,
		T: t,
		M: m,
		D: d,
	}
	return mess
}

func GenAuthMess(s int8, i int64, t int, c *int32, m string, d interface{}) ent.ReturnAuthMess {
	mess := ent.ReturnAuthMess{
		S: s,
		I: i,
		T: t,
		C: c,
		M: m,
		D: d,
	}
	return mess
}

func GetMess2(s int8, m string, w http.ResponseWriter) {
	mess := ent.ReturnMess{
		S: s,
		M: m,
	}
	json.NewEncoder(w).Encode(mess)
}

func GetClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if parts := strings.Split(xff, ","); len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	addr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return addr
	}
	return r.RemoteAddr
}

func GetToken(r *http.Request) string {
	reqToken := r.Header.Get("authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	return splitToken[1]
}

func ParseToken(r *http.Request) jwt.MapClaims {
	unparsedToken := GetToken(r)

	keyLookupFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("h.Config.SharedSecret"), nil
	}

	token, _ := jwt.ParseWithClaims(unparsedToken, jwt.MapClaims{}, keyLookupFunc)
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

func TokenEndUser(end ent.LoginDataEnduser) (string, error) {
	signingKey := []byte(cnf.Secret)
	end.Iat = time.Now().Unix()
	end.Exp = (time.Now().Local().Add(time.Second * time.Duration(86400)).Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, end)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func HeaderToArray(header http.Header) (res []string) {
	for name, values := range header {
		for _, value := range values {
			res = append(res, fmt.Sprintf("%s: %s", name, value))
		}
	}
	return
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}

func Strtodate(s string) time.Time {
	d, _ := time.Parse("2006-01-02", s)
	return d
}

func MaskString(token string) (masked string) {
	seo := strings.Split(token, " ")
	fmt.Println(cap(seo))
	var token_array []string

	for i := 0; i < cap(seo); i++ {
		for j, _ := range seo[i] {
			if j < len(seo[i])-(len(seo[i])-1) {
				token_array = append(token_array, " "+string(seo[i][j]))
			} else {
				token_array = append(token_array, "*")
			}
		}
	}
	return strings.TrimLeft(strings.Join(token_array[:], ""), " ")
}

func GetUserInfo(userid int64) (d ent.GetUserInfo) {
	err := cnf.DB.QueryRow("SELECT `identity`, usr_title FROM auth WHERE id=?", userid).Scan(&d.Identity, &d.FullName)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}
	return
}

func ValParser(values string) (id string, value string) {
	Variable := strings.Split(values, "#")
	return Variable[0], Variable[1]
}

func GenNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(999999-100000) + 100000
}

func HashString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func SetLog(subject string, notes string) {
	go func() {
		insert := "INSERT INTO logs(subject,notes) VALUES(?,?)"
		_, err := cnf.DB.Exec(insert, subject, notes)
		if err != nil {
			panic(err)
		}
	}()
}
