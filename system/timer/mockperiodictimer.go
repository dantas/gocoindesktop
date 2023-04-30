package timer

import "time"

type MockTimer struct {
	channel          chan struct{}
	SetIntervalValue time.Duration
}

func NewMockTimer() *MockTimer {
	return &MockTimer{
		channel: make(chan struct{}),
	}
}

func (m *MockTimer) Tick() <-chan struct{} {
	return m.channel
}

func (m *MockTimer) ForceTick() {
	m.channel <- struct{}{}
}

func (m *MockTimer) SetInterval(interval time.Duration) {
	m.SetIntervalValue = interval
}

func (m *MockTimer) Destroy() {
	close(m.channel)
}
