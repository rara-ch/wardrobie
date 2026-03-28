package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/rara-ch/wardrobie/internal/database"
)

func addHandler(s *state, args []string) error {
	// TODO: implement tests
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	itemFlags := flag.NewFlagSet("item", flag.ExitOnError)
	brand := itemFlags.String("brand", "", "sets the brand of the item")
	color := itemFlags.String("color", "", "sets the color of the item")
	material := itemFlags.String("material", "", "sets the material of the item")
	category := itemFlags.String("category", "", "sets the category of the item")
	itemFlags.Parse(args[1:])

	isBrandValid := isEmpty(*brand)
	isColorValid := isEmpty(*color)
	isMaterialValid := isEmpty(*material)
	isCategoryValid := isEmpty(*category)

	item, err := s.db.CreateItem(context.Background(), database.CreateItemParams{
		Type:     args[0],
		Color:    sql.NullString{String: *color, Valid: isColorValid},
		Brand:    sql.NullString{String: *brand, Valid: isBrandValid},
		Material: sql.NullString{String: *material, Valid: isMaterialValid},
		Category: sql.NullString{String: *category, Valid: isCategoryValid},
	})
	if err != nil {
		return fmt.Errorf("could not insert item into database: %s", err)
	}

	fmt.Println("===========================================================================")
	fmt.Println("Item Inserted Successfully")
	fmt.Printf("Created At: %s\n", item.CreatedAt.Format("2006-01-02T15:04:05 -070000"))
	fmt.Printf("Type: %s\n", item.Type)
	printNullString("Brand", item.Brand)
	printNullString("Color", item.Color)
	printNullString("Material", item.Material)
	printNullString("Category", item.Category)

	return nil
}

func isEmpty(s string) bool {
	if s != "" {
		return true
	} else {
		return false
	}
}

func printNullString(name string, value sql.NullString) {
	if value.Valid {
		fmt.Printf("%s: %s\n", name, value.String)
	} else {
		fmt.Printf("%s: NULL\n", name)
	}
}
