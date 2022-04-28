package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (point *Point) toString() string {
	return strconv.Itoa(point.x) + "," + strconv.Itoa(point.y)
}

func NewPoint(coordsString string) *Point {
	coords := strings.Split(coordsString, ",")
	return &Point{x: stringToInt(coords[0]), y: stringToInt(coords[1])}
}

type Move struct {
	from *Point
	to   *Point
}

func NewMove(move string) *Move {
	parts := strings.Split(move, " -> ")
	from := NewPoint(parts[0])
	to := NewPoint(parts[1])
	return &Move{from: from, to: to}
}

func (move *Move) touchedPoints() []*Point {
	var touchedPoints []*Point
	if move.from.x == move.to.x {
		for i := min(move.from.y, move.to.y); i <= max(move.from.y, move.to.y); i++ {
			touchedPoints = append(touchedPoints, &Point{x: move.from.x, y: i})
		}
		return touchedPoints
	}
	if move.from.y == move.to.y {
		for i := min(move.from.x, move.to.x); i <= max(move.from.x, move.to.x); i++ {
			touchedPoints = append(touchedPoints, &Point{x: i, y: move.to.y})
		}
		return touchedPoints
	}
	return touchedPoints
}

func (move *Move) allTouchedPoints() []*Point {
	touchedPoints := move.touchedPoints()
	if len(touchedPoints) > 0 {
		return touchedPoints
	}
	diff := abs(move.from.x - move.to.x)
	xSign := (move.to.x - move.from.x) / diff
	ySign := (move.to.y - move.from.y) / diff
	for i := 0; i <= diff; i++ {
		touchedPoints = append(touchedPoints, &Point{x: move.from.x + i*xSign, y: move.from.y + i*ySign})
	}
	return touchedPoints
}
