package tests

import (
	"assignment/config"
	"testing"
)

func BenchmarkInitDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					b.Fatalf("InitDB caused panic: %v", r)
				}
			}()
			config.InitDB()
		}()
	}
}

func BenchmarkInitRedis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					b.Fatalf("InitRedis caused panic: %v", r)
				}
			}()
			config.InitRedis()
		}()
	}
}
