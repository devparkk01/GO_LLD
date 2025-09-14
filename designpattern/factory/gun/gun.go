package main 

type Gun struct {
	name string 
	power int 
}

func(t *Gun) setName(name string) {
	t.name = name 
}

func (t *Gun) getName() string {
	return t.name 
}

func (t *Gun) setPower(power int) {
	t.power = power
}

func (t *Gun) getPower() int {
	return t.power
}