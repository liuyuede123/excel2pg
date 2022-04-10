package pkg

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var PathFlag string

var Db *gorm.DB

type Writer struct {
}

type TableInfo struct {
	TableSchema string
	TableName string
	OrdinalPosition string
	ColumnName string
	DataType string
	CharacterMaximumLength string
	NumericPrecision string
	NumericScale string
	IsNullable string
	ColumnDefault string
	Description string
}

func (lw Writer) Printf(format string, v ...interface{}) {

}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// ExecExcel 处理excel
func ExecExcel(filename string) error {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println("eeeeeee", err, filename)
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	// 获取表
	table, productId := GetTableName(filename)

	head := rows[0]
	if productId != "" {
		head = append(head, "商品ID")
	}
	head = append(head, "数据入库日期")

	// 判断如果没有这个表就创建一个
	if !CheckTableExist(table) {
		CreateTableAuto(table, head)
	}

	// 获取表字段
	tableFields := GetTableFields(table)
	tableColumnMap := GetTableFieldsMap(tableFields)
	for _, info := range head {
		if _,ok := tableColumnMap[strings.ToLower(info)];!ok {
			// 创建列
			Db.Exec("ALTER TABLE public."+table+" ADD COLUMN "+info+" VARCHAR(255);")
		}
	}

	importTime := GetImportTime(filename)

	rows = rows[1:]
	num := len(rows)
	for num > 0 {
		limit := 100
		if num < 100 {
			limit = num
		}
		data := rows[:limit]

		// 批量插入数据
		insertSql := "insert into public.\"" + table + "\" (\"" + strings.ToLower(strings.Join(head, "\",\"")) + "\") values "
		newRows := make([]string, 0)

		for _, row := range data {
			for i := 0; i < len(row); i++ {
				row[i] = strings.ReplaceAll(row[i], "'", "")
			}
			fillRow := make([]string, len(head)-1)
			if len(row) == 0 {
				continue
			}
			if len(row) < len(fillRow) {
				for k, s := range row {
					fillRow[k] = s
				}
			} else {
				fillRow = row[:len(fillRow)]
			}
			if productId != "" {
				fillRow[len(fillRow)-1] = productId
			}
			fillRow = append(fillRow, importTime)
			newRow := " ('" + strings.Join(fillRow, "','") + "')"
			newRows = append(newRows, newRow)
		}

		insertSql += strings.Join(newRows, ",")
		insertSql += " ;"
		err = Db.Exec(insertSql).Error
		if err != nil {
			fmt.Println("error 写入sql报错：", err)
			os.Exit(0)
		}

		rows = rows[limit:]
		num -= 100
	}

	return nil
}

func GetTableFieldsMap(columns []TableInfo) (res map[string]struct{}) {
	res = make(map[string]struct{})
	for _, s := range columns {
		res[s.ColumnName] = struct{}{}
	}
	return
}

func GetTableFields(tableName string) []TableInfo {
	var result []TableInfo
	Db.Raw(`select
col.table_schema,
col.table_name,
col.ordinal_position,
col.column_name,
col.data_type,
col.character_maximum_length,
col.numeric_precision,
col.numeric_scale,
col.is_nullable,
col.column_default,
des.description
from
information_schema.columns col left join pg_description des on
col.table_name::regclass = des.objoid
and col.ordinal_position = des.objsubid
where
table_schema = 'public'
and table_name = ?
order by
ordinal_position`, tableName).Scan(&result)
	return result
}

func GetTableName(filename string) (table string, productId string) {
	if strings.Contains(filename, "ExportOrderDetailList") {
		return "商家中心_订单表商品", ""
	}
	if strings.Contains(filename, "ExportOrderList") {
		return "商家中心_订单表明细", ""
	}
	files := strings.Split(filename, PathFlag)
	files = strings.Split(files[len(files)-1], "-")
	table = strings.Join(files[:len(files)-1], "-")
	table = strings.ReplaceAll(table, " ", "")
	table = strings.ReplaceAll(table, "-", "_")
	tables := strings.Split(table, "_")
	for _, s := range tables {
		if IsDigit(s) {
			if tables[len(tables)-1] == "销售分析商品1" {
				productId = s
			}
			table = strings.ReplaceAll(table, "_"+s, "")
		}
	}
	table = strings.Trim(table, "_")
	table = strings.ToLower(table)

	return table, productId
}

