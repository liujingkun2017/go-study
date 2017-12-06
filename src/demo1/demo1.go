package demo1


func less1() {

}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
