package world

import (
	"fmt"
	"math/rand"
	"strings"
)

type Location struct {
	ItemContainerImpl
	Doors      []*Door
	Usefuls    []*Useful
	Inspection string
	Welcome    string
}

func (l *Location) AddDoor(other *Location) *Door {
	door := &Door{false, rand.Int(), &Way{Location: l}, &Way{Location: other}}
	l.Doors = append(l.Doors, door)
	other.Doors = append(other.Doors, door)
	return door
}

func (l *Location) AddUseful(useful *Useful) {
	l.Usefuls = append(l.Usefuls, useful)
}

func (l *Location) FindUseful(findName string) (*Useful, bool) {
	for _, useful := range l.Usefuls {
		name := (*useful).Name()
		if name == findName {
			return useful, true
		}
	}
	return nil, false
}

func (l *Location) ShowWelcome() string {
	tmpl := l.Welcome
	if len(tmpl) == 0 {
		return l.ShowWays()
	}
	return fmt.Sprintf(tmpl, l.ShowWays())
}

func (l *Location) ShowWays() string {
	var wayNames []string
	for _, door := range l.Doors {
		way, ok := door.OtherWay(l)
		if !ok {
			continue
		}
		wayNames = append(wayNames, way.GetName())
	}
	return fmt.Sprintf("можно пройти - %s", strings.Join(wayNames, `, `))
}

func (l *Location) FindDoorTo(destination string) (*Door, bool) {
	for _, door := range l.Doors {
		if door.first.GetName() == destination || door.second.GetName() == destination {
			return door, true
		}
	}
	return nil, false
}

func (l *Location) Equals(other *Location) bool {
	return l.Name() == other.Name()
}
