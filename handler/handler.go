package handler

import (
	"SE_MIM22_WEBSHOP_LOGINSERVICE/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				_, responseErr := responseWriter.Write([]byte("{ERROR}"))
				errorHandler(responseErr)
				return
			}
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Id, Username, Password FROM users WHERE Username = ? AND Password = ?", user.Username, user.Password)
			errorHandler(err)
			var users []model.User
			for result.Next() {
				var user model.User
				err = result.Scan(&user.Id, &user.Username, &user.Password)
				errorHandler(err)
				users = append(users, user)
			}
			for _, iUser := range users {
				fmt.Println(user.Username + " " + user.Password)
				fmt.Println(iUser.Username + " " + iUser.Password)
				if iUser.Username == user.Username && iUser.Password == user.Password {
					_, responseErr := responseWriter.Write([]byte("{true}"))
					errorHandler(responseErr)
					return
				}
			}
			_, responseErr := responseWriter.Write([]byte("{false}"))
			errorHandler(responseErr)
			return
		}
		_, responseErr := responseWriter.Write([]byte("{false}"))
		errorHandler(responseErr)
		return
	default:
		_, responseErr := responseWriter.Write([]byte("THIS IS A POST REQUEST"))
		errorHandler(responseErr)
		return
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		print(err)
	}
}

func openDB() *sql.DB {
	fmt.Println("Opening DB")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	db.SetMaxIdleConns(0)
	errorHandler(err)
	return db
}
func errorHandler(err error) {
	if err != nil {
		print(err)
	}
}
