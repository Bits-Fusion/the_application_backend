package entities

type Action string

const (
	ActionView   Action = "view"
	ActionCreate Action = "create"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type Permission struct {
	ID       int32  `gorm:"primaryKey"`
	Action   Action `gorm:"type:permission_action_enum"`
	Resource string
}
