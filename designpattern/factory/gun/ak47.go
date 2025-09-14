package main 


type AK47 struct {
	Gun
}


func NewAK47() IGun {
	return &AK47{
		Gun: Gun{
			name: "ak47 Gun",
			power: 4,
		},
	}
}