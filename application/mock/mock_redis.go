package mock

import "context"

type MockRedis struct{}

func (u *MockRedis) Set(ctx context.Context, key string, value interface{}, exp int, duration string) error {
	return nil
}

func (u *MockRedis) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (u *MockRedis) Delete(ctx context.Context, keys ...string) error {
	return nil
}

type MockRedisSecret struct{}

func (u *MockRedisSecret) Set(ctx context.Context, key string, value interface{}, exp int, duration string) error {
	return nil
}

func (u *MockRedisSecret) Get(ctx context.Context, key string) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MTg4MTk5OTN9LCJVc2VySUQiOiIyMmIwOGU5YS1kZTI4LTRmN2ItODQ1My1hZjA2NjM5M2FmMzYiLCJVc2VybmFtZSI6ImZhamFyYXI3N0BnbWFpbC5jb20ifQ.Wwzljw_el6Dc6wNw8Bn6SnRFbSkT6ZYjbumTULPyuYo", nil
}

func (u *MockRedisSecret) Delete(ctx context.Context, keys ...string) error {
	return nil
}
