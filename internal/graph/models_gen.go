// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jordanknott/taskcafe/internal/db"
)

type AddTaskLabelInput struct {
	TaskID         uuid.UUID `json:"taskID"`
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type AssignTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
	UserID uuid.UUID `json:"userID"`
}

type ChecklistBadge struct {
	Complete int `json:"complete"`
	Total    int `json:"total"`
}

type CreateTaskChecklist struct {
	TaskID   uuid.UUID `json:"taskID"`
	Name     string    `json:"name"`
	Position float64   `json:"position"`
}

type CreateTaskChecklistItem struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Name            string    `json:"name"`
	Position        float64   `json:"position"`
}

type CreateTeamMember struct {
	UserID uuid.UUID `json:"userID"`
	TeamID uuid.UUID `json:"teamID"`
}

type CreateTeamMemberPayload struct {
	Team       *db.Team `json:"team"`
	TeamMember *Member  `json:"teamMember"`
}

type DeleteInvitedProjectMember struct {
	ProjectID uuid.UUID `json:"projectID"`
	Email     string    `json:"email"`
}

type DeleteInvitedProjectMemberPayload struct {
	InvitedMember *InvitedMember `json:"invitedMember"`
}

type DeleteInvitedUserAccount struct {
	InvitedUserID uuid.UUID `json:"invitedUserID"`
}

type DeleteInvitedUserAccountPayload struct {
	InvitedUser *InvitedUserAccount `json:"invitedUser"`
}

type DeleteProject struct {
	ProjectID uuid.UUID `json:"projectID"`
}

type DeleteProjectLabel struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type DeleteProjectMember struct {
	ProjectID uuid.UUID `json:"projectID"`
	UserID    uuid.UUID `json:"userID"`
}

type DeleteProjectMemberPayload struct {
	Ok        bool      `json:"ok"`
	Member    *Member   `json:"member"`
	ProjectID uuid.UUID `json:"projectID"`
}

type DeleteProjectPayload struct {
	Ok      bool        `json:"ok"`
	Project *db.Project `json:"project"`
}

type DeleteTaskChecklist struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
}

type DeleteTaskChecklistItem struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
}

type DeleteTaskChecklistItemPayload struct {
	Ok                bool                  `json:"ok"`
	TaskChecklistItem *db.TaskChecklistItem `json:"taskChecklistItem"`
}

type DeleteTaskChecklistPayload struct {
	Ok            bool              `json:"ok"`
	TaskChecklist *db.TaskChecklist `json:"taskChecklist"`
}

type DeleteTaskGroupInput struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
}

type DeleteTaskGroupPayload struct {
	Ok           bool          `json:"ok"`
	AffectedRows int           `json:"affectedRows"`
	TaskGroup    *db.TaskGroup `json:"taskGroup"`
}

type DeleteTaskGroupTasks struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
}

type DeleteTaskGroupTasksPayload struct {
	TaskGroupID uuid.UUID   `json:"taskGroupID"`
	Tasks       []uuid.UUID `json:"tasks"`
}

type DeleteTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
}

type DeleteTaskPayload struct {
	TaskID uuid.UUID `json:"taskID"`
}

type DeleteTeam struct {
	TeamID uuid.UUID `json:"teamID"`
}

type DeleteTeamMember struct {
	TeamID     uuid.UUID  `json:"teamID"`
	UserID     uuid.UUID  `json:"userID"`
	NewOwnerID *uuid.UUID `json:"newOwnerID"`
}

type DeleteTeamMemberPayload struct {
	TeamID           uuid.UUID    `json:"teamID"`
	UserID           uuid.UUID    `json:"userID"`
	AffectedProjects []db.Project `json:"affectedProjects"`
}

type DeleteTeamPayload struct {
	Ok       bool         `json:"ok"`
	Team     *db.Team     `json:"team"`
	Projects []db.Project `json:"projects"`
}

type DeleteUserAccount struct {
	UserID     uuid.UUID  `json:"userID"`
	NewOwnerID *uuid.UUID `json:"newOwnerID"`
}

type DeleteUserAccountPayload struct {
	Ok          bool            `json:"ok"`
	UserAccount *db.UserAccount `json:"userAccount"`
}

