package webhook

type JobsEvent struct {
	ObjectKind          string      `json:"object_kind"`
	Ref                 string      `json:"ref"`
	Tag                 bool        `json:"tag"`
	BeforeSha           string      `json:"before_sha"`
	Sha                 string      `json:"sha"`
	BuildID             int64       `json:"build_id"`
	BuildName           string      `json:"build_name"`
	BuildStage          string      `json:"build_stage"`
	BuildStatus         string      `json:"build_status"`
	BuildCreatedAt      string      `json:"build_created_at"`
	BuildStartedAt      string      `json:"build_started_at"`
	BuildFinishedAt     interface{} `json:"build_finished_at"`
	BuildDuration       float64     `json:"build_duration"`
	BuildQueuedDuration float64     `json:"build_queued_duration"`
	BuildAllowFailure   bool        `json:"build_allow_failure"`
	BuildFailureReason  string      `json:"build_failure_reason"`
	PipelineID          int         `json:"pipeline_id"`
	Runner              struct {
		ID          int      `json:"id"`
		Description string   `json:"description"`
		RunnerType  string   `json:"runner_type"`
		Active      bool     `json:"active"`
		IsShared    bool     `json:"is_shared"`
		Tags        []string `json:"tags"`
	} `json:"runner"`
	ProjectID   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
	User        struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Commit struct {
		ID          int         `json:"id"`
		Sha         string      `json:"sha"`
		Message     string      `json:"message"`
		AuthorName  string      `json:"author_name"`
		AuthorEmail string      `json:"author_email"`
		AuthorURL   string      `json:"author_url"`
		Status      string      `json:"status"`
		Duration    interface{} `json:"duration"`
		StartedAt   string      `json:"started_at"`
		FinishedAt  interface{} `json:"finished_at"`
	} `json:"commit"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
	Environment interface{} `json:"environment"`
}
