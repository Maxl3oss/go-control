package models

type Site struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	SiteTitle     string `json:"site_title"`
	SiteRemark    string `json:"site_remark"`
	SiteCommand   string `json:"site_command"`
	SiteFolder    string `json:"site_folder"`
	SiteClone     string `json:"site_clone"`
	SiteDeploy    string `json:"site_deploy"`
	WebUrl        string `json:"web_url"`
	GitRepository string `json:"git_repository"`
	GitBranch     string `json:"git_branch"`
	GitToken      string `json:"git_token"`
	Type          string `json:"type"`
}
