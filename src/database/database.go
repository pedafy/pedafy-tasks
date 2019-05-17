package database

// DatabaseHandler will interface all the database version struct
type DatabaseHandler interface {
	Connect(user, pass, dbName, url string) error
	IsNew() bool

	GetAllStatus() ([]Status, error)
	GetStatusByID(ID int) (Status, error)
	GetStatusByName(name string) (Status, error)

	GetAllTasks() ([]Tasks, error)
	GetAllTasksByOrder(order string) ([]Tasks, error)
	GetAllTasksByFilter(filter, value string) ([]Tasks, error)
	NewTask(task Tasks) (Tasks, error)
	ModifyTask(task Tasks, ID int) (Tasks, error)
	ArchiveTask(ID int) (Tasks, error)
}
