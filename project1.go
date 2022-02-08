package main

import (
	"fmt"
)
var gVar int = 200
func main(){
	
	basicArray(5,5,5,5);
	variadicFunction(5,5,5,5);
	
}
func basicArray(x...int){
	//x:=[5]int{1,2,3,4,5} การประกาศตัวแปร Array
	var total int
	for _ ,value:=range x{
		total +=value
	}
	fmt.Println(x);
	fmt.Println(total)
}
func recursiveFunction(num int)int{
	if num == 0{
		return 1
	}
	return num*recursiveFunction(num-1)
}//Recursive Function

func variadicFunction(args...int){
	var total int 
	for _ ,n:= range args{
		total += n
	}
	fmt.Println(args);
	fmt.Println(total)
}//Variadic Function

func textFunction(str string){
	fmt.Println(str);	
}//Input Value

func intFunction(a int ,b int){
	fmt.Println(a+b);
}//Input Value

func returnFunction(a int ,b int)int{
	output:= a+b
	return output
}//Return Value

func outsideMain(){
	num:=100;
	fmt.Println(gVar+num);
}//Local Variable
 
func variableValue(){
	var value int=100
	num:=200;
	hello:="Hello"
	text:="Poom";
	fulltext:=hello + " " + text;
	valueBoolean:= true;
	num1,num2:=20,30;
	fmt.Println(value);
	fmt.Println(num);
	fmt.Println(text);
	fmt.Println(valueBoolean);
	fmt.Println(num1,num2);
	fmt.Println(text[0:4]);
	fmt.Println(fulltext);
	outsideMain();
}//Basic Go

func ifelseCondition(){
	fmt.Print("input number: ");
	var input float64
	fmt.Scanf("%f",&input)
	if input >1 && input <10 {
		fmt.Print(input)
		fmt.Print(" :is between 1-10")
	}else {
		fmt.Print(input)
		fmt.Print(" :is not between 1-10")
	}
}//if else condition

func switchcase() {
	fmt.Print("input number: ");
	var input int
	fmt.Scanf("%d",&input)
	switch input{
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	default: fmt.Println("Not in case")
	}
}//Switch Case

func forLoop(){
	for i:=0;i<=10;i++{
		if i%2 == 0{
			fmt.Println(i,"even");	
		}else{
			fmt.Println(i,"odd");	
		}
	}
}//For Loop