type DuplicateTaskGroup struct {
	ProjectID   uuid.UUID `json:"projectID"`
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

type DuplicateTaskGroupPayload struct {
	TaskGroup *db.TaskGroup `json:"taskGroup"`
}

type FindProject struct {
	ProjectID uuid.UUID `json:"projectID"`
}

type FindTask struct {
	TaskID uuid.UUID `json:"taskID"`
}

type FindTeam struct {
	TeamID uuid.UUID `json:"teamID"`
}

type FindUser struct {
	UserID uuid.UUID `json:"userID"`
}

type InviteProjectMembers struct {
	ProjectID uuid.UUID      `json:"projectID"`
	Members   []MemberInvite `json:"members"`
}

type InviteProjectMembersPayload struct {
	Ok             bool            `json:"ok"`
	ProjectID      uuid.UUID       `json:"projectID"`
	Members        []Member        `json:"members"`
	InvitedMembers []InvitedMember `json:"invitedMembers"`
}

type InvitedMember struct {
	Email     string    `json:"email"`
	InvitedOn time.Time `json:"invitedOn"`
}

type InvitedUserAccount struct {
	ID        uuid.UUID   `json:"id"`
	Email     string      `json:"email"`
	InvitedOn time.Time   `json:"invitedOn"`
	Member    *MemberList `json:"member"`
}

type LogoutUser struct {
	UserID uuid.UUID `json:"userID"`
}

type MePayload struct {
	User         *db.UserAccount `json:"user"`
	TeamRoles    []TeamRole      `json:"teamRoles"`
	ProjectRoles []ProjectRole   `json:"projectRoles"`
}

type Member struct {
	ID          uuid.UUID    `json:"id"`
	Role        *db.Role     `json:"role"`
	FullName    string       `json:"fullName"`
	Username    string       `json:"username"`
	ProfileIcon *ProfileIcon `json:"profileIcon"`
	Owned       *OwnedList   `json:"owned"`
	Member      *MemberList  `json:"member"`
}

type MemberInvite struct {
	UserID *uuid.UUID `json:"userID"`
	Email  *string    `json:"email"`
}

type MemberList struct {
	Teams    []db.Team    `json:"teams"`
	Projects []db.Project `json:"projects"`
}

type MemberSearchFilter struct {
	SearchFilter string     `json:"SearchFilter"`
	ProjectID    *uuid.UUID `json:"projectID"`
}

type MemberSearchResult struct {
	Similarity int             `json:"similarity"`
	User       *db.UserAccount `json:"user"`
	Confirmed  bool            `json:"confirmed"`
	Invited    bool            `json:"invited"`
	Joined     bool            `json:"joined"`
}

type NewProject struct {
	TeamID *uuid.UUID `json:"teamID"`
	Name   string     `json:"name"`
}

type NewProjectLabel struct {
	ProjectID    uuid.UUID `json:"projectID"`
	LabelColorID uuid.UUID `json:"labelColorID"`
	Name         *string   `json:"name"`
}

type NewRefreshToken struct {
	UserID uuid.UUID `json:"userID"`
}

type NewTask struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

type NewTaskGroup struct {
	ProjectID uuid.UUID `json:"projectID"`
	Name      string    `json:"name"`
	Position  float64   `json:"position"`
}

type NewTaskGroupLocation struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Position    float64   `json:"position"`
}

type NewTaskLocation struct {
	TaskID      uuid.UUID `json:"taskID"`
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Position    float64   `json:"position"`
}

type NewTeam struct {
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organizationID"`
}

type NewUserAccount struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Initials string `json:"initials"`
	Password string `json:"password"`
	RoleCode string `json:"roleCode"`
}

type NotificationActor struct {
	ID   uuid.UUID `json:"id"`
	Type ActorType `json:"type"`
	Name string    `json:"name"`
}

type NotificationEntity struct {
	ID   uuid.UUID  `json:"id"`
	Type EntityType `json:"type"`
	Name string     `json:"name"`
}

type OwnedList struct {
	Teams    []db.Team    `json:"teams"`
	Projects []db.Project `json:"projects"`
}

type OwnersList struct {
	Projects []uuid.UUID `json:"projects"`
	Teams    []uuid.UUID `json:"teams"`
}

type ProfileIcon struct {
	URL      *string `json:"url"`
	Initials *string `json:"initials"`
	BgColor  *string `json:"bgColor"`
}

type ProjectRole struct {
	ProjectID uuid.UUID `json:"projectID"`
	RoleCode  RoleCode  `json:"roleCode"`
}

type ProjectsFilter struct {
	TeamID *uuid.UUID `json:"teamID"`
}

type RemoveTaskLabelInput struct {
	TaskID      uuid.UUID `json:"taskID"`
	TaskLabelID uuid.UUID `json:"taskLabelID"`
}

type SetTaskChecklistItemComplete struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Complete            bool      `json:"complete"`
}

