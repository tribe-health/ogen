// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func (s *ActionsCreateOrUpdateOrgSecretApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsCreateSelfHostedRunnerGroupForOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsEnterprisePermissions) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsListOrgSecrets) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsListRepoWorkflows) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsListSelfHostedRunnersForOrg) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsListSelfHostedRunnersForRepo) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsListSelfHostedRunnersInGroupForOrg) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsOrganizationPermissions) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsRepositoryPermissions) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsReviewPendingDeploymentsForRunApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsSetGithubActionsPermissionsOrganizationApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsSetGithubActionsPermissionsRepositoryApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ActionsUpdateSelfHostedRunnerGroupForOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AppPermissions) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AppsCreateContentAttachmentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'body' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    262144,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Body)); err != nil {
			failures = append(failures, validate.FieldError{Name: "body", Error: err})
		}
	}
	{
		// Validate 'title' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    1024,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Title)); err != nil {
			failures = append(failures, validate.FieldError{Name: "title", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AppsCreateInstallationAccessTokenApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AppsScopeTokenApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AuthenticationToken) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Authorization) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *BranchProtection) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *BranchWithProtection) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CheckRun) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CheckSuite) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ChecksListSuitesForRef) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CodeScanningAlertInstance) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CodeScanningAnalysis) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CodeScanningSarifsStatus) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CodeScanningUpdateAlertApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CodeScanningUploadSarifApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CommitComment) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CommitComparison) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ContentReferenceAttachment) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'body' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    262144,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Body)); err != nil {
			failures = append(failures, validate.FieldError{Name: "body", Error: err})
		}
	}
	{
		// Validate 'title' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    1024,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Title)); err != nil {
			failures = append(failures, validate.FieldError{Name: "title", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *DeploymentStatus) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'description' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    140,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Description)); err != nil {
			failures = append(failures, validate.FieldError{Name: "description", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *DiffEntry) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EnterpriseAdminListSelfHostedRunnersForEnterprise) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EnterpriseAdminSetGithubActionsPermissionsEnterpriseApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EnvironmentApprovals) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *FullRepository) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *FullRepositorySecurityAndAnalysis) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *FullRepositorySecurityAndAnalysisAdvancedSecurity) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *FullRepositorySecurityAndAnalysisSecretScanning) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GistComment) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'body' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    65535,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Body)); err != nil {
			failures = append(failures, validate.FieldError{Name: "body", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GistsCreateCommentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'body' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    65535,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Body)); err != nil {
			failures = append(failures, validate.FieldError{Name: "body", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GistsUpdateCommentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'body' field.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    65535,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Body)); err != nil {
			failures = append(failures, validate.FieldError{Name: "body", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GitCreateTagApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GitCreateTreeApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GitCreateTreeApplicationJSONRequestTreeItem) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GitRef) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GitRefObject) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'sha' field.
		validator := validate.String{
			MinLength:    40,
			MinLengthSet: true,
			MaxLength:    40,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Sha)); err != nil {
			failures = append(failures, validate.FieldError{Name: "sha", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Import) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *InstallationToken) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *InteractionLimit) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *InteractionLimitResponse) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *IssueComment) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *IssuesCreateMilestoneApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *IssuesLockApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *IssuesUpdateMilestoneApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Job) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *JobStepsItem) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MarkdownRenderApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MergedUpstream) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MigrationsSetLfsPreferenceApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MigrationsStartForAuthenticatedUserApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MigrationsStartForOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *MigrationsStartImportApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Milestone) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *NullableScopedInstallation) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *OrgMembership) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *OrganizationActionsSecret) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *OrgsCreateInvitationApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *OrgsSetMembershipForUserApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *OrgsUpdateApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *OrgsUpdateMembershipForAuthenticatedUserApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Page) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PagesHTTPSCertificate) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Project) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ProjectsAddCollaboratorApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ProjectsUpdateApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ProtectedBranchPullRequestReview) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullRequestReview) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullRequestReviewComment) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullsCreateReviewApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullsCreateReviewCommentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullsMergeApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullsSubmitReviewApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PullsUpdateApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Reaction) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForCommitCommentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForIssueApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForIssueCommentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForPullRequestReviewCommentApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForReleaseApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForTeamDiscussionCommentInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForTeamDiscussionCommentLegacyApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForTeamDiscussionInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReactionsCreateForTeamDiscussionLegacyApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Release) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReleaseAsset) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposAddCollaboratorApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposCreateCommitStatusApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposCreateDeploymentStatusApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposCreateDispatchEventApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'event_type' field.
		validator := validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    100,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.EventType)); err != nil {
			failures = append(failures, validate.FieldError{Name: "event_type", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposCreateInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposCreatePagesSiteApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposCreatePagesSiteApplicationJSONRequestSource) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposUpdateApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReposUpdateInvitationApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *RepositoryInvitation) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ReviewComment) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Runner) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *RunnerLabelsItem) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ScimProvisionAndInviteUserApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ScimSetInformationForProvisionedUserApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SecretScanningAlert) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SecretScanningUpdateAlertApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ShortBranch) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamFull) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamMembership) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsAddOrUpdateMembershipForUserInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsAddOrUpdateMembershipForUserLegacyApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsAddOrUpdateProjectPermissionsInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsAddOrUpdateProjectPermissionsLegacyApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsAddOrUpdateRepoPermissionsInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsAddOrUpdateRepoPermissionsLegacyApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsCreateApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsUpdateInOrgApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *TeamsUpdateLegacyApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *UsersSetPrimaryEmailVisibilityForAuthenticatedApplicationJSONRequest) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Workflow) Validate() error {
	var failures []validate.FieldError
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
