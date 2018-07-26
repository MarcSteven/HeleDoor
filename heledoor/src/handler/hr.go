package handler

import (
	"net/http"
)

//人力资源
func HrHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w, "hr", "", "")
}
