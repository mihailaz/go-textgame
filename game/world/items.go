package world

type Item interface {
	Name() string
	Wearable() bool
	Takeable() bool
	Showable() bool
}

type Useful interface {
	Name() string
	Use(*Item) (string, bool)
}

type Furniture struct {
	ItemContainerImpl
}

func (f *Furniture) Wearable() bool {
	return false
}

func (f *Furniture) Takeable() bool {
	return false
}

func (f *Furniture) Showable() bool {
	return len(f.Items()) > 0
}

type Backpack struct {
	ItemContainer
}

func (b *Backpack) Name() string {
	return "рюкзак"
}

func (b *Backpack) Wearable() bool {
	return true
}

func (b *Backpack) Takeable() bool {
	return false
}

func (b *Backpack) Showable() bool {
	return true
}

type Key struct {
	secret int
}

func (k *Key) Name() string {
	return "ключи"
}

func (k *Key) Wearable() bool {
	return false
}

func (k *Key) Takeable() bool {
	return true
}

func (k *Key) Showable() bool {
	return true
}

type DummyItem struct {
	_name     string
	_wearable bool
	_takeable bool
	_showable bool
}

func (i *DummyItem) Name() string {
	return i._name
}

func (i *DummyItem) Wearable() bool {
	return i._wearable
}

func (i *DummyItem) Takeable() bool {
	return i._takeable
}

func (i *DummyItem) Showable() bool {
	return i._showable
}
