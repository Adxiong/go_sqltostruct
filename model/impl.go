/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-15 23:47:22
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-15 23:49:17
 */
package model

type IModel interface {
	GenModel()
	GetDbName()
	GetPkgName()
	GetTableNames()
}
