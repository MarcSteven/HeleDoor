package handler

import (
	"net/http"
)

//商务合作
func CaseHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"cid": 0, "bid": 0}
	renderPanel(w, "case", "", data)
}

/*
//商务合作 - 分类
func CaseCategoryHandler(w http.ResponseWriter, r *http.Request) {
	cid := getInt(r, "cid")
	data := map[string]interface{}{"cid": cid, "bid": 0}
	renderPanel(w,"articles", "category", data)
}
*/

//商务合作 - 详情页
func CaseDetailHandler(w http.ResponseWriter, r *http.Request) {
	cid := getInt(r, "cid")
	bid := getInt(r, "bid")
	data := map[string]interface{}{"cid": cid, "bid": bid}
	renderPanel(w, "case", "", data)
}


// 商务合作 - 国际合作
func CaseInternationalHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w, "case", "international", "")
}


// 商务合作 - 国内合作
func CaseNationalHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w, "case", "national", "")
}

