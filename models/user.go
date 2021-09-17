package models

import (
	"database/sql"
	"fmt"
	"github.com/sugiantodenny01/apibook/db"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type user struct {
	Id 		 int    `json:"id"`
	Name	 string `json:"name"`
	Address  string `json:"address"`
	Telp	 string `json:"telp"`
	Password string	`json:"password"`
	Role 	 int    `json:"role"`
}




func FetchAllUser() (Response, error) {
	var obj user
	var arrObj []user
	var res Response

	con := db.CreateConn()
	sqlStatement := "SELECT * FROM users"

	rows,err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res,err
	}

	for rows.Next() {

		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Telp,&obj.Password, &obj.Role)
		if err != nil {
			return res,err
		}

		arrObj= append(arrObj,obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res,nil
}

func StoreUser(name string, address string, telp string, password string, role int)(Response, error)  {
	var res Response

	con := db.CreateConn()
	sqlStatement:="INSERT users (name,address,telp,password,role) VALUES (?,?,?,?,?)"
	stmt, err :=  con.Prepare(sqlStatement)

	if err !=nil {
		return res,err
	}

	result, err := stmt.Exec(name,address,telp, password, role)


	if err !=nil {
		return res,err
	}

	lastInsertedId, err := result.LastInsertId()

	if err !=nil {
		return res,err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{"last_inserted_id":lastInsertedId}

	return res, nil

}

func UpdateUser(id int, name string, address string, telp string, password string, role int) (Response,error) {

	var res Response
	con := db.CreateConn()
	sqlStatement := "UPDATE users SET name = ?, address = ?, telp = ?, role = ? WHERE id = ?"
	stmt, err :=  con.Prepare(sqlStatement)

	if err !=nil {
		return res,err
	}

	result, err := stmt.Exec(name,address,telp,password,role,id)

	if err !=nil {
		return res,err
	}

	rowsAffected, err := result.RowsAffected()

	if err !=nil {
		return res,err
	}


	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{"rows_affected":rowsAffected}

	return res, nil

}

func DeleteUser(id int)(Response, error)  {
	var res Response
	con := db.CreateConn()

	sqlStatement := "DELETE FROM users WHERE id = ?"
	stmt, err :=  con.Prepare(sqlStatement)

	if err !=nil {
		return res,err
	}

	result, err := stmt.Exec(id)

	if err !=nil {
		return res,err
	}

	rowsAffected, err := result.RowsAffected()

	if err !=nil {
		return res,err
	}


	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{"rows_affected":rowsAffected}

	return res, nil

}

func LoginUser(username, password string) (bool, *user, error) {
	var obj user
	var pwd string


	con := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE name = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Name, &obj.Address, &obj.Telp, &pwd, &obj.Role,
	)

	//fmt.Println(obj)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false,&obj, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false,&obj, err
	}

	match, err := CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false,&obj, err
	}

	return true, &obj, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func HashPassword(password string)(string, error)  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10);
	return string(bytes), err

}