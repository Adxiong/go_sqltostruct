/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-15 23:40:04
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-15 23:55:38
 */
package mysql

import "go_sqltoorm/model"

var MysqlModel mysqlModel

type mysqlModel struct{}

func GetModel() model.IModel {
	return &MysqlModel
}

func (m *mysqlModel) GenModel() {

}

func (m *mysqlModel) GetDbName() {}

func (m *mysqlModel) GetPkgName() {}

func (m *mysqlModel) GetTableNames() {}
