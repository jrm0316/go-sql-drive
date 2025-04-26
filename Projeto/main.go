package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type User struct {
	nome  string
	idade int
}

func main() {

	var err error

	db, err = sql.Open("mysql", "root:cruzalta2025@tcp(172.17.0.2:3306)/bank")
	if err != nil {
		panic(err)
	}

	user := User{
		nome:  "Carlos",
		idade: 20,
	}

	if insertError := insertUser(user); insertError != nil {
		panic(err)
	}

	users, err := getAllUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(*user)
	}
}

func getAllUsers() ([]*User, error) {
	res, err := db.Query("SELECT * FROM tabela")
	if err != nil {
		return nil, err
	}

	users := []*User{}

	for res.Next() {
		var user User

		if err := res.Scan(&user.nome, &user.idade); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func insertUser(user User) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO tabela VALUES('%s', '%d')", user.nome, user.idade))

	if err != nil {
		return err
	}
	fmt.Println("Usuario inserido com sucesso")
	return nil
}
