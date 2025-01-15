package types

import (
	"fmt"
)

func (leaderboard Leaderboard) Validate() error {
	// Check for duplicated player address in winners
	winnerInfoIndexMap := make(map[string]struct{})

	for index, elem := range leaderboard.Winners {
		if _, ok := winnerInfoIndexMap[elem.Address]; ok {
			return fmt.Errorf("duplicated address %s at index %d", elem.Address, index)
		}
		winnerInfoIndexMap[elem.Address] = struct{}{}
	}
	return nil
}
