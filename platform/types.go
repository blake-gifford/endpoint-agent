package platform

type Software struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type System struct {
	Name         string `json:"name"`
	Model        string `json:"model"`
	Serial       string `json:"serial"`
	Manufacturer string `json:"manufacturer"`
	Hostname     string `json:"hostname"`
}

type Data struct {
	Software   []Software `json:"software"`
	SystemInfo System     `json:"systemInfo"`
}
