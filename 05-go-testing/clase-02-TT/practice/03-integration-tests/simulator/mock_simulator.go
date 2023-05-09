package simulator

import "github.com/stretchr/testify/mock"

// constructor
func NewCatchSimulatorMockTestify() *CatchSimulatorMockTestify {
	return &CatchSimulatorMockTestify{}
}

type CatchSimulatorMockTestify struct {
	mock.Mock
}

func (m *CatchSimulatorMockTestify) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	args := m.Called(distance, speed, catchSpeed)

	r0 := args.Bool(0)

	return r0
}
func (m *CatchSimulatorMockTestify) GetLinearDistance(position [2]float64) float64 {
	args := m.Called(position)

	r0 := args.Get(0).(float64)

	return r0
}

// ------------------------------------------------------------
// manually
type CatchSimulatorMock struct {
	MethodCanCatch struct {Distance float64; Speed float64; CatchSpeed float64; Ret0 bool}
	MethodGetLinearDistance struct {Position [2]float64; Ret0 float64}
	// spy
	Spys map[string]bool
}

func (m *CatchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	// spy
	m.Spys["CanCatch"] = true

	m.MethodCanCatch.Distance = distance
	m.MethodCanCatch.Speed = speed
	m.MethodCanCatch.CatchSpeed = catchSpeed
	return m.MethodCanCatch.Ret0
}

func (m *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	// spy
	m.Spys["GetLinearDistance"] = true

	m.MethodGetLinearDistance.Position = position
	return m.MethodGetLinearDistance.Ret0
}