package ga

// CommonData includes all necessary data
type CommonData struct {
	// general
	Version    int    `url:"v"`
	TrackingID string `url:"tid"`

	// user
	ClientID string `url:"cid"`

	// t
	HitType string `url:"t"`

	// session
	UserIP    string `url:"uip"`
	UserAgent string `url:"ua"`

	// trafficsources
	DocumentReferer string `url:"dr,omitempty"`

	// system
	ScreenResolution string `url:"sr,omitempty"`
	ViewportSize     string `url:"vp,omitempty"`
	DocumentEncoding string `url:"de,omitempty"`
	ScreenColors     string `url:"sd,omitempty"`
	UserLanguage     string `url:"ul,omitempty"`

	// content
	DocumentLink  string `url:"dl"`
	DocumentTitle string `url:"dt,omitempty"`
}

// TimingData contains all fields of `HitType=timing`
type TimingData struct {
	PageLoadedTime     string `url:"plt,omitempty"`
	DNSTime            string `url:"dns,omitempty"`
	PageDownloadedTime string `url:"pdt,omitempty"`
	RedirectTime       string `url:"rrt,omitempty"`
	TCPTime            string `url:"tcp,omitempty"`
	ServerResponseTime string `url:"srt,omitempty"`
	DomInteractiveTime string `url:"dit,omitempty"`
	ContentLoadedTime  string `url:"clt,omitempty"`
}
