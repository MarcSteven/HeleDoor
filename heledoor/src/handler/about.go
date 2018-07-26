package handler

import (
	"net/http"
)


//关于我们
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w,"about", "", "")
}

//关于我们 - 企业文化
func AboutCultureHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w,"about", "culture", "")
}

//关于我们 - 荣誉证书
func AboutCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w,"about", "credentials", "")
}

//关于我们 - 资质证书
func AboutCertHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w,"about", "cert", "")
}

//关于我们 - 领导关怀
func AboutLeaderCertHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w,"about", "leader", "")
}
