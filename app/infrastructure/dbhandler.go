package infrastructure

type DBHandler struct {
}

func GetDBHandler() *DBHandler {
	dbHandler := new(DBHandler)
	return dbHandler
}
