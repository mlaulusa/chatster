package model

import (
	"log"

	_ "github.com/lib/pq"
)

type Message struct {
	Model
	User    string `json:"user"`
	Payload string `json:"payload"`
	Room    string `json:"room"`
	Time    string `json:"time"`
}

func SaveMessage(message *Message) error {

	insert := `INSERT INTO "Messages"("userId", "roomId", "payload", "time") VALUES($1, $2, $3, $4)`

	_, err := postgres.Exec(insert, message.User, message.Room, message.Payload, message.Time)

	if err != nil {
		return err
	}

	return nil
}

func GetAllMessages() ([]*Message, error) {

	rows, err := postgres.Query(`SELECT * FROM "Messages"`)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	results := make([]*Message, 0)

	for rows.Next() {

		var id, user, payload, room, time string

		err = rows.Scan(&id, &user, &payload, &room, &time)

		if err != nil {
			return nil, err
		}

		message := &Message{
			Model: Model{Id: id},
			User:    user,
			Payload: payload,
			Room:    room,
			Time:    time,
		}

		results = append(results, message)

	}

	return results, nil

}
