package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/doLogin", doLogin)
	http.ListenAndServe("localhost:9999", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	message := ""
	if req.Method == "POST" {
		/*
			req.ParseForm()
			if checkUsername(req.Form["username"][0]) {
				fmt.Fprint(rw, "用户名:", req.Form["username"][0], ",密码:", req.Form["password"][0])
				return
			} else {
				message = "不符合要求"
			}
		*/
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("upfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
		message = "success"
	}

	t, _ := template.ParseFiles("get.html")
	t.Execute(rw, message)

}

// func doLogin(rw http.ResponseWriter, req *http.Request) {
// 	req.ParseForm()

// 	if checkUsername(req.Form["username"][0]) {
// 		fmt.Fprint(rw, "用户名:", req.Form["username"][0], ",密码:", req.Form["password"][0])
// 	} else {
// 		t, _ := template.ParseFiles("get.html")
// 		t.Execute(rw, "用户名不合法")
// 	}

// }

/*
func checkUsername(username string) bool {
	if len(username) >= 6 && len(username) <= 16 {
		return true
	}
	return false
}
*/
