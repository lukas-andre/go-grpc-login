package dependencies

import "gorm.io/gorm"

type DaoDependency func(*DAO)

type DAO struct {
	Dao *gorm.DB
}

func NewDaoDependency(deps ...DaoDependency) *DAO {
	d := &DAO{}
	for _, dep := range deps {
		dep(d)
	}
	return d
}

func WithDao(dao *gorm.DB) DaoDependency {
	return func(d *DAO) {
		d.Dao = dao
	}
}
