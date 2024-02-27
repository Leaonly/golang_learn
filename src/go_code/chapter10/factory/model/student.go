package model

type student struct{
	Name string
	score float64
}

func NewStudent(n string, s float64)*student{
	return &student{
		Name: n,
		score: s,
	}
}

//如果score字段首字母小写，则，在其它包不可以直接方法，我们可以提供一个方法
func (s *student)GetScore()float64{
	return s.score
}