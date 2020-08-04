package models

import (
	"../db"
	"net/http"
)

type Item struct {
	Id 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Price 	int 	`json:"price"`
}

func FetchAllProducts()(Response, error) {
	var obj Item
	var arrObj []Item
	var res Response

	con := db.CreateCon()

	queryString := "SELECT * FROM items"

	rows, err := con.Query(queryString)
	defer rows.Close()
	if err!=nil {
		return res, err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Price)
		if err != nil{
			return res, err
		}
		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Success get data"
	res.Data = arrObj

	return res, nil
}

func FetchById(id int)(Response, error){
	var res Response
	var obj Item
	con := db.CreateCon()

	queryStatement := "SELECT * FROM items where id = ?"

	rows := con.QueryRow(queryStatement, id)

	_ = rows.Scan(&obj.Id, &obj.Name, &obj.Price)

	if obj.Id == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Item not found"
		res.Data = nil

		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func CreateItems(name string, price int)(Response, error){
	var res Response

	con := db.CreateCon()

	queryStatement := "INSERT INTO items(name, price) values (?,?)"

	stmt, err := con.Prepare(queryStatement)

	if err!=nil{
		return res, err
	}

	result, err := stmt.Exec(name, price)
	if err!=nil{
		return res, err
	}
	// variable lastInsertedId akan terisi id item yang dibuat
	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Success create item"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertedId,
	}

	return res, nil
}


func UpdateItem(id int, name string, price int)(Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE items SET name = ?, price = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, price, id)

	if err!=nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()


	res.Status = http.StatusOK
	res.Message = "Success update items"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteItem(id int)(Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM items WHERE id = ?"
	rows, err := con.Prepare(sqlStatement)

	if err!=nil{
		return res, err
	}

	_, err = rows.Exec(id)

	if err!=nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success delete items"
	res.Data = map[string]int{
		"item_id" : id,
	}

	return res, nil
}