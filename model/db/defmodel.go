/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-17 23:25:25
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-20 00:49:10
 */
package db

// DbInfo 数据库生成所需参数结构体
type DbInfo struct {
	PackageName string      // 包名
	DBName      string      // 数据库名
	TabList     []TableInfo // 表信息集合
}

type TableInfo struct {
	Name        string       // 表名
	Note        string       // 表注释
	SQLBuildstr string       // SQL 建表语句
	ColumnList  []ColumnInfo // 字段信息集合
}

type ColumnInfo struct {
	IsIncrement bool   // 是否自增
	IsNull      bool   // 是否空
	IsPrimary   string // 是否主键
	Index       string // 索引信息
	Comment     string // 字段注释
	Default     string //默认值
}
