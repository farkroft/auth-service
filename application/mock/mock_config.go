package mock

type MockConfig struct{}

func (c *MockConfig) GetString(str string) string {
	return ""
}

func (c *MockConfig) GetInt(str string) int {
	return 0
}
