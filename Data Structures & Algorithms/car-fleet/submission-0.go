import (
	"slices"
	"cmp"
)
	

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 { return 0 }

	type car struct {
		pos  int
		time float64
	}

	type node struct {
		val  car
		prev *node
	}

	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car {
			pos: position[i],
			time: float64(target - position[i]) / float64(speed[i]),
		}
	}

	slices.SortFunc(cars, func(a, b car) int {
		return cmp.Compare(b.pos, a.pos) // Descending order
	})

	var stack *node
	for _, car := range cars {
		if stack == nil || car.time > stack.val.time {
			stack = &node{
				val: car,
				prev: stack,
			}
		}
	}
	
	fleets := 0
	for stack != nil {
		stack = stack.prev
		fleets++
	}
	return fleets
}
