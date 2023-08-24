package entities

type Device struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	UserID   string `json:"user_id"`

	OSName         string `json:"os_name"`
	OSArch         string `json:"os_arch"`
	MacAddress     string `json:"mac_address"`
	LocalIPAddress string `json:"local_ip_address"`
	Port           string `json:"port"`
	FetchedUnix    int64  `json:"fetched_unix"`
	CountryCode    string `json:"country_code"`
	FlagImage      string `json:"flag_image"`
	RAMSize        string `json:"ram_size"`
	GPU            string `json:"gpu"`
	CPU            string `json:"cpu"` // Ek olarak CPU bilgisini de ekledik.
}
