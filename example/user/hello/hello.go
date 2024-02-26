package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

type Player struct {
	tgUid uint64
}

// ? Make sure that recepe and item have the same id and name
type PlayerBaseItem struct {
	itemId   uint64
	itemName string
	recipe   PlayerBaseItemRecipe
}

// ? requiredResources is a map of resourceId to how much of that resource is needed
type PlayerBaseItemRecipe struct {
	itemId            uint64
	itemName          string
	requiredResources map[uint64]uint16
}

type PlayerBase struct {
	owner         Player
	members       []Player
	size          uint8
	maxMembers    uint8
	interiorItems []PlayerBaseItem
}

const maxMembersDefault uint8 = 5
const defaultBaseSize uint8 = 1

func main() {
	ownerSession := Player{tgUid: 1}

	defaultBaseItemRecipe := PlayerBaseItemRecipe{
		itemId:            7009,
		itemName:          "Clown car",
		requiredResources: map[uint64]uint16{6002: 35, 6003: 35},
	}

	//? Maybe name and id should be taken directly from recipe ?
	defaultBaseItem := PlayerBaseItem{
		itemId:   7009,
		itemName: "Clown car",
		recipe:   defaultBaseItemRecipe,
	}

	membersList := make([]Player, maxMembersDefault)
	defaultInteriorItemsList := make([]PlayerBaseItem, maxMembersDefault)

	defaultBase := PlayerBase{
		owner:         ownerSession,
		size:          defaultBaseSize,
		members:       membersList,
		maxMembers:    maxMembersDefault,
		interiorItems: defaultInteriorItemsList,
	}

	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	fmt.Println("ownerSession = ", ownerSession)
	fmt.Println(defaultBase, emp1)
	fmt.Println(defaultBaseItemRecipe)
	fmt.Println(defaultBaseItem)

}
