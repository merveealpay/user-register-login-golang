package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {

	userName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	//Signup

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() //sunucudan gelen requesti parse et

		userName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")
		//disaridan gelen veri dolu mu bos mu?
		userNameCheck := helper.IsEmpty(userName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if userNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintln(w, "Errorcode is -10: There is empty data")
			return
		}
		//parola kontrolu
		if pwd == pwdConfirm {
			//Veritabanına kullanici kaydet.
			fmt.Fprintln(w, "Registration successful")
		} else {
			fmt.Fprintln(w, "Password information must be the same ")
		}

	})

	//Login

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")
		//email ve parola bos birakilmamali
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintln(w, "There is empty data")
			return
		}

		dbPwd := "123"
		dbEmail := "merveealpay@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login successful")
		} else {
			fmt.Println(w, "Login failed.")
		}

	})

	http.ListenAndServe(":8080", mux)

}
