package world

import (
	"fmt"
	"strings"
)

type ItemContainer interface {
	Name() string
	Items() []*Item
	AddItem(*Item)
	FindItem(string) (*Item, int, *ItemContainer)
	ShowItems() string
	RemoveItem(int)
	showTmpl() string
	emptyMsg() string
}

type ItemContainerImpl struct {
	_name     string
	_items    []*Item
	_showTmpl string
	_emptyMsg string
}

func (c *ItemContainerImpl) Name() string {
	return c._name
}
func (c *ItemContainerImpl) Items() []*Item {
	return c._items
}
func (c *ItemContainerImpl) showTmpl() string {
	return c._showTmpl
}
func (c *ItemContainerImpl) emptyMsg() string {
	return c._emptyMsg
}
func (c *ItemContainerImpl) AddItem(item *Item) {
	c._items = append(c._items, item)
}
func (c *ItemContainerImpl) FindItem(findName string) (*Item, int, *ItemContainer) {
	for i, item := range c.Items() {
		name := (*item).Name()
		if name == findName {
			var container ItemContainer = c
			return item, i, &container
		}
		container, ok := (*item).(ItemContainer)
		if ok && len(container.Items()) > 0 {
			found, idx, subContainer := container.FindItem(findName)
			if idx >= 0 {
				return found, idx, subContainer
			}
		}
	}
	return nil, -1, nil
}
func (c *ItemContainerImpl) RemoveItem(idx int) {
	c._items = append(c._items[:idx], c._items[idx+1:]...)
}
func (c *ItemContainerImpl) ShowItems() string {
	var itemNames []string
	for _, item := range c.Items() {
		if !(*item).Showable() {
			continue
		}
		name := (*item).Name()
		container, ok := (*item).(ItemContainer)
		if ok && len(container.Items()) > 0 {
			tmpl := container.showTmpl()
			if len(tmpl) == 0 {
				tmpl = container.Name() + `: %s`
			}
			itemNames = append(itemNames, fmt.Sprintf(tmpl, container.ShowItems()))
		} else {
			itemNames = append(itemNames, name)
		}
	}
	if len(itemNames) == 0 {
		emptyMsg := c.emptyMsg()
		if len(emptyMsg) == 0 {
			emptyMsg = "ничего интересного"
		}
		return emptyMsg
	}
	return strings.Join(itemNames, `, `)
}
