package demo

func ExampleDatabase() {
	db := NewDatabase("127.0.0.1", "3306")
	db.Connect("demo").ExecuteSql("select * from demo")

	// Output:
	// 模拟数据库连接... demo
	// 执行SQL语句... select * from demo
}
