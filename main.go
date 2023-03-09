package main

import("fmt")

type person struct{
	first string
}


func main(){

	p := person{
		first: "John Frusciante",
	}
	fmt.Println(p.first)
	ej2(&p)
	fmt.Printf(p.first)
}
func ej1(){
	x:=1
	fmt.Print(&x)
}

func ej2(p *person){
	p.first = "David Bowie"

}