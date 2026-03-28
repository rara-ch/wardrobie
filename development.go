package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, args []string) error {
	if err := s.db.DeleteItems(context.Background()); err != nil {
		return fmt.Errorf("could not delete items from database: %e", err)
	}

	fmt.Println("All items have been deleted")
	return nil
}
