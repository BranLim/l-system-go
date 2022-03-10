package main

import (
	"Coding_Night/painter"
)

func main() {
	paint := painter.New()
	var width, height = paint.ScreenSize()
	var size = 0
	if width < height {
		size = width
	} else {
		size = height
	}
	var mid = size / 2
	paint.Clear()
	paint.StartDrawing(size, 2)
	var rules = BuildRules("A", "AB", "B", "A")
	paint.DrawAlive('A', mid, 0)
	for i := 1; i < 8; i++ {
		var actual = rules.Produce("A", i)
		var chars = len(actual)
		for k, symbol := range actual {
			var x = mid - (chars / 2) + k
			paint.DrawAlive(symbol, x, i)
		}
	}
	paint.EndDrawing()
	paint.DrawAlive(' ', 1, height)
}

type Rules struct {
	rules map[string]string
}

func BuildRules(tokens ...string) *Rules {
	var first = ""
	var second = ""
	var ruleDictionary = map[string]string{}
	for i, token := range tokens {
		if i%2 == 0 {
			first = token
		} else {
			second = token
			ruleDictionary[first] = second
		}
	}
	return &Rules{
		rules: ruleDictionary,
	}
}

func (rules *Rules) produce(axiom string) string {
	var result = ""
	for _, character := range axiom {
		var charAsStr = string(character)
		production, ok := rules.rules[charAsStr]
		if ok {
			result += production
		} else {
			result += charAsStr
		}
	}
	return result
}

func (rules *Rules) Produce(axiom string, level int) string {
	var result = axiom
	for i := 0; i < level; i++ {
		result = rules.produce(result)
	}
	return result
}