type SetTaskComplete struct {
	TaskID   uuid.UUID `json:"taskID"`
	Complete bool      `json:"complete"`
}

type SortTaskGroup struct {
	TaskGroupID uuid.UUID            `json:"taskGroupID"`
	Tasks       []TaskPositionUpdate `json:"tasks"`
}

type SortTaskGroupPayload struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Tasks       []db.Task `json:"tasks"`
}

type TaskBadges struct {
	Checklist *ChecklistBadge `json:"checklist"`
}

type TaskPositionUpdate struct {
	TaskID   uuid.UUID `json:"taskID"`
	Position float64   `json:"position"`
}

type TeamRole struct {
	TeamID   uuid.UUID `json:"teamID"`
	RoleCode RoleCode  `json:"roleCode"`
}

type ToggleTaskLabelInput struct {
	TaskID         uuid.UUID `json:"taskID"`
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type ToggleTaskLabelPayload struct {
	Active bool     `json:"active"`
	Task   *db.Task `json:"task"`
}

type UnassignTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
	UserID uuid.UUID `json:"userID"`
}

type UpdateProjectLabel struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	LabelColorID   uuid.UUID `json:"labelColorID"`
	Name           string    `json:"name"`
}

type UpdateProjectLabelColor struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	LabelColorID   uuid.UUID `json:"labelColorID"`
}

type UpdateProjectLabelName struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	Name           string    `json:"name"`
}

type UpdateProjectMemberRole struct {
	ProjectID uuid.UUID `json:"projectID"`
	UserID    uuid.UUID `json:"userID"`
	RoleCode  RoleCode  `json:"roleCode"`
}

type UpdateProjectMemberRolePayload struct {
	Ok     bool    `json:"ok"`
	Member *Member `json:"member"`
}

type UpdateProjectName struct {
	ProjectID uuid.UUID `json:"projectID"`
	Name      string    `json:"name"`
}

type UpdateTaskChecklistItemLocation struct {
	TaskChecklistID     uuid.UUID `json:"taskChecklistID"`
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Position            float64   `json:"position"`
}

type UpdateTaskChecklistItemLocationPayload struct {
	TaskChecklistID uuid.UUID             `json:"taskChecklistID"`
	PrevChecklistID uuid.UUID             `json:"prevChecklistID"`
	ChecklistItem   *db.TaskChecklistItem `json:"checklistItem"`
}

type UpdateTaskChecklistItemName struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Name                string    `json:"name"`
}

type UpdateTaskChecklistLocation struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Position        float64   `json:"position"`
}

type UpdateTaskChecklistLocationPayload struct {
	Checklist *db.TaskChecklist `json:"checklist"`
}

type UpdateTaskChecklistName struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Name            string    `json:"name"`
}

type UpdateTaskDescriptionInput struct {
	TaskID      uuid.UUID `json:"taskID"`
	Description string    `json:"description"`
}

type UpdateTaskDueDate struct {
	TaskID  uuid.UUID  `json:"taskID"`
	DueDate *time.Time `json:"dueDate"`
}

type UpdateTaskGroupName struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Name        string    `json:"name"`
}

type UpdateTaskLocationPayload struct {
	PreviousTaskGroupID uuid.UUID `json:"previousTaskGroupID"`
	Task                *db.Task  `json:"task"`
}

type UpdateTaskName struct {
	TaskID uuid.UUID `json:"taskID"`
	Name   string    `json:"name"`
}

type UpdateTeamMemberRole struct {
	TeamID   uuid.UUID `json:"teamID"`
	UserID   uuid.UUID `json:"userID"`
	RoleCode RoleCode  `json:"roleCode"`
}

type UpdateTeamMemberRolePayload struct {
	Ok     bool      `json:"ok"`
	TeamID uuid.UUID `json:"teamID"`
	Member *Member   `json:"member"`
}

