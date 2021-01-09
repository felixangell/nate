package gui

import (
	"github.com/felixangell/strife"
)

type Component interface {
	SetPosition(x, y int)
	Translate(x, y int)
	Resize(w, h int)

	GetPos() (int, int)
	GetSize() (int, int)

	OnInit()
	OnUpdate() bool
	OnRender(*strife.Renderer)
	OnDispose()

	NumComponents() int
	AddComponent(c Component)
	GetComponents() []Component

	HandleEvent(evt strife.StrifeEvent)

	SetFocus(focus bool)
	HasFocus() bool
}

type BaseComponent struct {
	x, y    int
	w, h    int
	focused bool
}

func (b *BaseComponent) GetPos() (int, int) {
	return b.x, b.y
}

func (b *BaseComponent) GetSize() (int, int) {
	return b.w, b.h
}

func (b *BaseComponent) SetFocus(focus bool) {
	b.focused = focus
}

func (b *BaseComponent) HasFocus() bool {
	return b.focused
}

func (b *BaseComponent) HandleEvent(evt strife.StrifeEvent) {
	// NOP
}

func (b *BaseComponent) SetPosition(x, y int) {
	b.x = x
	b.y = y
}

func (b *BaseComponent) Translate(x, y int) {
	b.x += x
	b.y += y
}

func (b *BaseComponent) Resize(w, h int) {
	b.w = w
	b.h = h
}
