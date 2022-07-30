package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	CinnectToServer(ctx)

}

func DataBaseW(ctx context.Context, login string, password string) {
	connStr := "user=postgres dbname=golang password=1111 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	select {
	case <-ctx.Done():
		fmt.Println("Time ERr")
		return
	}
	//	time.Sleep(7 * time.Second)

	if err != nil {
		panic(err)
	}

	//fd, error := os.ReadFile("./schema/insert.sql")
	//if error != nil {
	//	panic(error.Error())
	//}
	//
	////str := string(fd)
	////insert := fmt.Sprintf(str, "user_login_02", "mega_hard_password")
	insert := fmt.Sprintf("INSERT INTO registration_data (user_login,user_pass) VALUES ('%s','%s')", login, password)
	fmt.Println(insert)

	_, err = db.Exec(insert)
	if err != nil {
		panic(err)
	}
	//db.Prepare(insert)
	//	defer stmt.Close()

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	// конец полключения 1
}

func CinnectToServer(cxt context.Context) {
	http.HandleFunc("/", mainPage)
	port := ":9090"
	fmt.Println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

var tpl = template.Must(template.ParseFiles("./HTTP_Files/serv.html"))

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Регистрация<h1>")) // передача текста
	tpl.Execute(w, nil)
	value_1 := (r.FormValue("value 1"))
	value_2 := (r.FormValue("value 2"))
	fmt.Println(value_1, value_2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	DataBaseW(ctx, value_1, value_2)

	//r.FormValue("result") = result

}
