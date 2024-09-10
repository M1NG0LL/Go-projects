package main

import (
	"fmt"
)

type mobile interface {
	ChargeAppleMobile()
}

type Apple struct {}

func (a *Apple) ChargeAppleMobile() {
	fmt.Println("Charging Apple Phone.")
}

type Android struct {}

func (a *Android) chargeAndroidMobile() {
	fmt.Println("Charging ANDROID mobile")
}

type androiddapter struct {
	android Android
}

func (a *androiddapter) ChargeAppleMobile() {
	a.android.chargeAndroidMobile()
}

type Clint struct {}

func (c *Clint) ChargeMobile(mob mobile) {
	mob.ChargeAppleMobile()
}

// func main()  {
// 	a := &Apple{}
// 	c := &Clint{}

// 	c.ChargeMobile(a)

// 	android := &Android{}
// 	adapter := &androiddapter{
// 		android: *android,
// 	}

// 	c.ChargeMobile(adapter)
// }