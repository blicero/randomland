// /home/krylon/go/src/github.com/blicero/randomland/model/model.go
// -*- mode: go; coding: utf-8; -*-
// Created on 24. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-24 19:22:33 krylon>

// Package model provides the data types to model the game and its world.
package model

import "math/rand/v2"

// Range represents a range of numbers
type Range struct {
	Lo, Hi int
}

// Rand returns a random number from the given range.
func (r *Range) Rand() int {
	return rand.IntN(r.Hi-r.Lo) - r.Lo
} // func (r *Range) Rand() int

// Item is an object the player can interact with and or take.
type Item struct {
	ID          int64
	Name        string
	Description string
	Weight      int
	Portable    bool
	Properties  map[string]string
}

// Lifeform represents a monster or NPC the player can interact with.
type Lifeform struct {
	ID          int64
	Name        string
	Description string
	HP          int
	MaxHP       int
	XP          int
	Inventory   map[string]*Item
	Attack      int
	Evade       int
	Armor       int
	Damage      Range
	Initiative  Range
}
