/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-15 23:47:22
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-20 23:15:37
 */
package model

type IModel interface {
	GenModel()
	GetDbName() string
	GetPkgName() string
	GetTableNames() (*[]string, error)
}
