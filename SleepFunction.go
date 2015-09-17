package math2

import "time"

func Sleeptest(i int, c chan string){
	
	sleepfunc(i)	
	c <-"2"
}

func sleepfunc(i int){
	<-time.After(time.Second * time.Duration(i))
}
        

    
