package typestruct

type Bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

// make a new Bill --> Like Constructor
func ConBill(n string, i map[string]float64, t float64) Bill {
	b := Bill{
		Name:  n,
		Items: i,
		Tip:   t,
	}
	return b
}
