package workflow

import (
	. "utils"
)

func ProcessCar(lineBase64 string) (err error, carJson string) {
	step := Get(lineBase64)
	step = Next(step, Base64ToBytes)
	step = Next(step, BytesToData)
	step = Next(step, TimestampData)
	step = Next(step, DataToJson)
	json, err := step(nil)
	if err != nil {
		Error.Error(err)
	} else {
		carJson = json.(string)
	}
	return
}
