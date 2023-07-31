package monagent

import (
	"fmt"
	"time"
)

type LogTailerConfig struct {
	TailConfigs    []TailConfig   `json:"tailConfigs" yaml:"tailConfigs"`
	RecoveryConfig RecoveryConfig `json:"recoveryConfig" yaml:"recoveryConfig"`
	// ProcessQueueCapacity Maximum capacity of file processing queue
	ProcessQueueCapacity int `json:"processQueueCapacity" yaml:"processQueueCapacity"`
}

type TailConfig struct {
	// LogDir Directory of files to be parsed (do not include file names)
	LogDir      string `json:"logDir" yaml:"logDir"`
	LogFileName string `json:"logFileName" yaml:"logFileName"`
	// ProcessLogInterval The interval at which logs are processed
	ProcessLogInterval time.Duration `json:"processLogInterval" yaml:"processLogInterval"`
	LogSourceType      string        `json:"logSourceType" yaml:"logSourceType"`
	LogAnalyzerType    string        `json:"logAnalyzerType" yaml:"logAnalyzerType"`
}

func (t TailConfig) GetLogFileRealPath() string {
	return fmt.Sprintf("%s/%s", t.LogDir, t.LogFileName)
}

type RecoveryConfig struct {
	// Enabled The function to restore the last tail location from a file is enabled
	Enabled bool `json:"enabled" yaml:"enabled"`
	// LastPositionStoreDir Persist the last queried file and the queried location information directory
	LastPositionStoreDir string `json:"lastPositionStoreDir" yaml:"lastPositionStoreDir"`
	// TriggerStoreThreshold How many lines of tail actively trigger the store action
	TriggerStoreThreshold uint64 `json:"triggerThreshold" yaml:"triggerThreshold"`
}
