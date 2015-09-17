package math1

import "testing"


type checkvaluesforfib struct{
	input uint
	output uint
}

var values = []checkvaluesforfib{
	{0,0},
	{1,1},
	{6,8},
}

func TestFib(t *testing.T){
	for _,x:=range values{
		y :=Fib(x.input)
		if(y!=x.output){
			t.Error("For", x.input, " expected " , x.output, " got" , y)		
		}	
	}
	
	
}
//func TestFib(t *testing.T){
//	f := Fib(6)
//	if (f!=8){
//	t.Error("expected 8, got",f)
//}
//}



func (r Rectangle) TestPerimeter(t *testing.T){
	rt := Rectangle{0,0,3,4}
	r_per:=rt.Perimeter()
	if(r_per!=14){
t.Error("Expected 14, got",r_per)}
} 

func (c Circle) TestPerimeter(t *testing.T){
	ct := Circle{0,0,1}
	c_per:=ct.Perimeter()
	if(c_per!=6.28){
t.Error("Expected 6.28, got",c_per)}
} 
