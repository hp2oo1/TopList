package Config

type Mysql struct {
}

func (mysql Mysql) GetConfig() map[interface{}]interface{} {
	config := make(map[interface{}]interface{})
	config["source"] = "root:55c5.Fa49mysql@tcp(localhost:3306)/mine?charset=utf8mb4"
	config["driver"] = "mysql"
	return config
}
