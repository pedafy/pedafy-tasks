package database

// DatabaseHandler will interface all the database version struct
type DatabaseHandler interface {
	Connect(user, pass, dbName, url string) error
	IsNew() bool

	GetAllStatus() ([]Status, error)
	GetStatusByID(ID int) (Status, error)
	GetStatusByName(name string) (Status, error)
}
