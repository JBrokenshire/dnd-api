package models

import "time"

type Subject string
type Action string

const (
	SubjectUser       Subject = "User"
	SubjectRole       Subject = "Role"
	SubjectPermission Subject = "Permission"
	SubjectCharacter  Subject = "Character"
	SubjectClass      Subject = "Class"
	SubjectRace       Subject = "Race"
)

const (
	ActionCreate Action = "Create"
	ActionRead   Action = "Read"
	ActionUpdate Action = "Update"
	ActionDelete Action = "Delete"
)

// SystemPermission holds a struct of all system permissions. This lets you only add actions that are relevant to the subject. By adding in a new
// Action and Subject here, it will automatically be added to the database when the API starts.
var SystemPermission = []struct {
	Subject Subject
	Action  Action
}{
	// User
	{Subject: SubjectUser, Action: ActionCreate},
	{Subject: SubjectUser, Action: ActionRead},
	{Subject: SubjectUser, Action: ActionUpdate},
	{Subject: SubjectUser, Action: ActionDelete},
	// Role
	{Subject: SubjectRole, Action: ActionCreate},
	{Subject: SubjectRole, Action: ActionRead},
	{Subject: SubjectRole, Action: ActionUpdate},
	{Subject: SubjectRole, Action: ActionDelete},
	// Permission
	{Subject: SubjectPermission, Action: ActionRead},
	// Character
	{Subject: SubjectCharacter, Action: ActionCreate},
	{Subject: SubjectCharacter, Action: ActionRead},
	{Subject: SubjectCharacter, Action: ActionUpdate},
	{Subject: SubjectCharacter, Action: ActionDelete},
	// Class
	{Subject: SubjectClass, Action: ActionCreate},
	{Subject: SubjectClass, Action: ActionRead},
	{Subject: SubjectClass, Action: ActionUpdate},
	{Subject: SubjectClass, Action: ActionDelete},
	// Race
	{Subject: SubjectRace, Action: ActionCreate},
	{Subject: SubjectRace, Action: ActionRead},
	{Subject: SubjectRace, Action: ActionUpdate},
	{Subject: SubjectRace, Action: ActionDelete},
}

type Permission struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"-"`
	Subject   Subject   `json:"subject" gorm:"type:text"`
	Action    Action    `json:"action" gorm:"type:text"`
	Roles     []Role    `json:"Roles,omitempty" gorm:"many2many:role_permissions;"`
}
