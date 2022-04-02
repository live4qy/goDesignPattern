package Factory

import "fmt"

//工厂模式

type Restaurant interface{
	GetFood()
}

type R1 struct{
}
func (r *R1) GetFood(){
	fmt.Println("now is R1.GetFood")
}

type R2 struct{
}
func (r *R2) GetFood(){
	fmt.Println("now is R2.GetFood")
}

func NewRestaurant(name string) Restaurant{
	/*
	根据标识(name)获取
	*/
	switch name{
	case "R1":return &R1{}
	case "R2":return &R2{}
	}
	return nil
}






























