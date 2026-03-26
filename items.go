package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rara-ch/wardrobie/internal/database"
)

func addHandler(s *state, args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	item, err := s.db.CreateItem(context.Background(), database.CreateItemParams{
		Type:  args[0],
		Color: sql.NullString{String: "", Valid: false},
	})
	if err != nil {
		return fmt.Errorf("could not insert item into database: %s", err)
	}

	fmt.Println("===========================================================================")
	fmt.Println("Item Inserted Successfully")
	fmt.Printf("Created At: %s\n", item.CreatedAt.Format("2006-01-02T15:04:05 -070000"))
	fmt.Printf("Type: %s\n", item.Type)

	if item.Color.Valid {
		fmt.Printf("Color: %s\n", item.Color.String)
	} else {
		fmt.Println("Color: NULL")
	}

	return nil
}
