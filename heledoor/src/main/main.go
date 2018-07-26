package main

import (
	"net/http"
	"config"
	"handler"
	"os"
)

func main() {

	basePath, _ := os.Getwd()
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(basePath + config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//首页
	mux.HandleFunc("/", handler.IndexHandler)

	//关于我们
	mux.HandleFunc("/about.html", handler.AboutHandler)
	mux.HandleFunc("/about/culture.html", handler.AboutCultureHandler)
	mux.HandleFunc("/about/credentials.html", handler.AboutCredentialsHandler)
	mux.HandleFunc("/about/cert.html", handler.AboutCertHandler)
	mux.HandleFunc("/about/leader.html", handler.AboutLeaderCertHandler)

	//公司新闻
	mux.HandleFunc("/articles.html", handler.ArticlesHandler)
	mux.HandleFunc("/articles/category.html", handler.ArticlesCategoryHandler)
	mux.HandleFunc("/articles/detail.html", handler.ArticlesDetailHandler)

	//公司产品
	mux.HandleFunc("/products.html", handler.ProductsHandler)
	mux.HandleFunc("/products/category.html", handler.ProductsCategoryHandler)
	mux.HandleFunc("/products/detail.html", handler.ProductsDetailHandler)

	//人力资源
	mux.HandleFunc("/hr.html", handler.HrHandler)

	//技术支持
	mux.HandleFunc("/support.html", handler.SupportHandler)

	//商务合作
	mux.HandleFunc("/case.html", handler.CaseHandler)
	mux.HandleFunc("/case/detail.html", handler.CaseDetailHandler)
	mux.HandleFunc("/case/international.html", handler.CaseInternationalHandler)
	mux.HandleFunc("/case/national.html", handler.CaseNationalHandler)

	server := &http.Server{
		Addr:		"0.0.0.0:8080",
		Handler:	mux,
	}
	server.ListenAndServe()
}
