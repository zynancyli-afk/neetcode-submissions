type Point struct {
	x int
	y int
}

type CountSquares struct {
    points map[Point]int
}

func Constructor() CountSquares {
    return CountSquares {
		points: make(map[Point]int),
	}
}

func (this *CountSquares) Add(point []int)  {
	p := Point{point[0],point[1]}
    val,_:=this.points[p] 
	this.points[p] = val+1
}

func (this *CountSquares) Count(point []int) int {
	total := 0
	px, py := point[0], point[1]
	for k, v := range this.points {
		dx, dy := px-k.x, py-k.y
		if dx == 0 || abs(dx) != abs(dy) {
			continue
		}
		total += v * this.points[Point{px, k.y}] * this.points[Point{k.x, py}]
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
