package sandbox

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/takuya-yone/golang-api-sandbox/models"
)

func DB_conn() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	/////////////////////////////////////////////////////////////
	// rows, err := db.Query(sqlStr, articleID)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer rows.Close()

	// articleArray := make([]models.Article, 0)

	// for rows.Next() {
	// 	var article models.Article
	// 	var createdTime sql.NullTime
	// 	err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	// 	if createdTime.Valid {
	// 		article.CreatedAt = createdTime.Time
	// 	}

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		articleArray = append(articleArray, article)
	// 	}
	// }
	// fmt.Printf("%+v\n", articleArray)
	/////////////////////////////////////////////////////////////
	// articleID := 1
	// const sqlStr = `
	// select * from articles
	// where article_id = ?;
	// `

	// row := db.QueryRow(sqlStr, articleID)

	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var article models.Article
	// var createdTime sql.NullTime

	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", article)
	/////////////////////////////////////////////////////////////
	article := models.Article{
		Title:    "insert test",
		Contents: "Can in Insert data correctly?",
		UserName: "yone",
	}
	const sqlStr = `
	insert into articles(title,contents,username,nice,created_at)values(?,?,?,0,now());
	`
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

}
