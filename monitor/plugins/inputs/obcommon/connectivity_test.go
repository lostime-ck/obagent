package obcommon

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestConnectable(t *testing.T) {
	configStr := `
        targets:
          t1: 'user:pass@tcp(127.0.0.1:9878)/oceanbase?timeout=5s'
    `
	var connectivityConfigMap map[string]interface{}
	_ = yaml.Unmarshal([]byte(configStr), &connectivityConfigMap)

	connectivityInput := &ConnectivityInput{}
	connectivityInput.Init(context.Background(), connectivityConfigMap)

	metrics, _ := connectivityInput.CollectMsgs()
	require.Equal(t, 1, len(metrics))
	value, exists := metrics[0].GetField("value")
	v, ok := value.(float64)
	require.True(t, exists)
	require.True(t, ok)
	require.Equal(t, 1.0, v)
}

func TestNotConnectable(t *testing.T) {
	configStr := `
        targets:
          t1: 'user:pass@tcp(0.0.0.0:2881)/oceanbase?timeout=5s'
    `
	var connectivityConfigMap map[string]interface{}
	_ = yaml.Unmarshal([]byte(configStr), &connectivityConfigMap)

	connectivityInput := &ConnectivityInput{}
	connectivityInput.Init(context.Background(), connectivityConfigMap)

	metrics, _ := connectivityInput.CollectMsgs()
	connectivityInput.Stop()
	require.Equal(t, 1, len(metrics))
	value, exists := metrics[0].GetField("value")
	v, ok := value.(float64)
	require.True(t, exists)
	require.True(t, ok)
	require.Equal(t, 0.0, v)
}

func TestDescription(t *testing.T) {
	configStr := `
        targets:
          t1: 'user:pass@tcp(0.0.0.0:9878)/oceanbase?timeout=5s'
    `
	var connectivityConfigMap map[string]interface{}
	_ = yaml.Unmarshal([]byte(configStr), &connectivityConfigMap)

	connectivityInput := &ConnectivityInput{}
	connectivityInput.Init(context.Background(), connectivityConfigMap)

	desStr := connectivityInput.Description()
	require.Equal(t, desStr, connectivityDescription)
}

func TestSampleConfig(t *testing.T) {
	configStr := `
        targets:
          t1: 'user:pass@tcp(0.0.0.0:9878)/oceanbase?timeout=5s'
    `
	var connectivityConfigMap map[string]interface{}
	_ = yaml.Unmarshal([]byte(configStr), &connectivityConfigMap)

	connectivityInput := &ConnectivityInput{}
	connectivityInput.Init(context.Background(), connectivityConfigMap)

	cfgStr := connectivityInput.SampleConfig()
	require.Equal(t, cfgStr, connectivitySampleConfig)
}
