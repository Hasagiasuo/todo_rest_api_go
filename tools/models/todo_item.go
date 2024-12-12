package models

type TodoItem struct {
	ID 			string 		`db:"id"`
	UID 		string 		`db:"uid"`
	Title 	string 		`db:"title"`
	Done 		bool 			`db:"done"`
}