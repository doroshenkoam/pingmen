package daemon

import (
	"pingmen/config"
	"sync"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xanzy/go-gitlab"
)

func TestTyp_createMsg(t1 *testing.T) {
	type fields struct {
		cfg         *config.Config
		bot         *tgbotapi.BotAPI
		wg          *sync.WaitGroup
		mrToBotChan <-chan *gitlab.MergeEvent
		doneChan    <-chan struct{}
	}
	type args struct {
		mr *gitlab.MergeEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "success",
			fields: fields{
				cfg: &config.Config{
					Users: struct {
						Dictionary []string `yaml:"dictionary"`
					}{
						Dictionary: []string{
							"first",
							"second",
						},
					},
				},
			},
			args: args{
				mr: &gitlab.MergeEvent{
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
						Description: "Что то поменяли",
						URL:         "gitlab.com/merge",
						Title:       "Merge",
						Action:      "open",
					},
				},
			},
			want: "open: Merge\ngitlab.com/merge\nЧто то поменяли\n@first @second",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Typ{
				cfg:         tt.fields.cfg,
				bot:         tt.fields.bot,
				wg:          tt.fields.wg,
				mrToBotChan: tt.fields.mrToBotChan,
				doneChan:    tt.fields.doneChan,
			}
			if got := t.createMsg(tt.args.mr); got != tt.want {
				t1.Errorf("createMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
