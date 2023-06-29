package player

import (
	"fmt"
	"game/world"
	"strings"
)

type Player struct {
	world.ItemContainerImpl
	CurrentLocation *world.Location
	Actions         map[string]Action
	TaskList        TaskList
}

func (p *Player) ShowInspection() string {
	tmpl := p.CurrentLocation.Inspection
	if len(tmpl) == 0 {
		tmpl = `%s. %s`
	}
	items := p.CurrentLocation.ShowItems()
	ways := p.CurrentLocation.ShowWays()
	if strings.Count(tmpl, `%s`) == 3 {
		tasks := p.TaskList.ShowTasks()
		return fmt.Sprintf(tmpl, items, tasks, ways)
	}
	return fmt.Sprintf(tmpl, items, ways)
}

func (p *Player) Do(command string) string {
	var actionAndArgs = strings.Split(command, ` `)
	action, ok := p.Actions[actionAndArgs[0]]
	if !ok {
		return "неизвестная команда"
	}
	answer, ok := action(p, actionAndArgs[1:])
	if ok {
		p.TaskList.Notify(p)
	}
	return answer
}

func (p *Player) Inventory() (world.ItemContainer, bool) {
	for _, item := range p.Items() {
		container, ok := (*item).(world.ItemContainer)
		if ok {
			return container, true
		}
	}
	return nil, false
}

func (p *Player) Move(location *world.Location) string {
	p.CurrentLocation = location
	return p.CurrentLocation.ShowWelcome()
}
