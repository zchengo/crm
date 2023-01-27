package config

// 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Jwt    Jwt    `mapstructure:"jwt"`
	File   File   `mapstructure:"file"`
	Mail   Mail   `mapstructure:"mail"`
	Alipay Alipay `mapstructure:"alipay"`
}

// 服务端启动配置
type Server struct {
	Port   int    `mapstructure:"port"`
	Runenv string `mapstructure:"runenv"`
}

// MySQL数据库配置
type Mysql struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

// Redis数据库配置
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

// JWT用户认证配置
type Jwt struct {
	SigningKey  string `mapstructure:"signingKey"`
	ExpiredTime int    `mapstructure:"expiredTime"`
}

// 文件存储配置
type File struct {
	Path string `mapstructure:"path"`
}

// 邮件服务配置
type Mail struct {
	Smtp   string `mapstructure:"smtp"`
	Secret string `mapstructure:"secret"`
	Sender string `mapstructure:"sender"`
}

// 支付宝支付服务配置
type Alipay struct {
	AppId            string `mapstructure:"appId"`
	PrivateKey       string `mapstructure:"privateKey"`
	AppPublicCert    string `mapstructure:"appPublicCert"`
	AlipayRootCert   string `mapstructure:"alipayRootCert"`
	AlipayPublicCert string `mapstructure:"alipayPublicCert"`
	ReturnURL        string `mapstructure:"returnURL"`
	NotifyURL        string `mapstructure:"notifyURL"`
}
