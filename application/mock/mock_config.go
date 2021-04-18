package mock

type MockConfig struct{}

func (c *MockConfig) GetString(str string) string {
	return ""
}

func (c *MockConfig) GetInt(str string) int {
	return 0
}

type MockConfigSecretKey struct{}

func (c *MockConfigSecretKey) GetString(str string) string {
	return "your secret"
}

func (c *MockConfigSecretKey) GetInt(str string) int {
	return 0
}
