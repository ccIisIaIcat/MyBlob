package main

import (
	"Project2/global"
	"fmt"
	"net/http"
	"text/template"
)

// main.go
func HomePage(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/html/homepage.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, "")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/html/contact.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, "")
}

func Monitor(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/html/Monitor.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, "")
}

func Services(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/html/Services.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, "")
}

func Diary(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/html/Diary.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 获取目录字典
	info := global.GetIndex("./data/Diary_index/index.csv", "./static/article_html/")

	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, info)
}

func Quant(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/html/quant.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	info := global.GetIndex("./data/Quant_index/index.csv", "./static/article_html/")

	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, info)
}

func test_html(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./static/article_html/article_demo.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	// 获取目录字典
	info := global.GetIndex("./data/ML_index/index.csv", "./static/article_html/")

	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, info)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/Diary", Diary)
	http.HandleFunc("/Quant", Quant)
	http.HandleFunc("/Services", Services)
	http.HandleFunc("/Monitor", Monitor)

	http.HandleFunc("/test", test_html)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP SERVER failed,err:", err)
		return
	}
}