func GetImportTime(filename string) string {
	files := strings.Split(filename, "-")
	res := strings.ReplaceAll(files[len(files)-1], ".xlsx", "")
	res = strings.ReplaceAll(res, ".xls", "")
	if IsDigit(res) {
		return res
	}
	return ""
}

func GetMaxImportTimeData(importTimeM map[string]int) string {
	maxData := ""
	for i, _ := range importTimeM {
		if i > maxData {
			maxData = i
		}
	}
	return maxData
}

func CreateTable() {
	// Db.Exec(config.CreateSql)

	sqlFile := GetCurrPath() + PathFlag + "sql"
	if !IsFile(sqlFile) {
		fmt.Println("error sql文件不存在：", sqlFile)
		os.Exit(0)
	}
	sqls, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		fmt.Println("error 读取sql文件失败：", "sql所在路径路径："+ GetCurrPath(), sqlFile, err)
		os.Exit(0)
	}

	sqlsli := strings.Split(string(sqls), ";")
	for _, s := range sqlsli {
		err = Db.Exec(s).Error
		if err != nil {
			fmt.Println("error 创建表报错：", err)
			os.Exit(0)
		}
	}
	fmt.Println("sql文件执行完成")
}

func SqlExec() {
	// Db.Exec(config.CreateSql)

	sqlFile := GetCurrPath() + PathFlag + "sql_exec"
	if !IsFile(sqlFile) {
		fmt.Println("error sql_exec文件不存在：", sqlFile)
		os.Exit(0)
	}
	sqls, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		fmt.Println("error 读取sql_exec文件失败：", "当前执行路径："+ GetCurrPath(), sqlFile, err)
		os.Exit(0)
	}

	err = Db.Exec(string(sqls)).Error
	if err != nil {
		fmt.Println("error 执行sql_exec报错：", err)
		os.Exit(0)
	}
	fmt.Println("sql_exec文件执行完成")
}

func IsSingleDigit(data string) bool {
	digit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, item := range digit {
		if data == item {
			return true
		}
	}
	return false
}

func IsDigit(data string) bool {
	if data == "" {
		return false
	}
	for _, item := range data {
		if IsSingleDigit(string(item)) {
			continue
		} else {
			return false
		}
	}
	return true
}

func CheckIsExcel(path string) bool {
	return IsFile(path) && filepath.Ext(path) == ".xlsx" || IsFile(path) && filepath.Ext(path) == ".xls"
}

func GetHeadKeyDateKey(rootPath string, path string) (string, string) {
	prePath := strings.Replace(path, filepath.Ext(path), "", 1)
	prePaths := strings.Split(prePath, "-")
	dateKey := prePaths[len(prePaths)-1]
	headKey := prePaths[0]
	headKey = strings.ReplaceAll(headKey, rootPath+PathFlag, "")
	headKeys := strings.Split(headKey, PathFlag)
	headKey = rootPath + PathFlag + headKeys[0]
	return headKey, dateKey
}

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func CheckTableExist(tableName string) bool {
	var count int64
	Db.Raw("select * from pg_class where relname = '"+tableName+"'").Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func CreateTableAuto(tableName string, fields []string)  {
	if len(fields) == 0 {
		return
	}
	sqlField := make([]string, 0)
	sql := "CREATE TABLE IF NOT EXISTS public.\""+tableName+"\" ( "
	for _, field := range fields {
		sqlField = append(sqlField, "\"" + strings.ToLower(field) + "\" varchar(255)")
	}
	sql += strings.Join(sqlField, ",")
	sql += " );"
	err := Db.Exec(sql).Error
	if err != nil {
		fmt.Println("error 判断无表创建sql失败:", err)
		os.Exit(0)
	}
}
