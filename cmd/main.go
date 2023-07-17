package main

import "github.com/fishmanDK/good_traning/intenal/pkg/app"


func main(){
	a, err := app.NewApp()
	if err != nil{
		return 
	}

	err = a.Run()
	if err != nil{
		return
	}
}

