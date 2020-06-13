package mysql

import (
	"fmt"
	"github.com/lonchura/irismovie/datamodels"
	"github.com/lonchura/irismovie/repositories"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// movieMySQLRepository就是一个"MovieRepository"
type movieMySQLRepository struct {
}

// NewMovieRepository返回一个新的基于内存的movie库。
// 库的类型在我们的例子中是唯一的。
func NewMovieRepository() repositories.MovieRepository {
	return &movieMySQLRepository{}
}

func (r *movieMySQLRepository) Exec(query repositories.Query, action repositories.Query, actionLimit int, mode int) (ok bool) {
	db,_:=sql.Open("mysql","go:123123@(127.0.0.1:32777)/go") // 设置连接数据库的参数
	defer db.Close()    //关闭数据库
	err:=db.Ping()      //连接数据库
	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}

	var q string
	if actionLimit == -1 {
		q = fmt.Sprintf("select id,name,year,genre,poster from movie ORDER BY rand()")
	} else {
		q = fmt.Sprintf("select id,name,year,genre,poster from movie limit %d", actionLimit)
	}
	// TODO debug
	//fmt.Println(q)

	rows,_:=db.Query(q)
	var id int64
	var year int
	var name,genre,poster string
	var movie datamodels.Movie
	for rows.Next(){ //循环显示所有的数据
		rows.Scan(&id,&name,&year,&genre,&poster)
		// TODO debug
		//fmt.Println(id,"--",name,"--",year,"--",genre,"--",poster)

		// init movie data
		movie.ID = id
		movie.Name = name
		movie.Year = year
		movie.Genre = genre
		movie.Poster = poster

		// add movie to result
		action(movie)
	}

	return
}

// SelectMany作用相同于Select但是它返回一个切片
// 切片包含一个或多个实例
// 如果传入的参数limit<=0则返回所有
func (r *movieMySQLRepository) SelectMany(query repositories.Query, limit int) (results []datamodels.Movie) {
	r.Exec(query, func(m datamodels.Movie) bool {
		results = append(results, m)
		return true
	}, limit, 0)

	return
}