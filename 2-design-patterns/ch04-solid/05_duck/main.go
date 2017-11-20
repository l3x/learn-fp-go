package main

import (
	"fmt"
	"errors"
	"log"
)
const DASHES = "----------------------"

type Pond struct {
	BugSupply       int
	StrokesRequired int
}

type StrokeBehavior interface {
	PaddleFoot(strokeSupply *int)
}

type EatBehavior interface {
	EatBug(strokeSupply *int)
}

type SurvivalBehaviors interface {
	StrokeBehavior
	EatBehavior
}


type Duck struct{}

func (Duck) Stroke(s StrokeBehavior, strokeSupply *int, p Pond) (err error) {
	for i := 0;  i < p.StrokesRequired; i++ {
		if *strokeSupply < p.StrokesRequired - i {
			err = errors.New("Our duck died!")
		}
		s.PaddleFoot(strokeSupply)
	}
	return err
}

func (Duck) Eat(e EatBehavior, strokeSupply *int, p Pond) {
	for i := 0;  i < p.BugSupply; i++ {
		e.EatBug(strokeSupply)
	}
}

type Foot struct{}
func (Foot) PaddleFoot(strokeSupply *int) {
	fmt.Println("- Foot, paddle!")
	*strokeSupply--
}

type Bill struct{}
func (Bill) EatBug(strokeSupply *int) {
	*strokeSupply++
	fmt.Println("- Bill, eat a bug!")
}

func (d Duck) SwimAndEat(se SurvivalBehaviors, strokeSupply *int, ponds []Pond) {
	for i := range ponds {
		pond := &ponds[i]
		err := d.Stroke(se, strokeSupply, *pond)
		if err != nil {
			log.Fatal(err)  // the duck died!
		}
		d.Eat(se, strokeSupply, *pond)
	}
}

type Capabilities struct {
	StrokeBehavior
	EatBehavior
	strokes int
}

func displayDuckStats(c Capabilities, ponds []Pond) {
	fmt.Printf("%s\n", DASHES)
	fmt.Printf("Ponds Processed:")
	for _, pond := range ponds {
		fmt.Printf("\n\t%+v", pond)
	}
	fmt.Printf("\nStrokes remaining: %+v\n", c.strokes)
	fmt.Printf("%s\n\n", DASHES)
}

func main() {
	var duck Duck
	capabilities := Capabilities{
		StrokeBehavior: Foot{},
		EatBehavior:    Bill{},
		strokes:        5,
	}

	ponds := []Pond{
		{BugSupply: 1, StrokesRequired: 3},
		{BugSupply: 1, StrokesRequired: 2},
	}
	duck.SwimAndEat(&capabilities, &capabilities.strokes, ponds)
	displayDuckStats(capabilities, ponds)

	ponds = []Pond{
		{BugSupply: 2, StrokesRequired: 3},
	}
	duck.SwimAndEat(&capabilities, &capabilities.strokes, ponds)
	displayDuckStats(capabilities, ponds)
}


