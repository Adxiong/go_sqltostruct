/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-15 23:40:04
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-20 23:16:32
 */
package mysql

import (
	"fmt"
	"go_sqltoorm/library/resource"
	"go_sqltoorm/model"
	"strings"
)

var MysqlModel mysqlModel

type mysqlModel struct{}

func GetModel() model.IModel {
	return &MysqlModel
}

func (m *mysqlModel) GenModel() {
	tableList, errTableList := m.GetTableNames()

	if errTableList != nil {
		fmt.Errorf("%s\n", errTableList.Error())
	}
	fmt.Printf("%+v\n", tableList)
}

func (m *mysqlModel) GetDbName() string {
	return resource.FlagConf.DataBase
}

func (m *mysqlModel) GetPkgName() string {
	return strings.ToLower(resource.FlagConf.DataBase)
}

func (m *mysqlModel) GetTableNames() (*[]string, error) {
	db := GetDb()
	res := make([]string, 0)

	if len(resource.FlagConf.TableList) > 0 {
		res = append(res, resource.FlagConf.TableList...)
		return &res, nil
	}

	row, err := db.Raw("show tables;").Rows()
	if err != nil {
		fmt.Errorf("show tables %s\n", err)
	}

	for row.Next() {
		var str string
		row.Scan(&str)
		res = append(res, str)
	}
	fmt.Printf("show tables result %+v\n", res)
	return &res, nil
}

func (m *mysqlModel) GetTableInfo() {
	// show create table;
}
