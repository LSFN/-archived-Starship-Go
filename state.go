package main

type State struct {
	Reactor
	Systems
	Engines
	Thrusters
	Sensors
}

type Reactor struct {
	Reactant    float64
	Coolant     float64
	PowerOutput float64
	Temperature float64
}

type Systems []*System

type System struct {
	ID         uint32
	Name       string
	State      bool
	PowerLevel float64
}

type Engines struct {
	Left  float64
	Right float64
}

type Thrusters struct {
	ForwardLeft  float64
	ForwardRight float64
	RearLeft     float64
	RearRight    float64
}

type Sensors []*SpaceObject

type SpaceObject struct {
	Type     SpaceObjectType
	Position Point
}

type SpaceObjectType uint8

const (
	SHIP SpaceObjectType = iota
	ASTEROID
)

type Point struct {
	X, Y float64
}
