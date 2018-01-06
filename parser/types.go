package parser

import (
	"fmt"
)

type WarehouseInfo struct {
	IsValid  int
	Location string
	Region   string
	District string
	Class    string
	Square   string
	Money    string
}

func (whi *WarehouseInfo) String() string {
	return fmt.Sprintf("location: %s, region: %s, district: %s, class: %s, square: %s, money: %s", whi.Location,
		whi.Region, whi.District, whi.Class, whi.Square, whi.Money)
}

func CreateWarehouseInfo() *WarehouseInfo {
	wi := &WarehouseInfo{
		IsValid:  0,
		Location: "",
		Region:   "",
		District: "",
		Class:    "",
		Square:   "",
		Money:    "",
	}
	return wi
}
