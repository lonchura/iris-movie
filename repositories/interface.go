package repositories

import (
	"github.com/lonchura/irismovie/datamodels"
	"github.com/lonchura/irismovie/repositories/sql"
)

// Query代表一种“访客”和它的查询动作。
type Query func(datamodels.Movie) bool

type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select([]*sql.Condition) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)

	//InsertOrUpdate(movie datamodels.Movie) (updatedMovie datamodels.Movie, err error)
	//Delete(query Query, limit int) (deleted bool)
}