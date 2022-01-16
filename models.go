package variti

type Service struct {
	Id           int
	Name         string
	IsSSL        int      `json:"ssl"`
	AdditionalIP []string `json:"dop_ip,omitempty"`
	IP           string
	IPv6         string `json:"ipv6,omitempty"`
	IsWAF        int    `json:"waf,omitempty"`
	IsEnabled    int    `json:"enable"`
	ABPMode      int    `json:"abp_mode"`
	CreatedAt    int    `json:"created"`
}

type ServiceItem struct {
	ID                 int
	Name               string
	HSTS               int
	IP                 string
	IPv6               string   `json:"ipv6,omitempty"`
	IsEnable           int      `json:"enable"`
	CreatedAt          int      `json:"created"`
	BlockBotLevel      int      `json:"block_bot_level"`
	IsForceSSL         int      `json:"forcessl"`
	IsSSL              int      `json:"ssl"`
	IsSSLAuto          int      `json:"ssl_auto"`
	SSLPrivateKey      string   `json:"ssl_key"`
	SSLCert            string   `json:"ssl_crt"`
	IsGlobalWhitelist  int      `json:"global_whitelist"`
	IsBotProtect       int      `json:"botprotect"`
	IsHttpsToHttp      int      `json:"https2http"`
	AdditionalIP       []string `json:"dop_ip,omitempty"`
	IsHTTP2            int      `json:"http2"`
	GeoIPMode          int      `json:"geoip_mode"`
	GeoIPList          string   `json:"geoip_list"`
	CDNProxyHost       string   `json:"cdn_proxy_host"`
	CDNHost            string   `json:"cdn_host"`
	IsCDN              int      `json:"cdn"`
	IsWAF              int      `json:"waf"`
	WAFMode            int      `json:"waf_mode"`
	IAuthEndpointPauth int      `json:"iauth_endpoint_pauth"`
	IAuthDisableL7     int      `json:"iauth_disable_l7"`
	ABPMode            int      `json:"abp_mode"`
	IPHash             int      `json:"iphash"`
	IsWWWRedirect      int      `json:"wwwredir"`
	IsWildAlias        int      `json:"wildalias"`
}

type Origin struct {
	ID         int
	Weight     int
	Mode       string
	DestIP     string `json:"data"`
	CreatedAt  int    `json:"created"`
	ModifiedAt int    `json:"modified"`
}
