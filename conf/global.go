package conf

type ContextKey string

var (

	// Config global configuration holder
	Config          Conf
	CtxTwitterConf  = ContextKey("twitter")
	CtxGoogleAPIKey = ContextKey("googleAPIKey")
)

func (c ContextKey) String() string {
	return string(c)
}

// Account ...
type Account struct {
	Twitter       Twitter  `yaml:"twitter"`
	Telegram      Telegram `yaml:"telegram"`
	GoogleMapsKey string   `yaml:"googlemaps"`
}

// Telegram ...
type Telegram struct {
	Token   string `yaml:"token"`
	Webhook string `yaml:"webhook"`
	BaseURL string `yaml:"baseurl"`
}

// SectionPID is sub section of config.
type SectionPID struct {
	Enabled  bool   `yaml:"enabled"`
	Path     string `yaml:"path"`
	Override bool   `yaml:"override"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Addr            string     `yaml:"addr"`
	CommandPrefix   string     `yaml:"command_prefix"`
	Port            string     `yaml:"port"`
	MaxNotification int64      `yaml:"max_notification"`
	WorkerNum       int64      `yaml:"worker_num"`
	QueueNum        int64      `yaml:"queue_num"`
	Mode            string     `yaml:"mode"`
	SSL             bool       `yaml:"ssl"`
	CertPath        string     `yaml:"cert_path"`
	KeyPath         string     `yaml:"key_path"`
	HTTPProxy       string     `yaml:"http_proxy"`
	PID             SectionPID `yaml:"pid"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Format      string `yaml:"format"`
	AccessLog   string `yaml:"access_log"`
	AccessLevel string `yaml:"access_level"`
	ErrorLog    string `yaml:"error_log"`
	ErrorLevel  string `yaml:"error_level"`
	HideToken   bool   `yaml:"hide_token"`
}

// Conf is config structure.
type Conf struct {
	Core    SectionCore `yaml:"core"`
	Log     SectionLog  `yaml:"log"`
	Account Account     `yaml:"account"`
}
