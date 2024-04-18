package config

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	var jsonExample = []byte(`{
"server": {
    "api_port": 9000,
    "global_timeout": 30000,
    "api_log_level": "debug"
  },
  "database": {
    "go-backend": {
      "master": {
        "driver": "postgres",
        "url": "user=test password=test dbname=test host=localhost port=5432 sslmode=disable",
        "max_idle": 30,
        "max_open_connection": 100,
        "connection_max_life_time": 10
      },
      "slave": {
        "driver": "postgres",
        "url": "user=test password=test dbname=test host=localhost port=5432 sslmode=disable",
        "max_idle": 30,
        "max_open_connection": 100,
        "connection_max_life_time": 10
      }
    }
  }
}
`)
	ConfigMap = viper.New()
	ConfigMap.SetConfigType("json")
	ConfigMap.ReadConfig(bytes.NewBuffer(jsonExample))
	// get default int
	assert.Equal(t, 10, getIntOrDefault("not found", 10))

	//get default string
	assert.Equal(t, "default", getStringOrDefault("not found", "default"))

	// get panic when string not found
	assert.Panics(t, func() {
		getStringOrPanic("not found")
	})

	// get panic when int not found
	assert.Panics(t, func() {
		getIntOrPanic("not found")
	})

	// get value from config when there is the value (string)
	assert.Equal(t, "default", getStringOrDefault("database.go.master.driver", "default"))

}

// testLoadFromEnvOrConfigEnv test whether the env var will replace
// existing value inside the config file
func testLoadFromEnvOrConfigEnv(v *viper.Viper) error {
	prevVal := v.GetString("SERVER_GLOBAL_TIMEOUT")
	expected := 20000

	// create matching env var from config file
	// and set the value to be different than the one
	// inside the config file
	os.Setenv("SERVER_GLOBAL_TIMEOUT", strconv.Itoa(expected))

	envVal := v.GetInt("SERVER_GLOBAL_TIMEOUT")
	if envVal != expected {
		return fmt.Errorf("unexpected value from SERVER_GLOBAL_TIMEOUT: expecting %d, got %d", expected, envVal)
	}

	// just in case if the above os.Setenv() effect is persistent, we need to return
	// the env var to its original value
	os.Setenv("SERVER_GLOBAL_TIMEOUT", prevVal)

	return nil
}

func Test_loadFromEnvAndConfigEnv(t *testing.T) {
	type args struct {
		path string
		name string
	}
	tests := []struct {
		name    string
		args    args
		tFn     func(*viper.Viper) error
		wantErr bool
	}{
		{
			name: "Load config file and env vars",
			args: args{
				path: "../../config",
			},
			tFn:     testLoadFromEnvOrConfigEnv,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadFromEnvAndConfigEnv(tt.args.path, "config")
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFromEnvAndConfigEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := tt.tFn(got); err != nil {
				t.Error(err)
				return
			}
		})
	}
}
