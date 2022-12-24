package main

import "fmt"

type Solder struct {
	on          bool
	ammo, power int
}

func (s *Solder) shot() bool {
	if s.ammo > 0 && s.on {
		s.ammo--
		return true
	} else {
		return false
	}
}

func (s *Solder) road() bool {
	if s.power > 0 && s.on {
		s.power--
		return true
	} else {
		return false
	}
}

func createSolder(on bool, ammo int, power int) Solder {
	return Solder{on, ammo, power}
}

func main() {
	/*testStruct := &Solder{}
	testStruct.On = true
	testStruct.ammo, testestStruct.Power = 10, 20
	*/
	Andrew := createSolder(true, 10, 20)
	fmt.Println(Andrew.shot(), Andrew.ammo)
	fmt.Println(Andrew.road(), Andrew.power)
	fmt.Println(Andrew.shot(), Andrew.ammo)
	fmt.Println(Andrew.road(), Andrew.power)

}
