/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-01-15 21:22:53
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-01-20 22:49:32
 */
package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"go_sqltoorm/library/resource"
	"go_sqltoorm/model"
	"go_sqltoorm/model/mysql"
)

var customConfPath string

func init() {
	initFlag()
}

func Execute() {

	flag.Parse()

	errReadConf := readConfig(customConfPath)
	if errReadConf != nil {
		fmt.Printf("%s\n", errReadConf.Error())
	}

	fmt.Printf("%+v\n", resource.FlagConf)

	var modeldb model.IModel
	switch resource.FlagConf.Type {
	case "mysql":
		modeldb = mysql.GetModel()

	}

	if modeldb == nil {
		return
	}

	modeldb.GenModel()

	// model.Generate(info)

}

func readConfig(path string) error {

	file, errPath := filepath.Abs(path)
	if errPath != nil {
		fmt.Errorf("%s", errPath.Error())
		return errPath
	}
	content, errIO := ioutil.ReadFile(file)
	if errIO != nil {
		fmt.Errorf("%s", errIO.Error())
		return errIO
	}
	err := toml.Unmarshal(content, &resource.FlagConf)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return err
	}

	return nil
}

func initFlag() {
	flag.StringVar(&customConfPath, "conf", "./data/config/flag.toml", "默认读取的flag config文件")
}
