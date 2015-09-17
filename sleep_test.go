package math2

import "testing"
import "time"

func TestSleeptest(t *testing.T){

	
	var c chan string = make(chan string)
	var res string
	go Sleeptest(2,c)
        select{
	case d:=<-c:
		res=d
	case <-time.After(time.Second *3):
		res="timeout"
}
	if(res!="2"){
		t.Error("expected 2, got timeout")}
	
}
