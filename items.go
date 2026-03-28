package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/rara-ch/wardrobie/internal/database"
)

type item struct {
	name     string
	brand    sql.NullString
	color    sql.NullString
	category sql.NullString
	material sql.NullString
}

func parseItem(args []string) (item, error) {
	if len(args) == 0 {
		return item{}, errors.New("there was no argument passed after command")
	}

	if string(args[0][0]) == "-" {
		return item{}, errors.New("a flag should not come directly after command")
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

	return item{
		name:     args[0],
		color:    sql.NullString{String: *color, Valid: isColorValid},
		brand:    sql.NullString{String: *brand, Valid: isBrandValid},
		material: sql.NullString{String: *material, Valid: isMaterialValid},
		category: sql.NullString{String: *category, Valid: isCategoryValid},
	}, nil
}

func addHandler(s *state, args []string) error {
	parsedItem, err := parseItem(args)
	if err != nil {
		return err
	}

	createdItem, err := s.db.CreateItem(context.Background(), database.CreateItemParams{
		Type:     parsedItem.name,
		Color:    parsedItem.color,
		Brand:    parsedItem.brand,
		Material: parsedItem.material,
		Category: parsedItem.category,
	})
	if err != nil {
		return fmt.Errorf("could not insert item into database: %s", err)
	}

	fmt.Println("===========================================================================")
	fmt.Println("Item Inserted Successfully")
	printDatabaseItem(createdItem)
	return nil
}

func getHandler(s *state, args []string) error {
	items, err := s.db.ReadItems(context.Background())
	if err != nil {
		return fmt.Errorf("could not get items from database: %s", err)
	}

	if len(items) == 0 {
		fmt.Println("There are no items")
	}

	for _, item := range items {
		fmt.Println("===========================================================================")
		printDatabaseItem(item)
	}

	return nil
}

func printDatabaseItem(item database.Item) {
	fmt.Printf("Created At: %s\n", item.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated At: %s\n", item.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Type: %s\n", item.Type)
	printNullString("Brand", item.Brand)
	printNullString("Color", item.Color)
	printNullString("Material", item.Material)
	printNullString("Category", item.Category)
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
