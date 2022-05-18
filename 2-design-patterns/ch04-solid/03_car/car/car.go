package car

import "fmt"

type Car struct {
	Make string
	Model string
}
func (c Car) Tires() int { return 4 }
func (c Car) PrintInfo() {
	fmt.Printf("%v has %d tires\n", c, c.Tires())
}

type CarWithSpare struct {
	Car
}
func (o CarWithSpare) Tires() int { return 5 }

func (c CarWithSpare) PrintInfo() {
	fmt.Printf("%v has %d tires\n", c, c.Tires())
}
//func (c CarWithSpare) PrintInfo(upCase bool) {
//	if upCase {
//		fmt.Printf("%v HAS %d TIRES\n", c, c.Tires())
//	} else {
//		fmt.Printf("%v has %d tires\n", c, c.Tires())
//	}
//}
