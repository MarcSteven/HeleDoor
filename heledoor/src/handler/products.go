package handler

import (
	"net/http"
)

//公司产品
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"cid": 0, "pid": 0}
	renderPanel(w,"products", "", data)
}

//公司产品 - 分类
func ProductsCategoryHandler(w http.ResponseWriter, r *http.Request) {
	cid := getInt(r, "cid")
	data := map[string]interface{}{"cid": cid, "pid": 0}
	renderPanel(w,"products", "category", data)
}

//公司产品 - 详情页
func ProductsDetailHandler(w http.ResponseWriter, r *http.Request) {
	cid := getInt(r, "cid")
	pid := getInt(r, "pid")
	data := map[string]interface{}{"cid": cid, "pid": pid}
	renderPanel(w,"products", "detail", data)
}
