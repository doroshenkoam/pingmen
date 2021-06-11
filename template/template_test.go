package template

import (
	"pingmen/config"
	"testing"
	"time"

	"github.com/xanzy/go-gitlab"
)

func TestCreateMsg(t *testing.T) {
	type args struct {
		cfg  *config.Config
		tmpl string
		mr   *gitlab.MergeEvent
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				cfg: &config.Config{
					Users: struct {
						Dictionary []string `yaml:"dictionary"`
						Field      string   `yaml:"_"`
					}{
						Field: "@first @second @third",
					},
				},
				tmpl: `Project Name - {{Project Name}}
Project Description - {{Project Description}}
ObjectAttributes Action - {{ObjectAttributes Action}}
ObjectAttributes Title - {{ObjectAttributes Title}}
ObjectAttributes Description - {{ObjectAttributes Description}}
ObjectAttributes URL - {{ObjectAttributes URL}}
ObjectAttributes AuthorID - {{ObjectAttributes AuthorID}}
ObjectAttributes MergeUserID - {{ObjectAttributes MergeUserID}}
ObjectAttributes MergeError - {{ObjectAttributes MergeError}}
ObjectAttributes MergeStatus - {{ObjectAttributes MergeStatus}}
Users - {{Users}}`,
				mr: &gitlab.MergeEvent{
					Project: struct {
						ID                int                    `json:"id"`
						Name              string                 `json:"name"`
						Description       string                 `json:"description"`
						AvatarURL         string                 `json:"avatar_url"`
						GitSSHURL         string                 `json:"git_ssh_url"`
						GitHTTPURL        string                 `json:"git_http_url"`
						Namespace         string                 `json:"namespace"`
						PathWithNamespace string                 `json:"path_with_namespace"`
						DefaultBranch     string                 `json:"default_branch"`
						Homepage          string                 `json:"homepage"`
						URL               string                 `json:"url"`
						SSHURL            string                 `json:"ssh_url"`
						HTTPURL           string                 `json:"http_url"`
						WebURL            string                 `json:"web_url"`
						Visibility        gitlab.VisibilityValue `json:"visibility"`
					}{
						Name:        "My Project",
						Description: "My Description",
					},
					ObjectAttributes: struct {
						ID                       int                 `json:"id"`
						TargetBranch             string              `json:"target_branch"`
						SourceBranch             string              `json:"source_branch"`
						SourceProjectID          int                 `json:"source_project_id"`
						AuthorID                 int                 `json:"author_id"`
						AssigneeID               int                 `json:"assignee_id"`
						AssigneeIDs              []int               `json:"assignee_ids"`
						Title                    string              `json:"title"`
						CreatedAt                string              `json:"created_at"`
						UpdatedAt                string              `json:"updated_at"`
						StCommits                []*gitlab.Commit    `json:"st_commits"`
						StDiffs                  []*gitlab.Diff      `json:"st_diffs"`
						MilestoneID              int                 `json:"milestone_id"`
						State                    string              `json:"state"`
						MergeStatus              string              `json:"merge_status"`
						TargetProjectID          int                 `json:"target_project_id"`
						IID                      int                 `json:"iid"`
						Description              string              `json:"description"`
						Position                 int                 `json:"position"`
						LockedAt                 string              `json:"locked_at"`
						UpdatedByID              int                 `json:"updated_by_id"`
						MergeError               string              `json:"merge_error"`
						MergeParams              *gitlab.MergeParams `json:"merge_params"`
						MergeWhenBuildSucceeds   bool                `json:"merge_when_build_succeeds"`
						MergeUserID              int                 `json:"merge_user_id"`
						MergeCommitSHA           string              `json:"merge_commit_sha"`
						DeletedAt                string              `json:"deleted_at"`
						ApprovalsBeforeMerge     string              `json:"approvals_before_merge"`
						RebaseCommitSHA          string              `json:"rebase_commit_sha"`
						InProgressMergeCommitSHA string              `json:"in_progress_merge_commit_sha"`
						LockVersion              int                 `json:"lock_version"`
						TimeEstimate             int                 `json:"time_estimate"`
						Source                   *gitlab.Repository  `json:"source"`
						Target                   *gitlab.Repository  `json:"target"`
						LastCommit               struct {
							ID        string     `json:"id"`
							Message   string     `json:"message"`
							Timestamp *time.Time `json:"timestamp"`
							URL       string     `json:"url"`
							Author    struct {
								Name  string `json:"name"`
								Email string `json:"email"`
							} `json:"author"`
						} `json:"last_commit"`
						WorkInProgress bool                 `json:"work_in_progress"`
						URL            string               `json:"url"`
						Action         string               `json:"action"`
						OldRev         string               `json:"oldrev"`
						Assignee       gitlab.MergeAssignee `json:"assignee"`
					}{
						Action:      "open",
						Title:       "New merge",
						Description: "Somethings",
						URL:         "gitlab.com/merge_requests/1",
						AuthorID:    1,
						MergeUserID: 2,
						MergeError:  "error",
						MergeStatus: "failed",
					},
				},
			},
			want: `Project Name - My Project
Project Description - My Description
ObjectAttributes Action - open
ObjectAttributes Title - New merge
ObjectAttributes Description - Somethings
ObjectAttributes URL - gitlab.com/merge_requests/1
ObjectAttributes AuthorID - 1
ObjectAttributes MergeUserID - 2
ObjectAttributes MergeError - error
ObjectAttributes MergeStatus - failed
Users - @first @second @third`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMsg(tt.args.cfg, tt.args.tmpl, tt.args.mr); got != tt.want {
				t.Errorf("CreateMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
