package database

type Database interface {
	Start() error
}

type Gorm struct {
	Database
}

func (db *Gorm) Start() error {
	return nil
}
