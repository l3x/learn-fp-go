package main

import "fmt"

var languages map[string]string

func init(){
	languages= make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["go"] = "Golang"
}
func Get(key string) (string){
	return languages[key]
}
func Add(key,value string){
	languages[key]=value
}
func GetAll() (map[string]string){
	return languages
}

func main() {
	fmt.Printf("languages: %v\n", languages)
	fmt.Printf("Get('Ruby'): %v\n", Get("rb"))
	fmt.Println("Add('Perl')")
	Add("pl", "Perl")
	fmt.Printf("GetAll(): %v\n", GetAll())
}
