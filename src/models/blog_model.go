package models

//package entities

import (
	"database/sql"
	"entities"
	//"src/entities"
)

type BlogModel struct {
	Db *sql.DB
}

func (blogModel BlogModel) FindAll() (Blog []entities.Blog, err error) {
	rows, err := blogModel.Db.Query("SELECT * FROM Tag")
	if err != nil {
		return nil, err
	} else {
		var blogs []entities.Blog
		for rows.Next() {
			var id int64
			var title string
			var body string
			var created_at string
			var updated_at string
			err2 := rows.Scan(&id, &title, &body, &created_at, &updated_at)
			if err2 != nil {
				return nil, err2
			} else {
				blog := entities.Blog{
					Id:        id,
					Title:     title,
					Body:      body,
					CreatedAt: created_at,
					UpdatedAt: updated_at,
				}
				blogs = append(blogs, blog)
			}
		}
		return blogs, nil
	}
}

// func responseWithError(w http.ResponseWriter, code int, msg string) {
// 	responseWithJson(w, code, map[string]string{"error": msg})
// }

// func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
