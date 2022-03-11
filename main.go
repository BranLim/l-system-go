package main

import (
	"Coding_Night/painter"
	"time"

	"github.com/golang-collections/collections/stack"
)

func produceTree(channel chan string) {
	rules := BuildRules("1", "11", "0", "1[0]0")
	axiom := '0'
	//should be 7 but the result is too big
	for i := 0; i <= 6; i++ {
		channel <- rules.Produce(string(axiom), i)
	}
	close(channel)
}

func main() {
	paint := painter.New()
	var width, height = paint.ScreenSize()
	var size = 0
	if width < height {
		size = width
	} else {
		size = height - 3
	}
	var mid = size / 2
	paint.Clear()
	paint.StartDrawing(size, 2)
	messages := make(chan string)
	go produceTree(messages)

	for actual := range messages {
		var stack = stack.New()

		paint.StartDrawing(size, 2)
		drawString(paint, actual, mid, size-2, stack)
		paint.EndDrawing()
		time.Sleep(1 * time.Second)
		paint.Clear()
	}

}

func drawString(paint *painter.Painter, actual string, initial_x int, initial_y int, stack *stack.Stack) {
	var point *painter.Point
	paint.DrawText(1, 1, actual)
	for k, symbol := range actual {
		if k == 0 {
			point = painter.BuildPoint(symbol, initial_x, initial_y, painter.N)
		}
		if symbol == '[' {
			stack.Push(point)
		} else if symbol == ']' {
			point = stack.Pop().(*painter.Point)
		} else if symbol == '0' {
			point = point.Next(symbol)
			paint.DrawPoint(point.AsRune(), point.X, point.Y)
		} else {
			paint.DrawPoint(point.AsRune(), point.X, point.Y)
		}
		point = point.Next(symbol)
	}
	paint.DrawPoint(point.AsRune(), point.X, point.Y)
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
