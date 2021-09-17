package models

import (
	"github.com/sugiantodenny01/apibook/db"
	"net/http"
)

type book struct {
	Id 		 int    `json:"id"`
	Name	 string `json:"name"`
	Author   string `json:"author"`
	Stock	 int 	`json:"stock"`
	Price    int	`json:"price"`
}

func FetchAllBook() (Response, error) {
	var obj book
	var arrObj []book
	var res Response

	con := db.CreateConn()
	sqlStatement := "SELECT * FROM books"

	rows,err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res,err
	}

	for rows.Next() {

		err = rows.Scan(&obj.Id, &obj.Name, &obj.Author, &obj.Stock,&obj.Price)
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

func StoreBook(name string, author string, stock int, price int)(Response, error)  {
	var res Response

	con := db.CreateConn()
	sqlStatement:="INSERT books (name,author,stock,price) VALUES (?,?,?,?)"
	stmt, err :=  con.Prepare(sqlStatement)

	if err !=nil {
		return res,err
	}

	result, err := stmt.Exec(name,author,stock, price)


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

func UpdateBook(id int, name string, author string, stock int, price int) (Response,error) {

	var res Response
	con := db.CreateConn()
	sqlStatement := "UPDATE books SET name = ?, author = ?, stock = ?, price = ? WHERE id = ?"
	stmt, err :=  con.Prepare(sqlStatement)

	if err !=nil {
		return res,err
	}

	result, err := stmt.Exec(name,author,stock,price,id)

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

func DeleteBook(id int)(Response, error)  {
	var res Response
	con := db.CreateConn()

	sqlStatement := "DELETE FROM books WHERE id = ?"
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
