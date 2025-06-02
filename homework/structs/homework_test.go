package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = name
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = int32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mana = int32(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.health = int32(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respect = int32(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strength = int32(strength)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experience = int32(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.level = int32(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= 1 << 0
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= 1 << 1
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= 1 << 2
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags = (person.flags &^ (3 << 3)) | (byte(personType) << 3)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	name       string
	x          int32
	y          int32
	z          int32
	gold       int32
	mana       int32
	health     int32
	respect    int32
	strength   int32
	experience int32
	level      int32
	flags      byte
}

func NewGamePerson(options ...Option) GamePerson {
	p := GamePerson{}
	for _, option := range options {
		option(&p)
	}
	return p
}

func (p *GamePerson) Name() string {
	return p.name
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.mana)
}

func (p *GamePerson) Health() int {
	return int(p.health)
}

func (p *GamePerson) Respect() int {
	return int(p.respect)
}

func (p *GamePerson) Strength() int {
	return int(p.strength)
}

func (p *GamePerson) Experience() int {
	return int(p.experience)
}

func (p *GamePerson) Level() int {
	return int(p.level)
}

func (p *GamePerson) HasHouse() bool {
	return (p.flags & (1 << 0)) != 0
}

func (p *GamePerson) HasGun() bool {
	return (p.flags & (1 << 1)) != 0
}

func (p *GamePerson) HasFamilty() bool {
	return (p.flags & (1 << 2)) != 0
}

func (p *GamePerson) Type() int {
	return int((p.flags >> 3) & 3)
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
