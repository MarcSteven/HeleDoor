package handler

import (
	"net/http"
	"html/template"
	"os"
	"strconv"
)

//首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w, "index","", "")
}


/*
*	私有函数, handler包内其它函数可以调用以下函数
 */
func renderPanel(w http.ResponseWriter, controller string, action string, data interface{}){
	basePath, _ := os.Getwd()

	framePath := basePath + "/template/frame.html"
	headerPath := basePath + "/template/header.html"
	footerPath := basePath + "/template/footer.html"
	leftPath := basePath + "/template/" + controller + "/left_menu.html"
	var panelPath string
	if action == "" {
		panelPath = basePath + "/template/" + controller + "/panel_index.html"
	}else{
		panelPath = basePath + "/template/" + controller + "/panel_" + action + ".html"
	}
	t, error := template.ParseFiles(framePath, headerPath, footerPath, leftPath, panelPath)

	if error != nil {
		w.WriteHeader(500)
		errorPath := basePath + "/template/500.html"
		t1, _ := template.ParseFiles(errorPath)
		t1.Execute(w, "view file [" + panelPath + "] does not exist" )
	}else {
		/*
		funcMap := template.FuncMap{"active": activeLink}
		t.Funcs(funcMap)
		*/
		viewData := map[string]interface{}{"controller": controller, "action": action, "static": "/static", "data":data}
		t.ExecuteTemplate(w, "frame", viewData)
	}
}

/*
func activeLink(arg1 string, arg2 string) string{
	if arg1 == arg2 {
		return "active"
	}else{
		return "sibling"
	}
}
*/

func get(r *http.Request, name string) string {
	r.ParseForm()
	if len(r.Form[name]) > 0 {
		return r.Form[name][0]
	}else{
		return ""
	}
}

func getInt(r *http.Request, name string) int {
	g := get(r, name)
	i, err := strconv.Atoi(g)
	if err != nil {
		i = 0
	}
	return i
}