type UpdateUserInfo struct {
	Name     string `json:"name"`
	Initials string `json:"initials"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type UpdateUserInfoPayload struct {
	User *db.UserAccount `json:"user"`
}

type UpdateUserPassword struct {
	UserID   uuid.UUID `json:"userID"`
	Password string    `json:"password"`
}

type UpdateUserPasswordPayload struct {
	Ok   bool            `json:"ok"`
	User *db.UserAccount `json:"user"`
}

type UpdateUserRole struct {
	UserID   uuid.UUID `json:"userID"`
	RoleCode RoleCode  `json:"roleCode"`
}

type UpdateUserRolePayload struct {
	User *db.UserAccount `json:"user"`
}

type ActionLevel string

const (
	ActionLevelOrg     ActionLevel = "ORG"
	ActionLevelTeam    ActionLevel = "TEAM"
	ActionLevelProject ActionLevel = "PROJECT"
)

var AllActionLevel = []ActionLevel{
	ActionLevelOrg,
	ActionLevelTeam,
	ActionLevelProject,
}

func (e ActionLevel) IsValid() bool {
	switch e {
	case ActionLevelOrg, ActionLevelTeam, ActionLevelProject:
		return true
	}
	return false
}

func (e ActionLevel) String() string {
	return string(e)
}

func (e *ActionLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionLevel", str)
	}
	return nil
}

func (e ActionLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ActionType string

const (
	ActionTypeTaskMemberAdded ActionType = "TASK_MEMBER_ADDED"
)

var AllActionType = []ActionType{
	ActionTypeTaskMemberAdded,
}

func (e ActionType) IsValid() bool {
	switch e {
	case ActionTypeTaskMemberAdded:
		return true
	}
	return false
}

func (e ActionType) String() string {
	return string(e)
}

func (e *ActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

func (e ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ActorType string

const (
	ActorTypeUser ActorType = "USER"
)

var AllActorType = []ActorType{
	ActorTypeUser,
}

func (e ActorType) IsValid() bool {
	switch e {
	case ActorTypeUser:
		return true
	}
	return false
}

func (e ActorType) String() string {
	return string(e)
}

func (e *ActorType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActorType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActorType", str)
	}
	return nil
}

func (e ActorType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EntityType string

const (
	EntityTypeTask EntityType = "TASK"
)

var AllEntityType = []EntityType{
	EntityTypeTask,
}

func (e EntityType) IsValid() bool {
	switch e {
	case EntityTypeTask:
		return true
	}
	return false
}

func (e EntityType) String() string {
	return string(e)
}

func (e *EntityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntityType", str)
	}
	return nil
}

func (e EntityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ObjectType string

const (
	ObjectTypeOrg               ObjectType = "ORG"
	ObjectTypeTeam              ObjectType = "TEAM"
	ObjectTypeProject           ObjectType = "PROJECT"
	ObjectTypeTask              ObjectType = "TASK"
	ObjectTypeTaskGroup         ObjectType = "TASK_GROUP"
	ObjectTypeTaskChecklist     ObjectType = "TASK_CHECKLIST"
	ObjectTypeTaskChecklistItem ObjectType = "TASK_CHECKLIST_ITEM"
)

var AllObjectType = []ObjectType{
	ObjectTypeOrg,
	ObjectTypeTeam,
	ObjectTypeProject,
	ObjectTypeTask,
	ObjectTypeTaskGroup,
	ObjectTypeTaskChecklist,
	ObjectTypeTaskChecklistItem,
}

func (e ObjectType) IsValid() bool {
	switch e {
	case ObjectTypeOrg, ObjectTypeTeam, ObjectTypeProject, ObjectTypeTask, ObjectTypeTaskGroup, ObjectTypeTaskChecklist, ObjectTypeTaskChecklistItem:
		return true
	}
	return false
}

func (e ObjectType) String() string {
	return string(e)
}

func (e *ObjectType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ObjectType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ObjectType", str)
	}
	return nil
}

func (e ObjectType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleCode string

const (
	RoleCodeOwner    RoleCode = "owner"
	RoleCodeAdmin    RoleCode = "admin"
	RoleCodeMember   RoleCode = "member"
	RoleCodeObserver RoleCode = "observer"
)

var AllRoleCode = []RoleCode{
	RoleCodeOwner,
	RoleCodeAdmin,
	RoleCodeMember,
	RoleCodeObserver,
}

func (e RoleCode) IsValid() bool {
	switch e {
	case RoleCodeOwner, RoleCodeAdmin, RoleCodeMember, RoleCodeObserver:
		return true
	}
	return false
}

func (e RoleCode) String() string {
	return string(e)
}

func (e *RoleCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleCode", str)
	}
	return nil
}

func (e RoleCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleLevel string

const (
	RoleLevelAdmin  RoleLevel = "ADMIN"
	RoleLevelMember RoleLevel = "MEMBER"
)

var AllRoleLevel = []RoleLevel{
	RoleLevelAdmin,
	RoleLevelMember,
}

func (e RoleLevel) IsValid() bool {
	switch e {
	case RoleLevelAdmin, RoleLevelMember:
		return true
	}
	return false
}

func (e RoleLevel) String() string {
	return string(e)
}

func (e *RoleLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleLevel", str)
	}
	return nil
}

func (e RoleLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
