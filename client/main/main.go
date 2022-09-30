package main

import (
	"fmt"
	"os"
	"go_code/ChatRoom/client/processes"
)


//gernerate main page
func main(){
	var key int
	var loop bool = true
	for {

		fmt.Println("---------------------Welcome to ChatRoom---------------------")
		fmt.Println("\t\t\t 1、LOGIN IN")
		fmt.Println("\t\t\t 2、REGISTE")
		fmt.Println("\t\t\t 3、EXIT")
		fmt.Println("-------------------------------------------------------------")
		fmt.Print("Please choose one:")
		fmt.Scanf("%d\n",&key)
		var UserId int
		var UserName string
		var Password string
		//create a UserProcess case
		up := &processes.UserProcess{
		}
		switch key {
		case 1:
			fmt.Println("welcome to LOGIN")
			fmt.Print("Login UserId:")
			fmt.Scanln(&UserId)
			fmt.Print("Login Password:")
			fmt.Scanln(&Password)
			err := up.Login(UserId,Password)
			if err != nil {
				fmt.Println("login fail")
			}
			loop = false
		case 2:
			fmt.Println("welcome to REGISTE")
			fmt.Print("Registe UserId:")
			fmt.Scanln(&UserId)
			fmt.Print("Registe UserName:")
			fmt.Scanln(&UserName)
			fmt.Print("Registe Password:")
			fmt.Scanln(&Password)
			err := up.Register(UserId,UserName,Password)
			if err != nil {
				fmt.Println("Register fail")
			}
			loop = false
		case 3:
			fmt.Println("you have exit the ChatRoom!")
			os.Exit(0)
		default:
			fmt.Printf("there is no mode %d,try again!\n",key)
		}
		if !loop {
			break
		} 
	}

}