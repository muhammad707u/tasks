package task

import "math"

func RingS(r1 float64,r2 float64) (float64){
	var res float64
	p:=3.14
	r1=p*r1*r1
	r2=p*r2*r2
	res=math.Abs(r2-r1)
	return res
}