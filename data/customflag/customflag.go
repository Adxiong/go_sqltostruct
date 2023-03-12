/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-15 21:41:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-20 23:07:19
 */
package customflag

type FlagConf struct {
	Host      string
	Port      uint64
	User      string
	Password  string
	DataBase  string
	Outdir    string
	Type      string
	TableList []string
}
