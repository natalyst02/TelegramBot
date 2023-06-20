package webhook

type DeploymentEventsLoad struct {
	ObjectKind      string `json:"object_kind"`
	Status          string `json:"status"`
	StatusChangedAt string `json:"status_changed_at"`
	DeploymentID    int    `json:"deployment_id"`
	DeployableID    int    `json:"deployable_id"`
	DeployableURL   string `json:"deployable_url"`
	Environment     string `json:"environment"`
	Project         struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ShortSha string `json:"short_sha"`
	User     struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	UserURL     string `json:"user_url"`
	CommitURL   string `json:"commit_url"`
	CommitTitle string `json:"commit_title"`
}
