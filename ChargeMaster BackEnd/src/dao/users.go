package dao

import (
	"demo/src/consts"
	"demo/src/output"
	"fmt"
)

func UsersQueryById(id string) (SchemaUser, bool) {
	var user SchemaUser
	user.Id = consts.NotExistId

	// query in Users
	sql := "select * from users where id=?"
	output.Printer("Dao", sql+" "+id)
	rows, err := db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Auth, &user.Account, &user.Password, &user.Name, &user.CarArray, &user.Balance)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return user, true
	}
	return user, false
}

func UsersQueryByAccount(account string) (SchemaUser, bool) {
	var user SchemaUser
	user.Id = consts.NotExistId

	// query in Users
	sql := "select * from users where account=?"
	output.Printer("Dao", sql+" "+account)
	rows, err := db.Query(sql, account)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	// judge login state
	if rows.Next() {
		// store a instance into user
		err = rows.Scan(&user.Id, &user.Auth, &user.Account, &user.Password, &user.Name, &user.CarArray, &user.Balance)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return user, true
	}
	return user, false
}

func UsersQueryByAP(account string, password string) (SchemaUser, bool) {
	var user SchemaUser
	user.Id = consts.NotExistId

	// query in Users
	sql := "select * from users where account=? and password=?"
	output.Printer("Dao", sql+" "+account+" "+password)
	rows, err := db.Query(sql, account, password)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Auth, &user.Account, &user.Password, &user.Name, &user.CarArray, &user.Balance)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return user, true
	}
	return user, false
}

func UsersQueryAdminByAP(account string, password string) (SchemaUser, bool) {
	var user SchemaUser
	user.Id = consts.NotExistId

	// query in Users
	sql := "select * from users where auth=1 and account=? and password=?"
	output.Printer("Dao", sql+" "+account+" "+password)
	rows, err := db.Query(sql, account, password)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Auth, &user.Account, &user.Password, &user.Name, &user.CarArray, &user.Balance)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return user, true
	}
	return user, false
}

func UsersQueryByIdP(id string, password string) (SchemaUser, bool) {
	var user SchemaUser
	user.Id = consts.NotExistId

	// query in Users
	sql := "select * from users where id=? and password=?"
	output.Printer("Dao", sql+" "+id)
	rows, err := db.Query(sql, id, password)
	if err != nil {
		fmt.Println(err)
		return user, false
	}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Auth, &user.Account, &user.Password, &user.Name, &user.CarArray, &user.Balance)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return user, true
	}
	return user, false
}

func UsersInsert(account string, password string) (SchemaUser, bool) {
	// query for exist
	user, exist := UsersQueryByAccount(account)
	if exist == true {
		user.Id = consts.RepeatId
		return user, false
	}

	// insert
	sql := "insert into users (" +
		AttributeUser["Auth"] + ", " +
		AttributeUser["Account"] + ", " +
		AttributeUser["Password"] + ", " +
		AttributeUser["Name"] + ", " +
		AttributeUser["CarArray"] + ", " +
		AttributeUser["Balance"] +
		") values (0,?,?,'未命名','',0)"
	output.Printer("Dao", sql+" "+account+" "+password)

	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(account, password); err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return user, false
	}
	return UsersQueryByAccount(account)
}

func UsersUpdate(id string, attribute string, update string) bool {
	sql := "update users set " + attribute + "=? where id=?"
	output.Printer("Dao", sql+" "+" "+update+" "+id)
	_, err := db.Exec(sql, update, id)
	if err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}
