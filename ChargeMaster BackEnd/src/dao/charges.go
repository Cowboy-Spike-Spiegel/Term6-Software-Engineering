package dao

import (
	"demo/src/consts"
	"demo/src/output"
	"fmt"
)

func ChargesQueryAll() []SchemaCharge {
	var charges []SchemaCharge

	// query in Charges
	sql := "select * from charges"
	output.Printer("Dao", sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return charges
	}

	for rows.Next() {
		var charge SchemaCharge
		err = rows.Scan(&charge.Id, &charge.Name, &charge.Mode, &charge.State,
			&charge.TotalCharge, &charge.TotalChargeTimes, &charge.TotalChargeDuration,
			&charge.TotalChargeFee, &charge.TotalChargeService, &charge.TotalFee)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		charges = append(charges, charge)
	}
	return charges
}

func ChargesQueryById(id string) (SchemaCharge, bool) {
	var charge SchemaCharge
	charge.Id = consts.NotExistId

	// query in Charges
	sql := "select * from charges where id=?"
	output.Printer("Dao", sql+" "+id)
	rows, err := db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
		return charge, false
	}

	if rows.Next() {
		err = rows.Scan(&charge.Id, &charge.Name, &charge.Mode, &charge.State,
			&charge.TotalCharge, &charge.TotalChargeTimes, &charge.TotalChargeDuration,
			&charge.TotalChargeFee, &charge.TotalChargeService, &charge.TotalFee)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return charge, true
	}
	return charge, false
}

func ChargesUpdateById(id string, attribute string, update string) bool {
	/*
		// query in Charges for exist
		sql := "select * from charges where id=?"
		output.Printer("Dao", sql+" "+id)
		rows, err := db.Query(sql, id)
		if err != nil {
			fmt.Println(err)
			return false
		}
		exist := false
		if rows.Next() {
			exist = true
		}
		if exist == false {
			return false
		}

	*/

	sql := "update charges set " + attribute + "=? where id=?"
	output.Printer("Dao", sql+" "+" "+update+" "+id)
	_, err := db.Exec(sql, update, id)
	if err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}
