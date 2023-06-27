package main

import "fmt"

func Run() error {
	fmt.Printf("App starting....")
	return nil

}

func main() {
	fmt.Println("Go Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
