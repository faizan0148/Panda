package greetings 

//GreetingsString is used for doing Meh things
var GreetingsString = "Hello Meh World from Wajeeh"

//PrintGreetings does just about anything you want it to do. 
func PrintGreetings (name string) string {	
	return GreetingsString + "-" + name; 
}