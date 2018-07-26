package handler

import (
	"net/http"
)

//公司新闻
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"cid": 0, "aid": 0}
	renderPanel(w,"articles", "", data)
}

//公司新闻 - 分类
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	cid := getInt(r, "cid")
	data := map[string]interface{}{"cid": cid, "aid": 0}
	renderPanel(w,"articles", "category", data)
}

//公司新闻 - 详情页
func ArticlesDetailHandler(w http.ResponseWriter, r *http.Request) {
	cid := getInt(r, "cid")
	aid := getInt(r, "aid")
	data := map[string]interface{}{"cid": cid, "aid": aid}
	renderPanel(w,"articles", "detail", data)
}
