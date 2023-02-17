package constantx

const (
	MYSQL_Dsn string = "titok:tdk@tcp(127.0.0.1:3306)/titok?charset=utf8mb4&parseTime=True&loc=Local"
)

// config
const (
	MYSQL_SkipDefaultTransaction bool = true // close default tx
	MYSQL_PrepareStmt            bool = true // cache precompile sentence
)
