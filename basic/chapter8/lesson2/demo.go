// 包注释 - 我的包名叫demo
package demo

import "fmt"

// 常量注释 - 最大连接数
const (
	MaxConnections = 100
)

// 全局变量注释 - 数据库对象实例
var (
	Db *Database
)

// 数据库结构对象
type Database struct {
	Host string
	Port string
}

// 方法注释 - 连接到一个数据库
func (db *Database) Connect(dbName string) *Database {
	fmt.Println("模拟数据库连接...", dbName)
	return db
}

// 方法注释 - 执行SQL语句
func (db *Database) ExecuteSql(sql string) {
	fmt.Println("执行SQL语句...", sql)
}

// 函数注释 - 创建一个数据库对象实例
func NewDatabase(host, port string) *Database {
	return &Database{
		Host: host,
		Port: port,
	}
}
