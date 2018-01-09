package models

import (
	"log"
	"github.com/graphql-go/graphql"
)

type User struct {
	Id      int    `json:"id",db:"id"`
	Name    string `json:"name",db:"name"`
	Email   string `json:"email",db:"email"`
	Phone   string `json:"phone",db:"phone"`
	Deleted int    `json:"deleted",db:"deleted"`
}

// define custom GraphQL ObjectType `UserType` for our Golang struct `User`
// Note that
// - the fields in our UserType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"deleted": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func AllUsers() ([]*User, error) {
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs := make([]*User, 0)

	for rows.Next() {
		usr := new(User)
		err := rows.Scan(&usr.Name, &usr.Email, &usr.Phone, &usr.Deleted)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usrs, nil
}

func NewUser(usr User) (error) {
	stmt, err := db.Prepare("INSERT INTO users(name, email, phone, deleted) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(usr.Name, usr.Email, usr.Phone, usr.Deleted)
	if err != nil {
		return err
	}
	log.Printf("Result = %+v", res)
	return nil
}
