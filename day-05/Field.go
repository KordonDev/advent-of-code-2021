package main

type Field struct {
	positions map[string]int
}

func (field *Field) increasePoint(point *Point) {
	counter, ok := field.positions[point.toString()]
	if ok {
		field.positions[point.toString()] = counter + 1
	} else {
		field.positions[point.toString()] = 1
	}
}

func (field *Field) highPoints() int {
	result := 0
	for _, counter := range field.positions {
		if counter >= 2 {
			result = result + 1
		}
	}
	return result
}

func NewField() *Field {
	return &Field{positions: make(map[string]int)}
}
