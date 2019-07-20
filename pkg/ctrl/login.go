package ctrl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoginCtrl interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type loginCtrl struct {
}

func NewLoginCtrl() LoginCtrl {
	return &loginCtrl{}
}

func (c *loginCtrl) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth.Login()")
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(body)
}
