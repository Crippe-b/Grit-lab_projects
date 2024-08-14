package sqlite

import (
	"fmt"
	"real-time-forum-base/internal/data"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func AddUserToTable(userStruct data.Registration) {
	stmt, err := data.DB.Prepare("INSERT INTO users(username, email, password, date, first_name, last_name, age, gender, session_token) values(?,?,?,?,?,?,?,?,?);")
	if err != nil {
		fmt.Println("errorstuff", err)
	}
	_, err = stmt.Exec(userStruct.Username, userStruct.Email, userStruct.Password, userStruct.RegistrationDate, userStruct.Firstname, userStruct.Lastname, userStruct.Age, userStruct.Gender, userStruct.Cokkie)
	if err != nil {
		fmt.Println("ohh yeah", err)
	}
}

func AddPostToTable(userPosts data.Post) {
	stmt, err := data.DB.Prepare("INSERT INTO posts(author_id, topic, title, body, created) values(?,?,?,?,?);")
	if err != nil {
		fmt.Println("errorstuff", err)
	}
	_, err = stmt.Exec(userPosts.AuthorId, userPosts.Categories, userPosts.Title, userPosts.Body, userPosts.Created)
	if err != nil {
		fmt.Println("ohh yeah", err)
	}
}

func AddCommentToTable(userComment data.Comment) {
	stmt, err := data.DB.Prepare("INSERT INTO posts_comments(post_id, author_id, created, body) values(?,?,?,?);")
	if err != nil {
		fmt.Println("errorstuff", err)
	}
	_, err = stmt.Exec(userComment.PostId, userComment.AuthorId, userComment.DateCreated, userComment.CommentText)
	if err != nil {
		fmt.Println("ohh yeah", err)
	}
}

func AddCategory(newCategory data.Category) {
	stmt, err := data.DB.Prepare("INSERT INTO posts_categories(category) values(?);")
	if err != nil {
		fmt.Println("errorstuff", err)
	}
	_, err = stmt.Exec(newCategory.Category)
	if err != nil {
		fmt.Println("double fuck", err)
	}
}

func CheckDataExistence(REGDATA string, dataType string) (int, string, bool) {
	var uid int
	var password string
	sqlStmt := "SELECT password,uid FROM users WHERE " + dataType + " = ?;"
	err := data.DB.QueryRow(sqlStmt, REGDATA).Scan(&password, &uid)
	if err != nil {
		return -1, "you suck", false
	}
	return uid, password, true
}

func CheckSessionExist(sessionString string) (int, error) {
	stmt := "SELECT uid FROM users WHERE session_token = ?"
	var uid int
	err := data.DB.QueryRow(stmt, sessionString).Scan(&uid)
	return uid, err
}

func UpdateSessionToken(token string, uid int) {
	stmt, err := data.DB.Prepare("UPDATE users SET session_token = ? WHERE uid = ?;")
	if err != nil {
		fmt.Println("you fucked up", err)
	}

	result, err := stmt.Exec(token, uid)
	if err != nil {
		fmt.Println("stmt execution error:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("rows affected error:", err)
	}

	fmt.Println("rows affected:", rowsAffected)
	_, err = stmt.Exec(token, uid)

	if err != nil {
		fmt.Println("exec err")
	}
}

func GetUsername(id int) string {
	stmt := "SELECT username FROM users WHERE uid = ?;"
	res := ""
	err := data.DB.QueryRow(stmt, id).Scan(&res)
	if err != nil {
		fmt.Println(err, "error in GetUsername")
	}
	return res
}

func getPostsForMainPage() []data.SPostMain {
	posts := []data.SPostMain{}
	rows, err := data.DB.Query("SELECT id,author_id,title,created,topic FROM posts;")
	if err != nil {
		fmt.Println("query issue", err)
	}
	for rows.Next() {
		var useless []byte
		var singlePost data.SPostMain
		var authorId int
		var categories string
		err := rows.Scan(&singlePost.PostId, &authorId, &singlePost.Title, &singlePost.Created, &categories)
		if err != nil {
			fmt.Println("scan issue", err)
		}
		singlePost.Categories = strings.Split(categories, "?")
		singlePost.Author = GetUsername(authorId)
		posts = append(posts, singlePost)
		_ = useless
	}
	return posts

}

func getAllCategories() []data.Topic {
	topics := []data.Topic{}
	rows, err := data.DB.Query("SELECT id,category FROM categories;")
	if err != nil {
		fmt.Println(err, "error in getAllCategories")
	}
	for rows.Next() {
		cat := data.Topic{}
		err = rows.Scan(&cat.Id, &cat.Category)
		if err != nil {
			fmt.Println(err, "error in rows.next function for categories")
		}
		topics = append(topics, cat)
	}
	return topics
}

func MainPageDataGathering() data.MainPage {
	var data data.MainPage
	data.Posts = getPostsForMainPage()
	data.Categories = getAllCategories()
	return data
}

func getAllComments(postId int) ([]data.PostComment, error) {
	res := []data.PostComment{}
	rows, err := data.DB.Query("SELECT id, author_id, created, body from posts_comments WHERE post_id = ?", postId)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var temp data.PostComment
		var uid int
		err = rows.Scan(&temp.Id, &uid, &temp.Created, &temp.CommentText)
		if err != nil {
			fmt.Println("lol u suck", err)
		}
		temp.Author = GetUsername(uid)
		res = append(res, temp)
	}
	if err != nil {
		return res, err
	}
	return res, err
}

func SinglePostDataGathering(postId int) data.SinglePost {
	res := data.SinglePost{}
	var uid int
	var body []byte
	var topics string
	stmt := "SELECT author_id, title, body, created, topic FROM posts WHERE id = ?"
	data.DB.QueryRow(stmt, postId).Scan(&uid, &res.Title, &body, &res.Created, &topics)
	res.Author = GetUsername(uid)
	res.Body = string(body)
	res.Categories = strings.Split(topics, "?")
	var err error
	res.Comments, err = getAllComments(postId)
	if err != nil {
		fmt.Println("error somewhere lol")
	}
	return res
}
