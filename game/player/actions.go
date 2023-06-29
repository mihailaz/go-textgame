package player

type Action func(*Player, []string) (string, bool)

var Actions = map[string]Action{
	`осмотреться`: inspectAction,
	`идти`:        goAction,
	`применить`:   useAction,
	`взять`:       takeAction,
	`надеть`:      wearAction,
	`инвентарь`:   inventoryAction,
}

func inspectAction(player *Player, args []string) (string, bool) {
	if len(args) > 0 {
		return "неизвестная команда", false
	}
	return player.ShowInspection(), true
}

func goAction(player *Player, args []string) (string, bool) {
	if len(args) != 1 {
		return "неизвестная команда", false
	}
	door, ok := player.CurrentLocation.FindDoorTo(args[0])
	if !ok {
		return "нет пути в " + args[0], false
	}
	way, ok := door.ComeIn(player.CurrentLocation)
	if !ok {
		return "дверь закрыта", false
	}
	return player.Move(way.Location), true
}

func useAction(player *Player, args []string) (string, bool) {
	if len(args) != 2 {
		return "неизвестная команда", false
	}
	item, idx, _ := player.FindItem(args[0])
	if idx < 0 {
		return "нет предмета в инвентаре - " + args[0], false
	}
	useful, ok := player.CurrentLocation.FindUseful(args[1])
	if !ok {
		return "не к чему применить", false
	}
	return (*useful).Use(item)
}

func takeAction(player *Player, args []string) (string, bool) {
	if len(args) != 1 {
		return "неизвестная команда", false
	}
	item, idx, container := player.CurrentLocation.FindItem(args[0])
	if idx < 0 {
		return "нет такого", false
	}
	if !(*item).Takeable() {
		return "невозможно", false
	}
	inventory, ok := player.Inventory()
	if !ok {
		return "некуда класть", false
	}
	(*container).RemoveItem(idx)
	inventory.AddItem(item)
	return "предмет добавлен в инвентарь: " + (*item).Name(), true
}

func wearAction(player *Player, args []string) (string, bool) {
	if len(args) != 1 {
		return "неизвестная команда", false
	}
	item, idx, container := player.CurrentLocation.FindItem(args[0])
	if idx < 0 {
		return "нет такого", false
	}
	if !(*item).Wearable() {
		return "невозможно", false
	}
	(*container).RemoveItem(idx)
	player.AddItem(item)
	return "вы надели: " + (*item).Name(), true
}

func inventoryAction(player *Player, args []string) (string, bool) {
	if len(args) > 0 {
		return "неизвестная команда", false
	}
	return player.ShowItems(), true
}
