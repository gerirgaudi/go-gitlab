package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// EpicsService handles communication with the epic related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/epics.html
type EpicsService struct {
	client    *Client
	timeStats *timeStatsService
}

// IssueAuthor represents a author of the epic.
type EpicAuthor struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// Epics represents a GitLab epic.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/epics.html
type Epic struct {
	ID                      int             `json:"id"`
	IID                     int             `json:"iid"`
	GroupID                 int             `json:"group_id"`
	Title                   string          `json:"string"`
	Description             string          `json:"description"`
	Author                  *EpicAuthor     `json:"author"`
	StartDate               *ISOTime      	`json:"start_date"`
	StartDateIsFixed        bool            `json:"start_date_is_fixed"`
	StartDateFixed          string          `json:"start_date_fixed"`
	StartDateFromMilestone  string          `json:"start_date_from_milestone"`
	EndDate                 *ISOTime        `json:"end_date"`
	DueDate                 *ISOTime      	`json:"end_date"`
	DueDateIsFixed          bool            `json:"end_date_is_fixed"`
	DueDateFixed            string          `json:"end_date_fixed"`
	DueDateFromMilestone    string          `json:"end_date_from_milestone"`
	UpdatedAt               *time.Time      `json:"updated_at"`
	CreatedAt               *time.Time      `json:"created_at"`
	Labels                  []string        `json:"labels"`
}

func (i Epic) String() string {
	return Stringify(i)
}

// ListGroupEpicsOptions represents the available ListGroupEpics() options.
//
// https://docs.gitlab.com/ee/api/epics.html#list-epics-for-a-group
type ListGroupEpicsOptions struct {
	ListOptions
	Id                      int 			`url:"id,omitempty" json:"id,omitempty"`
	AuthorId                int 			`url:"author_id,omitempty" json:"author_id,omitempty"`
	Labels                  Labels     		`url:"labels,comma,omitempty" json:"labels,omitempty"`
	OrderBy                 *string    		`url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                    *string    		`url:"sort,omitempty" json:"sort,omitempty"`
	Search                  *string    		`url:"search,omitempty" json:"search,omitempty"`
	State                   *string    		`url:"state,omitempty" json:"state,omitempty"`
}

// ListGroupEpics gets a list of group epics. This function accepts
// pagination parameters page and per_page to return the list of group epics.
//
// https://docs.gitlab.com/ee/api/epics.html#list-epics-for-a-group
func (s *EpicsService) ListGroupEpics(pid interface{}, opt *ListGroupEpicsOptions, options ...OptionFunc) ([]*Epic, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/epics", url.QueryEscape(group))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var i []*Epic
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}
