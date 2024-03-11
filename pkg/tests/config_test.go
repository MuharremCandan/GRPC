package tests

import (
	"os"
	"reflect"
	"test-grpc-project/pkg/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test Configuration Path
	testConfigPath := "./" // You may need to adjust this path

	// Create a test configuration file
	err := os.WriteFile(testConfigPath+"config.yaml", []byte(`
httpserver:
  host: "localhost"
  port: "8080"
grpcserver:
  host: "localhost"
  port: "50051"
database:
  host: "localhost"
  port: "5432"
  user: "testuser"
  pass: "testpass"
  name: "testdb"
`), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testConfigPath + "config.yaml")

	// Test cases
	tests := []struct {
		name       string
		args       string
		wantConfig config.Config
		wantErr    bool
	}{
		{
			name: "Valid Config",
			args: testConfigPath,
			wantConfig: config.Config{
				HttpServer: struct {
					Host string `yaml:"thost"`
					Port string `yaml:"tport"`
				}{
					Host: "localhost",
					Port: "8080",
				},
				GrpcServer: struct {
					Host string "yaml:\"ghost\""
					Port string "yaml:\"gport\""
				}{
					Host: "localhost",
					Port: "50051",
				},
				Database: struct {
					Host string "yaml:\"host\""
					Port string "yaml:\"port\""
					User string "yaml:\"user\""
					Pass string "yaml:\"pass\""
					Name string "yaml:\"name\""
				}{
					Host: "localhost",
					Port: "5432",
					User: "testuser",
					Pass: "testpass",
					Name: "testdb",
				},
			},
			wantErr: false,
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := config.LoadConfig(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("LoadConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
