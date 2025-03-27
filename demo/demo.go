// /home/krylon/go/src/github.com/blicero/randomland/demo/demo.go
// -*- mode: go; coding: utf-8; -*-
// Created on 27. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-27 18:37:59 krylon>

// Package demo provides a simple World for testing the game engine and mechanics.
package demo

import "github.com/blicero/randomland/model"

var idcnt int64

func newid() int64 {
	idcnt++
	return idcnt
} // func newid() int64

func item(name, desc string, weight int, portable bool, prop map[string]string) *model.Item {
	return &model.Item{
		ID:          newid(),
		Name:        name,
		Description: desc,
		Weight:      weight,
		Portable:    portable,
		Properties:  prop,
	}
} // func item(name, desc string, weight int, portable bool, prop map[string]string) *model.Item

func lifeform(name, desc string, hp, hpmax, xp, attack, evade, armor int, dmg, ini model.Range, inv map[int64]*model.Item) *model.Lifeform {
	return &model.Lifeform{
		ID:          newid(),
		Name:        name,
		Description: desc,
		HP:          hp,
		MaxHP:       hpmax,
		XP:          xp,
		Attack:      attack,
		Evade:       evade,
		Armor:       armor,
		Damage:      dmg,
		Initiative:  ini,
		Inventory:   inv,
	}
} // func lifeform(name, desc string, hp, hpmax, xp, attack, evade, armor int, dmg, ini model.Range, inv map[int64]*model.Item) *model.Lifeform
