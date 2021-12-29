// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

var spells = []string{
	"magic_missile",
	"drain",
	"shield",
	"poison",
	"recharge",
}

var manaCosts = map[string]int{
	"magic_missile": 53,
	"drain":         73,
	"shield":        113,
	"poison":        173,
	"recharge":      229,
}

var dp = map[string]int{}

func main() {
	you := player{
		name:  "you",
		hp:    50,
		mana:  500,
		timer: make(map[string]int),
	}
	boss := player{
		name:   "boss",
		hp:     58,
		damage: 9,
		timer:  make(map[string]int),
	}
	fmt.Println(minMana(you, boss, "easy"))
	dp = make(map[string]int)
	fmt.Println(minMana(you, boss, "hard"))
}

func debug(you, boss player) {
	if you.name == "you" {
		fmt.Printf("-- Player Turn --\n")
	} else {
		you, boss = boss, you
		fmt.Printf("-- Boss Turn --\n")
	}
	fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", you.hp, you.armor, you.mana)
	fmt.Printf("- Boss has %d hit points\n", boss.hp)
	fmt.Println(you.manaPool, you.timer)
	fmt.Println()
}

func minMana(p1, p2 player, difficulty string) int {
	hard := difficulty == "hard"
	if hard {
		if p1.name == "you" {
			p1.hp--
		} else {
			p2.hp--
		}
	}
	p1.underEffect(&p2)
	p2.underEffect(&p1)
	//debug(p1, p2)

	if p1.dead() {
		switch p1.name {
		case "you":
			return math.MaxInt
		default:
			return p2.manaPool
		}
	}
	key := ""
	if p1.name == "you" {
		key += fmt.Sprint(p1)
		key += ":"
		key += fmt.Sprint(p2)
	} else {
		key += fmt.Sprint(p2)
		key += ":"
		key += fmt.Sprint(p1)
	}
	if cost, ok := dp[key]; ok {
		return cost
	}

	switch p1.name {
	case "you":
		bestCost := math.MaxInt
		for _, spell := range spells {
			p2c := p2.clone()
			p1c := p1.clone()
			if p1c.magic(&p2c, spell) {
				cost := minMana(p2c, p1c, difficulty)
				if cost < bestCost {
					bestCost = cost
				}
			}
		}
		dp[key] = bestCost
		return bestCost
	default:
		p1.attack(&p2)
		return minMana(p2, p1, difficulty)
	}
}

type player struct {
	name          string
	armor, damage int
	hp, mana      int
	timer         map[string]int
	manaPool      int // Stores the mana used by the player
}

func (p player) clone() player {
	timer := make(map[string]int)
	for k, v := range p.timer {
		timer[k] = v
	}
	return player{
		name:     p.name,
		armor:    p.armor,
		damage:   p.damage,
		hp:       p.hp,
		mana:     p.mana,
		manaPool: p.manaPool,
		timer:    timer,
	}
}

func (p *player) dead() bool {
	return p.hp <= 0
}

func (p *player) attack(other *player) {
	other.hp -= max(p.damage-other.armor, 1)
}

func (p *player) magic(other *player, spell string) bool {
	dmg, ok := p.cast(spell)
	if !ok {
		return false
	}
	other.hp -= dmg
	return true
}

func (p *player) cast(spell string) (int, bool) {
	if v, ok := p.timer[spell]; ok && v > 0 {
		return 0, false
	}
	cost := manaCosts[spell]
	if p.mana < cost {
		return 0, false
	}

	p.manaPool += cost
	p.mana -= cost
	damage := 0
	switch spell {
	case "magic_missile":
		damage = 4
	case "drain":
		p.hp += 2
		damage = 2
	case "shield":
		p.timer[spell] = 5
		p.armor = 7
	case "poison":
		p.timer[spell] = 6
	case "recharge":
		p.timer[spell] = 5
	}
	return damage, true
}

func (p *player) underEffect(other *player) {
	dmg := other.effect()
	p.hp -= dmg
}

func (p *player) effect() int {
	damage := 0
	for spell, duration := range p.timer {
		remaining := duration - 1
		active := remaining > -1

		if active {
			switch spell {
			case "shield":
				p.armor = 7
			case "poison":
				damage += 3
			case "recharge":
				p.mana += 101
			}
		} else {
			if spell == "shield" {
				p.armor = 0
			}
		}
		if active {
			p.timer[spell] = remaining
		} else {
			delete(p.timer, spell)
		}
	}
	return damage
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
