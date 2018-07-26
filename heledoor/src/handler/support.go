package handler

import (
	"net/http"
)

//技术支持
func SupportHandler(w http.ResponseWriter, r *http.Request) {
	renderPanel(w, "support", "", "")
}
