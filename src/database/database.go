package database

// DatabaseHandler will interface all the database version struct
type DatabaseHandler interface {
	Connect(user, pass, dbName, url string) error
	IsNew() bool
}
