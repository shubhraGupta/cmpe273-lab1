package math1

import ("math")

func distance(x1,y1,x2,y2 float64) float64{
    x := x2 -x1
    y := y2 -y1
    result := math.Sqrt(x*x + y*y)
    return result
}

type shape interface{
    area() float64
    Perimeter() float64
}


type Rectangle struct{
    x1,y1,x2,y2 float64
}

type Circle struct{
    a,b,rad float64
}

func (c Circle) area() float64{
    return (math.Pi *c.rad *c.rad)
}    
func (r Rectangle) area() float64{
    l := distance(r.x1,r.y1,r.x2,r.y1)
    b := distance(r.x1,r.y1,r.x1,r.y2)
    return l*b
}


func (c Circle) Perimeter() float64{
    return math.Pi * c.rad * 2
}

func (r Rectangle) Perimeter() float64{
    l := distance(r.x1,r.y1,r.x2,r.y1)
    b := distance(r.x1,r.y1,r.x1,r.y2)
    return 2*(l+b)
}




