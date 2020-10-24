package model

import (
	"log"
)

type Room struct {
	Model
	Name string `json:"name"`
}

func SaveRoom (room *Room) error {

	_, err := postgres.Exec(`INSERT INTO "Rooms"("name") values($1)`, room.Name)

	if err != nil {
		return err
	}

	return nil
}

func GetAllRooms () ([]*Room, error) {

	rows, err := postgres.Query(`SELECT * FROM "Rooms"`)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	results := make([]*Room, 0)

	for rows.Next() {

		var id, name string

		err = rows.Scan(&id, &name)

		if err != nil {
			return nil, err
		}

		room := &Room{
			Model: Model{Id: id},
			Name:  name,
		}

		results = append(results, room)

	}

	return results, nil
}
