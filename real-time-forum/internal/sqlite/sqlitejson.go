package sqlite

import (
	"fmt"
	"real-time-forum-base/internal/data"
)

func GetallMessages(user1, user2 string) []data.MsgJ {
	messages := []data.MsgJ{}
	rows, err := data.DB.Query("SELECT sender,msg,date FROM msg WHERE sender = ? AND reciever = ? OR sender = ? AND reciever = ?", user1, user2, user2, user1)
	if err != nil {
		fmt.Println(err, "error in data.Query; getallmessages")

	}
	for rows.Next() {
		msg := data.MsgJ{}
		err = rows.Scan(&msg.Sender, &msg.Msg, &msg.Date)
		if err != nil {
			fmt.Println(err, "error in rows.Scan messages")

		}
		messages = append(messages, msg)
	}

	return messages
}
