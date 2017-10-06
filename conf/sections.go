package conf

import "encoding/json"

// RedisSection Redis instance information
type RedisSection struct {
	Host   string `json:"host" split-default-value:"localhost" split-cli-option:"redis-host" split-cli-description:"Redis server hostname"`
	Port   int    `json:"port" split-default-value:"6379" split-cli-option:"redis-port" split-cli-description:"Redis Server port"`
	Db     int    `json:"db" split-default-value:"0" split-cli-option:"redis-db" split-cli-description:"Redis DB"`
	Pass   string `json:"password" split-default-value:"" split-cli-option:"redis-pass" split-cli-description:"Redis password"`
	Prefix string `json:"prefix" split-default-value:"" split-cli-option:"redis-prefix" split-cli-description:"Redis key prefix"`

	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `json:"network" split-default-value:"tcp" split-cli-option:"redis-network" split-cli-description:"Redis network protocol"`

	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int `json:"maxRetries" split-default-value:"0" split-cli-option:"redis-max-retries" split-cli-description:"Redis connection max retries"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout int `json:"dialTimeout" split-default-value:"5" split-cli-option:"redis-dial-timeout" split-cli-description:"Redis connection dial timeout"`

	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is 10 seconds.
	ReadTimeout int `json:"readTimeout" split-default-value:"10" split-cli-option:"redis-read-timeout" split-cli-description:"Redis connection read timeout"`

	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is 3 seconds.
	WriteTimeout int `json:"writeTimeout" split-default-value:"5" split-cli-option:"redis-write-timeout" split-cli-description:"Redis connection write timeout"`

	// Maximum number of socket connections.
	// Default is 10 connections.
	PoolSize int `json:"poolSize" split-default-value:"10" split-cli-option:"redis-pool" split-cli-description:"Redis connection pool size"`
}

// LogSection log instance configuration
type LogSection struct {
	VerboseOn       bool   `json:"verbose" split-default-value:"false" split-cli-option:"log-verbose" split-cli-description:"Enable verbose mode"`
	DebugOn         bool   `json:"debug" split-default-value:"false" split-cli-option:"log-debug" split-cli-description:"Enable debug mode"`
	StdoutOn        bool   `json:"stdout" split-default-value:"false" split-cli-option:"log-stdout" split-cli-description:"Enable log standard output"`
	File            string `json:"file" split-default-value:"/tmp/split-agent.log" split-cli-option:"log-file" split-cli-description:"Set the log file"`
	FileMaxSize     int64  `json:"fileMaxSizeBytes" split-cli-option:"log-file-max-size" split-default-value:"2000000" split-cli-description:"Max file log size in bytes"`
	FileBackupCount int    `json:"fileBackupCount" split-cli-option:"log-file-backup-count" split-default-value:"3" split-cli-description:"Number of last log files to keep in filesystem"`
	SlackChannel    string `json:"slackChannel" split-default-value:"" split-cli-option:"log-slack-channel" split-cli-description:"Set the Slack channel or user"`
	SlackWebhookURL string `json:"slackWebhookURL" split-default-value:"" split-cli-option:"log-slack-webhook-url" split-cli-description:"Set the Slack webhook url"`
}

// ConfigData main configuration container
type ConfigData struct {
	APIKey              string             `json:"apiKey" split-cli-option:"api-key" split-default-value:"YOUR API KEY" split-cli-description:"Your Split API-KEY"`
	Proxy               InMemorySection    `json:"proxy" split-cli-option-group:"true"`
	Redis               RedisSection       `json:"redis" split-cli-option-group:"true"`
	Producer            ProducerSection    `json:"sync" split-cli-option-group:"true"`
	Logger              LogSection         `json:"log" split-cli-option-group:"true"`
	ImpressionListener  ImpressionListener `json:"impressionListener" split-cli-option-group:"true"`
	SplitsFetchRate     int                `json:"splitsRefreshRate" split-cli-option:"split-refresh-rate" split-default-value:"60" split-cli-description:"Refresh rate of splits fetcher"`
	SegmentFetchRate    int                `json:"segmentsRefreshRate" split-default-value:"60" split-cli-option:"segment-refresh-rate" split-cli-description:"Refresh rate of segments fetcher"`
	ImpressionsPostRate int                `json:"impressionsRefreshRate" split-default-value:"60" split-cli-option:"impressions-post-rate" split-cli-description:"Post rate of impressions recorder"`
	ImpressionsPerPost  int64              `json:"impressionsPerPost" split-cli-option:"impressions-per-post" split-default-value:"1000" split-cli-description:"Number of impressions to send in a POST request"`
	ImpressionsThreads  int                `json:"impressionsThreads" split-default-value:"1" split-cli-option:"impressions-recorder-threads" split-cli-description:"Number of impressions recorder threads"`
	MetricsPostRate     int                `json:"metricsRefreshRate" split-default-value:"60" split-cli-option:"metrics-post-rate" split-cli-description:"Post rate of metrics recorder"`
	HTTPTimeout         int64              `json:"httpTimeout" split-default-value:"60" split-cli-option:"http-timeout" split-cli-description:"Timeout specifies a time limit for requests"`
}

//MarshalBinary exports ConfigData to JSON string
func (c ConfigData) MarshalBinary() (data []byte, err error) {
	return json.MarshalIndent(c, "", "  ")
}

// ProducerSection wrapper for all producer configurations
type ProducerSection struct {
	ProducerAdmin ProducerAdmin `json:"admin" split-cli-option-group:"true"`
	// TODO migrate Redis into this section.
	//Redis RedisSection `json:"redis" split-cli-option-group:"true"`
}

// ProducerAdmin represents configuration for sync admin endpoints
type ProducerAdmin struct {
	Port          int    `json:"port" split-default-value:"3010" split-cli-option:"sync-admin-port" split-cli-description:"Sync admin port to listen connections"`
	AdminUsername string `json:"adminUsername" split-default-value:"" split-cli-option:"sync-admin-username" split-cli-description:"HTTP basic auth username for admin endpoints"`
	AdminPassword string `json:"adminPassword" split-default-value:"" split-cli-option:"sync-admin-password" split-cli-description:"HTTP basic auth password for admin endpoints"`
}

// InMemorySection represents configuration for in memory proxy
type InMemorySection struct {
	Port               int    `json:"port" split-default-value:"3000" split-cli-option:"proxy-port" split-cli-description:"Proxy port to listen connections"`
	AdminPort          int    `json:"adminPort" split-default-value:"3010" split-cli-option:"proxy-admin-port" split-cli-description:"Proxy port for admin endpoints"`
	AdminUsername      string `json:"adminUsername" split-default-value:"" split-cli-option:"proxy-admin-username" split-cli-description:"HTTP basic auth username for admin endpoints"`
	AdminPassword      string `json:"adminPassword" split-default-value:"" split-cli-option:"proxy-admin-password" split-cli-description:"HTTP basic auth password for admin endpoints"`
	PersistMemoryPath  string `json:"persistInFilePath" split-default-value:"" split-cli-option:"proxy-mmap-path" split-cli-description:"File path to persist memory in proxy mode"`
	ImpressionsMaxSize int64  `json:"impressionsMaxSize" split-default-value:"10485760" split-cli-option:"proxy-impressions-max-size" split-cli-description:"Max size, in bytes, to send impressions in proxy mode"`
	Auth               Auth   `json:"auth" split-cli-option-group:"true"`
}

// Auth struct for proxy authentication
type Auth struct {
	// ApiKeys list of alloweb API-Keys for SDKs
	// split-default-value must be set as SDK_API_KEY just to write config file by cli (see func getDefaultConfigData() at parser.go)
	APIKeys []string `json:"sdkAPIKeys" split-default-value:"SDK_API_KEY" split-cli-option:"proxy-apikeys" split-cli-description:"List of allowed custom API Keys for SDKs"`
}

type ImpressionListener struct {
	Endpoint string `json:"endpoint" split-default-value:"" split-cli-option:"impression-listener-endpoint" split-cli-description:"HTTP endpoint where impression bulks will be posted"`
}
