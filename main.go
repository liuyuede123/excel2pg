package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"kqqMysqlUpload/pkg"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var distPath string
var dbName string
var dbHost string
var dbUser string
var dbPassword string
var dbPort string

func main() {
	sysType := runtime.GOOS
	pkg.PathFlag = "/"
	if sysType == "windows" {
		pkg.PathFlag = "\\"
	}
	args := os.Args
	args = args[1:]
	for _, arg := range args {
		vars := strings.Split(arg, "=")
		if len(vars) < 2 {
			continue
		}
		switch vars[0] {
		case "distPath":
			distPath = vars[1]
		case "dbName":
			dbName = vars[1]
		case "dbHost":
			dbHost = vars[1]
		case "dbUser":
			dbUser = vars[1]
		case "dbPassword":
			dbPassword = vars[1]
		case "dbPort":
			dbPort = vars[1]
		}
	}
	if len(distPath) < 1 {
		fmt.Println("error 缺少参数distPath，请输入文件目录")
		os.Exit(0)
	}
	/*if len(dbName) < 1 {
		fmt.Println("error 缺少参数dbName，请输入数据库名称")
		os.Exit(0)
	}
	if len(dbHost) < 1 {
		fmt.Println("error 缺少参数dbHost，请输入数据库地址")
		os.Exit(0)
	}
	if len(dbUser) < 1 {
		fmt.Println("error 缺少参数dbUser，请输入数据库账号")
		os.Exit(0)
	}
	if len(dbPassword) < 1 {
		fmt.Println("error 缺少参数dbPassword，请输入数据库密码")
		os.Exit(0)
	}*/
	dir := distPath
	if len(dbPort) <= 0 {
		dbPort = "5432"
	}
	if len(dbName) <= 0 {
		dbName = "bi_master_test"
	}
	if len(dbHost) <= 0 {
		dbHost = "localhost"
	}
	if len(dbUser) <= 0 {
		dbUser = "superset"
	}
	if len(dbPassword) <= 0 {
		dbPassword = "superset"
	}


	// fmt.Println("当前执行路径：", pkg.GetCurrPath())

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	/*dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	pkg.Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})*/

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	pkg.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
	}


	allFileM := make(map[string]string)
	latestDateM := make(map[string]string)
	ImportTimeMap := make(map[string]int)

	if !pkg.IsDir(dir) {
		fmt.Println("error 文件路径不存在：", dir)
		os.Exit(0)
	}

	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if pkg.CheckIsExcel(path) {
				headKey, dateKey := pkg.GetHeadKeyDateKey(dir, path)
				allFileM[path] = dateKey
				if headKey != "" {
					if val, ok := latestDateM[headKey]; ok {
						if val < dateKey {
							latestDateM[headKey] = dateKey
						}
					} else {
						latestDateM[headKey] = dateKey
					}
				}

			}

			return nil
		})

	if err != nil {
		fmt.Println("error 文件处理失败：", err)
		os.Exit(0)
	}

	pkg.CreateTable()

	for k, v := range allFileM {
		headKey, dateKey := pkg.GetHeadKeyDateKey(dir, k)
		date,ok :=latestDateM[headKey]
		if !ok {
			continue
		}
		if date != dateKey {
			continue
		}

		table, _ := pkg.GetTableName(k)
		pkg.Db.Exec("delete from public.\"" + table+"\" where 数据入库日期='"+v+"'")
	}

	for k, v := range allFileM {
		headKey, dateKey := pkg.GetHeadKeyDateKey(dir, k)
		date,ok :=latestDateM[headKey]
		if !ok {
			continue
		}
		if date != dateKey {
			continue
		}

		// 剩下的就是需要处理的文件
		if _, ok = ImportTimeMap[v]; ok {
			ImportTimeMap[v]++
		} else {
			ImportTimeMap[v] = 1
		}
		err = pkg.ExecExcel(k)
		if err != nil {
			fmt.Println("error 处理Excel失败：", err)
			os.Exit(0)
		}
	}

	maxKey := pkg.GetMaxImportTimeData(ImportTimeMap)
	err = pkg.Db.Exec("insert into public.触发器事件表 (入库日期,更新表数量) values (?,?)", maxKey, ImportTimeMap[maxKey]).Error
	if err != nil {
		fmt.Println("error 插入触发器事件表报错：", err)
		os.Exit(0)
	}
	pkg.SqlExec()

	fmt.Println("success 脚本执行成功")
}
