package issue

import (
	"github.com/bikun-bikun/go-github-api/resources/repository"
	"github.com/bikun-bikun/go-github-api/resources/user"
	"time"
)

type Issue struct {
	Action string `json:"action"`
	Issue  struct {
		URL               string        `json:"url,omitempty"`
		RepositoryURL     string        `json:"repository_url,omitempty"`
		LabelsURL         string        `json:"labels_url,omitempty"`
		CommentsURL       string        `json:"comments_url,omitempty"`
		EventsURL         string        `json:"events_url,omitempty"`
		HTMLURL           string        `json:"html_url,omitempty"`
		ID                int           `json:"id,omitempty"`
		NodeID            string        `json:"node_id,omitempty"`
		Number            int           `json:"number,omitempty"`
		Title             string        `json:"title,omitempty"`
		User              user.User     `json:"user,omitempty"`
		Labels            []Label       `json:"labels,omitempty"`
		State             string        `json:"state,omitempty"`
		Locked            bool          `json:"locked,omitempty"`
		Assignee          interface{}   `json:"assignee,omitempty"`
		Assignees         []interface{} `json:"assignees,omitempty"`
		Milestone         interface{}   `json:"milestone,omitempty"`
		Comments          int           `json:"comments,omitempty"`
		CreatedAt         time.Time     `json:"created_at,omitempty"`
		UpdatedAt         time.Time     `json:"updated_at,omitempty"`
		ClosedAt          interface{}   `json:"closed_at,omitempty"`
		AuthorAssociation string        `json:"author_association,omitempty"`
		Body              string        `json:"body,omitempty"`
	} `json:"issue,omitempty"`
	Changes struct {
	} `json:"changes,omitempty"`
	Repository repository.Repository `json:"repository,omitempty"`
	Sender     Sender                `json:"sender,omitempty"`
}

type Sender struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
