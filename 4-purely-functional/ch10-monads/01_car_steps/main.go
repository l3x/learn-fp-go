package main

import (
	"workflow"
	. "utils"
	"bufio"
	"os"
)

func init() {
	GetOptions()
	InitLog("trace.log", os.Stdout, os.Stdout, os.Stderr)
	Info.Println("AppEnv:", Config.AppEnv)
}

func main() {
	carCntr := 0
	if file, err := os.Open(Config.DataFilepath); err == nil {
	    defer file.Close()
		Info.Println("----")
	    scanner := bufio.NewScanner(file)
	    for scanner.Scan() {
	        carCntr += 1
			Info.Println("Processing car #", carCntr)
	        line :=  scanner.Text()
	        Info.Println("IN :", line)
	        err, carJson := workflow.ProcessCar(line)
	        if err == nil {
				Info.Println("OUT:", carJson)
			}
			Info.Println("----")
	    }
	    if err = scanner.Err(); err != nil {
			Error.Error(err)
	    }
	} else {
		Error.Error(err)
	}
}

