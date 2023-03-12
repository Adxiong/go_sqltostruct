/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-20 22:34:49
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-20 22:41:01
 */
package mysql

import (
	"fmt"
	"go_sqltoorm/library/resource"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	conf := resource.FlagConf
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DataBase), // DSN data source name
		DefaultStringSize:         256,                                                                                                                                     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                                                    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                                                    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                                                    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                                                   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		fmt.Errorf("%s", err)
	}

	return db
}
