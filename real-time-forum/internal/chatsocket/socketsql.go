package chatsocket

import (
	"fmt"
	"real-time-forum-base/internal/data"
	"time"
)

func GetAllUsers(username string) []data.Users {
	allUsers := []data.Users{}
	rows, err := data.DB.Query("SELECT username FROM users;")
	if err != nil {
		fmt.Println(err, "error in data Query")
	}
	for rows.Next() {
		usr := data.Users{}
		err = rows.Scan(&usr.Username)
		if err != nil {
			fmt.Println(err, "error in rows.Scan; usernames)")
		}
		hasMsg, lastMsg := hasSentMessage(username, usr.Username)
		if hasMsg {
			usr.HasSentmsg = true
			usr.LastMsg = lastMsg
		}

		if _, ok := AllClients[usr.Username]; ok {
			usr.Loggedin = true
		}
		if usr.Username == username {
			usr.You = true
		} else {
			usr.You = false
		}
		allUsers = append(allUsers, usr)
	}
	return allUsers
}

func hasSentMessage(username1, username2 string) (bool, time.Time) {
	rows, err := data.DB.Query("SELECT date FROM msg WHERE sender = ? AND reciever = ? OR sender = ? AND reciever = ?;", username1, username2, username2, username1)
	if err != nil {
		fmt.Println("whoopsie", err, "user1", username1, "user2", username2)
	}
	res := time.Time{}
	resB := false
	for rows.Next() {
		err = rows.Scan(&res)
		resB = true
		if err != nil {
			fmt.Println("Problem with rows", err)
		}
	}
	return resB, res
}
