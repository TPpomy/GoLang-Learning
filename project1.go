package main

import (
	"fmt"
	s "strings"
)
var gVar int = 200
var p =fmt.Println
func main(){
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ",s.Count("test","t"))
	p("HasPreFix: ",s.HasPrefix("test","te"))
	p("HasSuffix: ",s.HasSuffix("test","st"))
	p("Index: ",s.Index("test","e"))
	p("Join: ",s.Join([]string{"a","b"},"-"))
	p("Repeat: ",s.Repeat("a",5))
	p("Replace: ",s.Replace("foo","o","0",-1))
	p("Replace: ",s.Replace("fooo","o","0",2))
	p("Split: ",s.Split("a-b-c-d-e","-"))
	p("ToLower: ",s.ToLower("LOWER"))
	p("ToUPper: ",s.ToUpper("upper"))
	p("Len: ",len("Hello"))
	


}
func goRoutine(){
	f:=func(n int) {
		for i:=0; i<=5; i++{
			fmt.Println(n , " " , i)
			
		}
	}
	go f(0)
	var input string
	fmt.Scanln(&input);
}
func basicStructure(){
	type books struct{
		title string
		subtitle string
		author string
		price float64
	}
	book:=books{title:"GoLang Learning",subtitle:"in 2022",author:"Poom",price:231.21}
	fmt.Println(book);
	
}
func closurefunction(){
	app:=func (x,y int) int{
		return x+y
	}
	fmt.Println(app(3,4))
}
func mapFunction(){
	x:=map[string] string{
		"TH":"Thailand",
		"JP":"Japan",
	}
	fmt.Println(x["TH"]);
}
func sliceArray(){
	// slice1:=[]int{1,2,3}
	// slice2:=append(slice1,4,5)
	// fmt.Println(slice2);
	slice1:=[]int{1,2,3,4,5}
	slice2:=make([]int,3)
	copy(slice2,slice1)
	fmt.Println(slice2)
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
	fmt.Println(total/int(5))
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