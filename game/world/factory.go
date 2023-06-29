package world

func Create() *Location {
	kitchen := Location{
		ItemContainerImpl: ItemContainerImpl{_name: "кухня"},
		Inspection:        "ты находишься на кухне, %s, %s. %s",
		Welcome:           "кухня, ничего интересного. %s",
	}
	hall := Location{ItemContainerImpl: ItemContainerImpl{_name: "коридор"}, Welcome: "ничего интересного. %s"}
	room := Location{ItemContainerImpl: ItemContainerImpl{_name: "комната", _emptyMsg: "пустая комната"}, Welcome: "ты в своей комнате. %s"}
	outdoor := Location{ItemContainerImpl: ItemContainerImpl{_name: "улица"}, Welcome: "на улице весна. %s"}

	hall.AddDoor(&kitchen)
	hall.AddDoor(&room)
	exitDoor := hall.AddDoor(&outdoor)

	var exitDoorUseful Useful = exitDoor
	hall.AddUseful(&exitDoorUseful)
	outdoor.AddUseful(&exitDoorUseful)

	exitDoor.first.name = "домой"
	exitDoor.Locked = true
	var exitKey Item = &Key{secret: exitDoor.secret}

	var backpack Item = &Backpack{&ItemContainerImpl{_showTmpl: "в рюкзаке"}}
	var book Item = &DummyItem{"конспекты", false, true, true}
	var roomTable = &Furniture{ItemContainerImpl{_name: "стол", _showTmpl: "на столе: %s"}}
	roomTable.AddItem(&exitKey)
	roomTable.AddItem(&book)
	roomChair := &Furniture{ItemContainerImpl{_name: "стул", _showTmpl: "на стуле: %s"}}
	roomChair.AddItem(&backpack)

	var roomTableItem Item = roomTable
	var roomChairItem Item = roomChair
	room.AddItem(&roomTableItem)
	room.AddItem(&roomChairItem)

	var tea Item = &DummyItem{"чай", false, false, true}
	kitchenTable := &Furniture{ItemContainerImpl{_name: "стол", _showTmpl: "на столе: %s"}}
	kitchenTable.AddItem(&tea)
	var kitchenTableItem Item = kitchenTable
	kitchen.AddItem(&kitchenTableItem)

	return &kitchen
}
