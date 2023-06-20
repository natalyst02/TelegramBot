package webhook

type Issues struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		Username  string      `json:"username"`
		AvatarURL interface{} `json:"avatar_url"`
		Email     string      `json:"email"`
	} `json:"user"`
	Project struct {
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
	ObjectAttributes struct {
		AuthorID            int           `json:"author_id"`
		ClosedAt            interface{}   `json:"closed_at"`
		Confidential        bool          `json:"confidential"`
		CreatedAt           string        `json:"created_at"`
		Description         string        `json:"description"`
		DiscussionLocked    interface{}   `json:"discussion_locked"`
		DueDate             interface{}   `json:"due_date"`
		ID                  int           `json:"id"`
		Iid                 int           `json:"iid"`
		LastEditedAt        interface{}   `json:"last_edited_at"`
		LastEditedByID      interface{}   `json:"last_edited_by_id"`
		MilestoneID         interface{}   `json:"milestone_id"`
		MovedToID           interface{}   `json:"moved_to_id"`
		DuplicatedToID      interface{}   `json:"duplicated_to_id"`
		ProjectID           int           `json:"project_id"`
		RelativePosition    interface{}   `json:"relative_position"`
		StateID             int           `json:"state_id"`
		TimeEstimate        int           `json:"time_estimate"`
		Title               string        `json:"title"`
		UpdatedAt           string        `json:"updated_at"`
		UpdatedByID         interface{}   `json:"updated_by_id"`
		Weight              interface{}   `json:"weight"`
		URL                 string        `json:"url"`
		TotalTimeSpent      int           `json:"total_time_spent"`
		TimeChange          int           `json:"time_change"`
		HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
		HumanTimeChange     interface{}   `json:"human_time_change"`
		HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
		AssigneeIds         []interface{} `json:"assignee_ids"`
		AssigneeID          interface{}   `json:"assignee_id"`
		Labels              []interface{} `json:"labels"`
		State               string        `json:"state"`
		Severity            string        `json:"severity"`
		Action              string        `json:"action"`
	} `json:"object_attributes"`
	Labels  []interface{} `json:"labels"`
	Changes struct {
		AuthorID struct {
			Previous interface{} `json:"previous"`
			Current  int         `json:"current"`
		} `json:"author_id"`
		CreatedAt struct {
			Previous interface{} `json:"previous"`
			Current  string      `json:"current"`
		} `json:"created_at"`
		Description struct {
			Previous interface{} `json:"previous"`
			Current  string      `json:"current"`
		} `json:"description"`
		ID struct {
			Previous interface{} `json:"previous"`
			Current  int         `json:"current"`
		} `json:"id"`
		Iid struct {
			Previous interface{} `json:"previous"`
			Current  int         `json:"current"`
		} `json:"iid"`
		ProjectID struct {
			Previous interface{} `json:"previous"`
			Current  int         `json:"current"`
		} `json:"project_id"`
		Title struct {
			Previous interface{} `json:"previous"`
			Current  string      `json:"current"`
		} `json:"title"`
		UpdatedAt struct {
			Previous interface{} `json:"previous"`
			Current  string      `json:"current"`
		} `json:"updated_at"`
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}
