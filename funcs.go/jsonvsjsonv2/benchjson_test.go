package tjson

import (
	"encoding/json"
	"testing"
	"time"

	jsonv2 "encoding/json/v2"
)

// Structs used for benchmarking
type BenchmarkUser struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Email    string                 `json:"email"`
	Age      int                    `json:"age"`
	Active   bool                   `json:"active"`
	Created  time.Time              `json:"created"`
	Tags     []string               `json:"tags"`
	Metadata map[string]interface{} `json:"metadata"`
}

type BenchmarkUserV2 struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Age      int            `json:"age"`
	Active   bool           `json:"active"`
	Created  time.Time      `json:"created"`
	Tags     []string       `json:"tags"`
	Metadata map[string]any `json:"metadata"`
}

// Sample instance for tests
var testUser = BenchmarkUser{
	ID:      1,
	Name:    "Jeff Otoni from Testing for Benchmark",
	Email:   "jeff.teste@exemplo.com.br",
	Age:     35,
	Active:  true,
	Created: time.Now(),
	Tags:    []string{"admin", "user", "premium", "beta", "tester"},
	Metadata: map[string]interface{}{
		"department": "information Technology",
		"salary":     7500.50,
		"projects":   []string{"projeto1", "projeto2", "projeto3"},
		"skills":     []string{"Go", "JavaScript", "Python", "Docker"},
		"certifications": map[string]string{
			"aws":    "solutions-architect",
			"google": "cloud-engineer",
			"azure":  "fundamentals",
		},
	},
}

var testUserV2 = BenchmarkUserV2{
	ID:      1,
	Name:    "Jeff Otoni from Testing for Benchmark",
	Email:   "jeff.teste@exemplo.com.br",
	Age:     35,
	Active:  true,
	Created: time.Now(),
	Tags:    []string{"admin", "user", "premium", "beta", "tester"},
	Metadata: map[string]any{
		"department": "information Technology",
		"salary":     7500.50,
		"projects":   []string{"projeto1", "projeto2", "projeto3"},
		"skills":     []string{"Go", "JavaScript", "Python", "Docker"},
		"certifications": map[string]string{
			"aws":    "solutions-architect",
			"google": "cloud-engineer",
			"azure":  "fundamentals",
		},
	},
}

// Sample JSON string for Unmarshal benchmarks
var testJSON = `{
	"id": 1,
	"name": "Jeff Otoni from Testing for Benchmark",
	"email": "jeff.teste@exemplo.com.br",
	"age": 35,
	"active": true,
	"created": "2024-01-15T10:30:00Z",
	"tags": ["admin", "user", "premium", "beta", "tester"],
	"metadata": {
		"department": "information Technology",
		"salary": 7500.50,
		"projects": ["projeto1", "projeto2", "projeto3"],
		"skills": ["Go", "JavaScript", "Python", "Docker"],
		"certifications": {
			"aws": "solutions-architect",
			"google": "cloud-engineer",
			"azure": "fundamentals"
		}
	}
}`

// Benchmarks for Marshal
func BenchmarkMarshalV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(testUser)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonv2.Marshal(testUserV2)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for Unmarshal
func BenchmarkUnmarshalV1(b *testing.B) {
	jsonData := []byte(testJSON)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var user BenchmarkUser
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalV2(b *testing.B) {
	jsonData := []byte(testJSON)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var user BenchmarkUserV2
		err := jsonv2.Unmarshal(jsonData, &user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for Marshal with large slice
func BenchmarkMarshalSliceV1(b *testing.B) {
	users := make([]BenchmarkUser, 1000)
	for i := range users {
		users[i] = testUser
		users[i].ID = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(users)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalSliceV2(b *testing.B) {
	users := make([]BenchmarkUserV2, 1000)
	for i := range users {
		users[i] = testUserV2
		users[i].ID = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonv2.Marshal(users)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for Unmarshal with large slice
func BenchmarkUnmarshalSliceV1(b *testing.B) {
	// Create a large JSON array
	users := make([]BenchmarkUser, 100)
	for i := range users {
		users[i] = testUser
		users[i].ID = i
	}
	jsonData, _ := json.Marshal(users)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []BenchmarkUser
		err := json.Unmarshal(jsonData, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalSliceV2(b *testing.B) {
	// Create a large JSON array
	users := make([]BenchmarkUserV2, 100)
	for i := range users {
		users[i] = testUserV2
		users[i].ID = i
	}
	jsonData, _ := jsonv2.Marshal(users)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []BenchmarkUserV2
		err := jsonv2.Unmarshal(jsonData, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for round-trip (Marshal + Unmarshal)
func BenchmarkRoundTripV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(testUser)
		if err != nil {
			b.Fatal(err)
		}

		var user BenchmarkUser
		err = json.Unmarshal(data, &user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRoundTripV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := jsonv2.Marshal(testUserV2)
		if err != nil {
			b.Fatal(err)
		}

		var user BenchmarkUserV2
		err = jsonv2.Unmarshal(data, &user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for complex map structures
func BenchmarkMarshalMapV1(b *testing.B) {
	complexMap := map[string]interface{}{
		"users": []BenchmarkUser{testUser, testUser, testUser},
		"config": map[string]interface{}{
			"debug":     true,
			"timeout":   30,
			"endpoints": []string{"api1", "api2", "api3"},
		},
		"metadata": map[string]interface{}{
			"version": "1.0.0",
			"build":   12345,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(complexMap)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalMapV2(b *testing.B) {
	complexMap := map[string]any{
		"users": []BenchmarkUserV2{testUserV2, testUserV2, testUserV2},
		"config": map[string]any{
			"debug":     true,
			"timeout":   30,
			"endpoints": []string{"api1", "api2", "api3"},
		},
		"metadata": map[string]any{
			"version": "1.0.0",
			"build":   12345,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonv2.Marshal(complexMap)
		if err != nil {
			b.Fatal(err)
		}
	}
}
