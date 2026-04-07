func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 { return 0 }

	type car struct {
		pos  int
		time float64
	}

	type node struct {
		val  *car
		prev *node
	}

	cars := make([]*car, target + 1)
	for i := 0; i < n; i++ {
		pos := position[i]
		cars[pos] = &car {
			pos: pos,
			time: float64(target - position[i]) / float64(speed[i]),
		}
	}

	var stack *node
	for i := target; i >= 0; i-- {
		if cars[i] == nil {
			continue
		}
		car := cars[i]
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
