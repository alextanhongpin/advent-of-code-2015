// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var dp map[dpkey]int

type dpkey struct {
	dmg, arm, cost int
}

func init() {
	dp = make(map[dpkey]int)
}

func main() {
	boss := player{
		name: "boss",
		hp:   104,
		dmg:  8,
		arm:  1,
	}
	you := player{name: "you", hp: 100}
	items := parse(input)
	fmt.Println("part 1:", cheapestCostToWin(you, boss, items, 3, len(items[3])-1))

	dp = make(map[dpkey]int)
	fmt.Println("part 2:", mostExpensiveButLost(you, boss, items, 3, len(items[3])-1))
}

func mostExpensiveButLost(you, enemy player, items map[int][]item, inv int, inv0 int) int {
	if inv < 1 || inv0 < 0 {
		if len(you.weapons) < 1 {
			return -1
		}
		key := dpkey{dmg: you.damage(), arm: you.armor(), cost: you.cost()}
		if v, ok := dp[key]; ok {
			return v
		}
		cost := -1
		if len(you.rings) > 2 {
			cost = -1
		} else {
			winner, _ := simulate(you, enemy)
			cost = you.cost()
			if winner.name == "you" {
				cost = -1
			}
		}
		dp[key] = cost
		return cost
	}

	eqp := you
	item := items[inv][inv0]

	switch inv {
	case 1:
		eqp.weapons = append(eqp.weapons, item)
		return max(
			mostExpensiveButLost(eqp, enemy, items, inv-1, 0),    // Use this weapon.
			mostExpensiveButLost(you, enemy, items, inv, inv0-1), // Choose another weapon.
		)
	case 2:
		eqp.armors = append(eqp.armors, item)
		return max(
			mostExpensiveButLost(eqp, enemy, items, inv-1, len(items[inv-1])-1), // Equip an armor, now go to weapons.
			mostExpensiveButLost(you, enemy, items, inv, inv0-1),                // Pick the next armor.
			mostExpensiveButLost(you, enemy, items, inv-1, len(items[inv-1])-1), // Don't equip any armor.
		)
	case 3:
		eqp.rings = append(eqp.rings, item)
		return max(
			mostExpensiveButLost(eqp, enemy, items, inv-1, len(items[inv-1])-1), // Equip this ring, and select next items.
			mostExpensiveButLost(eqp, enemy, items, inv, inv0-1),                // Already equipped, equip another ring.
			mostExpensiveButLost(you, enemy, items, inv-1, len(items[inv-1])-1), // Don't equip any ring
			mostExpensiveButLost(you, enemy, items, inv, inv0-1),                // Pick another ring.
		)
	default:
		return -1
	}
}

func cheapestCostToWin(you, enemy player, items map[int][]item, inv int, inv0 int) int {
	if inv < 1 || inv0 < 0 {
		key := dpkey{dmg: you.damage(), arm: you.armor(), cost: you.cost()}
		if v, ok := dp[key]; ok {
			return v
		}
		cost := math.MaxInt
		if len(you.rings) > 2 {
			cost = math.MaxInt
		} else {
			winner, _ := simulate(you, enemy)
			cost = winner.cost()
			if winner.name != "you" {
				cost = math.MaxInt
			}
		}
		dp[key] = cost
		return cost
	}

	eqp := you
	item := items[inv][inv0]

	switch inv {
	case 1:
		eqp.weapons = append(eqp.weapons, item)
		return min(
			cheapestCostToWin(eqp, enemy, items, inv-1, -1),   // Use this weapon.
			cheapestCostToWin(you, enemy, items, inv, inv0-1), // Choose another weapon.
		)
	case 2:
		eqp.armors = append(eqp.armors, item)
		return min(
			cheapestCostToWin(eqp, enemy, items, inv-1, len(items[inv-1])-1), // Equip an armor, now go to rings.
			cheapestCostToWin(you, enemy, items, inv, inv0-1),                // Pick the next armor.
			cheapestCostToWin(you, enemy, items, inv-1, len(items[inv-1])-1), // Don't equip any armor.
		)
	case 3:
		eqp.rings = append(eqp.rings, item)
		return min(
			cheapestCostToWin(eqp, enemy, items, inv-1, len(items[inv-1])-1), // Equip this ring, and select next items.
			cheapestCostToWin(eqp, enemy, items, inv, inv0-1),                // Already equipped, equip another ring.
			cheapestCostToWin(you, enemy, items, inv-1, len(items[inv-1])-1), // Don't equip any ring
			cheapestCostToWin(you, enemy, items, inv, inv0-1),                // Pick another ring.
		)
	default:
		return math.MaxInt
	}
}

func max(a int, nums ...int) int {
	for _, n := range nums {
		if n > a {
			a = n
		}
	}
	return a
}

func min(a int, nums ...int) int {
	for _, n := range nums {
		if n < a {
			a = n
		}
	}
	return a
}

type player struct {
	name         string
	hp, dmg, arm int
	weapons      []item
	armors       []item
	rings        []item
}

func (p player) damageInflicted(other player) int {
	return max(p.damage()-other.armor(), 1)
}

func (p player) damage() int {
	dmg := p.dmg
	for _, w := range p.weapons {
		dmg += w.damage
	}
	for _, r := range p.rings {
		dmg += r.damage
	}
	return dmg
}

func (p player) armor() int {
	arm := p.arm
	for _, w := range p.armors {
		arm += w.armor
	}
	for _, r := range p.rings {
		arm += r.armor
	}
	return arm
}

func (p player) cost() int {
	total := 0
	for _, w := range p.weapons {
		total += w.cost
	}
	for _, a := range p.armors {
		total += a.cost
	}
	for _, r := range p.rings {
		total += r.cost
	}
	return total
}

func simulate(p1, p2 player) (player, int) {
	rounds := 0

	for {
		if rounds%2 == 0 {
			p2.hp -= p1.damageInflicted(p2)
		} else {
			p1.hp -= p2.damageInflicted(p1)
		}
		rounds++
		if p1.hp <= 0 {
			return p2, rounds
		}
		if p2.hp <= 0 {
			return p1, rounds
		}
	}
}

type item struct {
	name                string
	cost, damage, armor int
}

func parse(input string) map[int][]item {
	lines := strings.Split(input, "\n")
	items := make(map[int][]item)
	typ := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "Weapons") || strings.HasPrefix(line, "Armor") || strings.HasPrefix(line, "Rings") {
			typ++
			continue
		}
		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		first := strings.Join(parts[:len(parts)-3], " ")
		second := parts[len(parts)-3]
		third := parts[len(parts)-2]
		fourth := parts[len(parts)-1]

		items[typ] = append(items[typ], item{
			name:   first,
			cost:   toInt(second),
			damage: toInt(third),
			armor:  toInt(fourth),
		})
	}
	return items
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var input = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`
