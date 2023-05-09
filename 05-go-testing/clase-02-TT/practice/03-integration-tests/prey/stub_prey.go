package prey

func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

type PreyStub struct {
	GetSpeedFn func() float64
}

func (s *PreyStub) GetSpeed() float64 {
	return s.GetSpeedFn()
}