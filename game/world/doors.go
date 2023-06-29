package world

type Door struct {
	Locked bool
	secret int
	first  *Way
	second *Way
}

func (d *Door) Name() string {
	return "дверь"
}

func (d *Door) Use(item *Item) (string, bool) {
	key, ok := (*item).(*Key)
	if !ok {
		return "не к чему применить", false
	}
	locked, ok := d.UseKey(key)
	if !locked {
		return "дверь открыта", ok
	}
	return "дверь закрыта", ok
}

func (d *Door) OtherWay(from *Location) (*Way, bool) {
	if d.first.Location.Equals(from) {
		return d.second, true
	}
	if d.second.Location.Equals(from) {
		return d.first, true
	}
	return nil, false
}

func (d *Door) ComeIn(from *Location) (*Way, bool) {
	if d.Locked {
		return nil, false
	}
	return d.OtherWay(from)
}

func (d *Door) UseKey(key *Key) (bool, bool) {
	if key.secret == d.secret {
		d.Locked = !d.Locked
		return d.Locked, true
	}
	return d.Locked, false
}

type Way struct {
	Location *Location
	name     string
}

func (w *Way) GetName() string {
	if len(w.name) > 0 {
		return w.name
	}
	return w.Location.Name()
}
