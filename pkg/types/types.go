package types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type Guest struct {
	ID        string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Username  string             `gorm:"not null;min:1;max:20"`
	State     bitfield.Bitfield8 `gorm:"not null;default:0;"`
	CreatedAt time.Time
	ExpiresAt time.Time
}

type GuestSession struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"not null"`
	UserAgent string `gorm:"mediumtext;not null"`
	Origin    string `gorm:"mediumtext;not null"`
	IP        string `gorm:"mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	Guest *Guest `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type User struct {
	ID        string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Username  string             `gorm:"unique;not null;min:1;max:20"`
	Email     string             `gorm:"unique;not null;min:1;max:255"`
	Password  string             `gorm:"type:mediumtext"`
	Secret    string             `gorm:"type:mediumtext"`
	State     bitfield.Bitfield8 `gorm:"not null;default:0;"`
	AvatarID  *string
	BannerID  *string
	Name      string             `gorm:"type:tinytext"`
	Bio       string             `gorm:"type:mediumtext"`
	Location  string             `gorm:"type:tinytext"`
	Website   string             `gorm:"type:tinytext"`
	Language  string `gorm:"type:tinytext;default:'en';not null"`
	Theme     string `gorm:"type:tinytext;default:'system';not null"`
	IsPublic  bool   `gorm:"not null;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Avatar          *Image             `gorm:"foreignKey:AvatarID;references:ID;constraint:OnDelete:SET NULL;"`
	Banner          *Image             `gorm:"foreignKey:BannerID;references:ID;constraint:OnDelete:SET NULL;"`
	UserGameSaves   []*UserGameSave    `gorm:"foreignKey:UserID"`
	DeveloperMember []*DeveloperMember `gorm:"foreignKey:UserID"`
	GameComments    []*GameComment     `gorm:"foreignKey:UserID"`
}

type UserSession struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"not null"`
	UserAgent string `gorm:"mediumtext;not null"`
	Origin    string `gorm:"mediumtext;not null"`
	IP        string `gorm:"mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserGoogle struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserDiscord struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserGitHub struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserTOTP struct {
	UserID    string `gorm:"not null"`
	Secret    string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Verification struct {
	UserID    string `gorm:"type:char(26);not null"`
	Code      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type RecoveryCode struct {
	UserID    string `gorm:"type:char(26);not null"`
	Code      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type EmailChangeToken struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"type:char(26);not null"`
	NewEmail  string `gorm:"type:varchar(255);not null"`
	Token     string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Achievement struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID          string `gorm:"not null"`
	DeveloperGameID string `gorm:"not null"`
	Description     string `gorm:"type:tinytext;not null"`
	Points          uint64 `gorm:"not null;default:0"`
	IconID          *string
	CreatedAt       time.Time

	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
	Icon          *Image         `gorm:"foreignKey:IconID;references:ID;constraint:OnDelete:SET NULL;"`
}

type SystemEvent struct {
	ID         string `gorm:"primaryKey;type:char(26);unique;not null"`
	EventID    string
	Details    string `gorm:"type:tinytext"`
	Successful bool
	CreatedAt  time.Time

	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserEvent is used to log changes to a user account, ranging from authentication and account changes to account errors.
type UserEvent struct {
	ID         string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID     string `gorm:"not null"`
	EventID    string
	Details    string `gorm:"type:tinytext"`
	Successful bool
	CreatedAt  time.Time

	User  *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserReport represents a user-submitted report on another user.
// By specifying a ReportTag, the user can specify the kind of the report.
// If the ReportTag is null, the user can fill out the ReportDetails field for a custom report.
type UserReport struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID          string `gorm:"type:char(26);not null"`
	SubmittedUserID string `gorm:"type:char(26);not null"`
	ReportTagID     *string
	Details         string `gorm:"type:mediumtext"`
	CreatedAt       time.Time

	SubmittedUser *User      `gorm:"foreignKey:SubmittedUserID;references:ID;constraint:OnDelete:CASCADE;"`
	User          *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	ReportTag     *ReportTag `gorm:"foreignKey:ReportTagID;references:ID;constraint:OnDelete:CASCADE;"`
}

// DeveloperGameReport represents a user-submitted report on a developer's game.
// By specifying a ReportTag, the user can specify the kind of the report.
// If the ReportTag is null, the user can fill out the ReportDetails field for a custom report.
type DeveloperGameReport struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	SubmittedUserID string `gorm:"not null"`
	DeveloperGameID string `gorm:"not null"`
	ReportTagID     *string
	Details         string `gorm:"type:mediumtext"`
	CreatedAt       time.Time

	SubmittedUser *User          `gorm:"foreignKey:SubmittedUserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
	ReportTag     *ReportTag     `gorm:"foreignKey:ReportTagID;references:ID;constraint:OnDelete:CASCADE;"`
}

// DeveloperReport represents a user-submitted report on a developer's profile.
// By specifying a ReportTag, the user can specify the kind of the report.
// If the ReportTag is null, the user can fill out the ReportDetails field for a custom report.
type DeveloperReport struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	SubmittedUserID string `gorm:"not null"`
	DeveloperID     string `gorm:"not null"`
	ReportTagID     *string
	Details         string `gorm:"type:mediumtext"`
	CreatedAt       time.Time

	SubmittedUser *User      `gorm:"foreignKey:SubmittedUserID;references:ID;constraint:OnDelete:CASCADE;"`
	Developer     *Developer `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
	ReportTag     *ReportTag `gorm:"foreignKey:ReportTagID;references:ID;constraint:OnDelete:CASCADE;"`
}

// DeveloperEvent is used to log changes to a developer profile, ranging from memberships to approvals.
type DeveloperEvent struct {
	ID          string `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string
	EventID     string
	Details     string `gorm:"type:tinytext"`
	Successful  bool
	CreatedAt   time.Time

	Developer *Developer `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
	Event     *Event     `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserGameSave is used to store a user's game save data.
type UserGameSave struct {
	UserID          string `gorm:"primaryKey;type:char(26);not null"`
	DeveloperGameID string `gorm:"primaryKey;type:char(26);not null"`
	SaveSlot        uint8  `gorm:"not null;min:1;max:10"`
	SaveData        string `gorm:"type:mediumtext"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserPlayedGame is used to track games a user has played.
type UserPlayedGame struct {
	UserID          string `gorm:"primaryKey;type:char(26);not null"`
	DeveloperGameID string `gorm:"primaryKey;type:char(26);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Friend represents a friendship between two users.
type Friend struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"type:char(26);not null;index"`
	FriendID  string `gorm:"type:char(26);not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Friend *User `gorm:"foreignKey:FriendID;references:ID;constraint:OnDelete:CASCADE;"`
}

// FriendRequest represents a friend request between two users.
type FriendRequest struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	SenderID  string `gorm:"type:char(26);not null;index"`
	ReceiverID string `gorm:"type:char(26);not null;index"`
	Status    string `gorm:"type:varchar(20);not null;default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Sender   *User `gorm:"foreignKey:SenderID;references:ID;constraint:OnDelete:CASCADE;"`
	Receiver *User `gorm:"foreignKey:ReceiverID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Blocklist represents a user blocking another user.
type Blocklist struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"type:char(26);not null;index"`
	BlockedID string `gorm:"type:char(26);not null;index"`
	CreatedAt time.Time

	User    *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Blocked *User `gorm:"foreignKey:BlockedID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Message represents a chat message between two users.
type Message struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	SenderID  string `gorm:"type:char(26);not null;index"`
	ReceiverID string `gorm:"type:char(26);not null;index"`
	Content   string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Sender   *User `gorm:"foreignKey:SenderID;references:ID;constraint:OnDelete:CASCADE;"`
	Receiver *User `gorm:"foreignKey:ReceiverID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Notification represents a user notification.
type Notification struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"type:char(26);not null;index"`
	Type      string `gorm:"type:varchar(50);not null"`
	Message   string `gorm:"type:mediumtext;not null"`
	Read      bool   `gorm:"not null;default:false"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Developer represents a game developer.
type Developer struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string             `gorm:"type:tinytext;not null;default:''"`
	Description string             `gorm:"type:mediumtext"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`
	BannerID    *string
	AvatarID    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Banner           *Image             `gorm:"foreignKey:BannerID;references:ID;constraint:OnDelete:SET NULL;"`
	Avatar           *Image             `gorm:"foreignKey:AvatarID;references:ID;constraint:OnDelete:SET NULL;"`
	DeveloperMembers []*DeveloperMember `gorm:"foreignKey:DeveloperID"`
}

// DeveloperGame represents a game created by a developer.
type DeveloperGame struct {
	ID          string `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string `gorm:"type:tinytext;not null;default:''"`
	Description string `gorm:"type:mediumtext"`
	DeveloperID string
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`
	ThumbnailID *string
	CreatedAt   time.Time

	Thumbnail     *Image          `gorm:"foreignKey:ThumbnailID;references:ID;constraint:OnDelete:SET NULL;"`
	Developer     *Developer      `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
	Features      []*FeatureTag   `gorm:"many2many:developer_game_features;"`
	UserGameSaves []*UserGameSave `gorm:"foreignKey:DeveloperGameID"`
	GameComments  []*GameComment  `gorm:"foreignKey:DeveloperGameID"`
}

// GameComment represents a comment on a game by a user.
type GameComment struct {
	ID              string `gorm:"primaryKey;type:char(26);not null"`
	UserID          string `gorm:"type:char(26);not null"`
	DeveloperGameID string `gorm:"type:char(26);not null"`
	ParentID        *string
	Content         string `gorm:"type:mediumtext;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Relationships
	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
	Parent        *GameComment   `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:CASCADE;"`
	Replies       []*GameComment `gorm:"foreignKey:ParentID"`
}

// DeveloperMember represents memberships between a user and a developer account.
type DeveloperMember struct {
	UserID      string             `gorm:"not null"`
	DeveloperID string             `gorm:"not null"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`

	User      *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Developer *Developer `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Image represents an image stored on the server's hosted folder.
type Image struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	Link      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Event is a generic entity used to de-duplicate events across the system.
type Event struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null"`
	Description string `gorm:"type:tinytext"`
	LogLevel    uint8
}

// FeatureTag is a generic entity used to de-duplicate features for games.
type FeatureTag struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null;"`
	Description string `gorm:"type:tinytext"`
}

// ReportTag is a generic entity used to de-duplicate report types.
type ReportTag struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null;"`
	Description string `gorm:"type:tinytext"`
	IsUser      bool
	IsDeveloper bool
	IsGame      bool
}

// UserPoint represents a user's points balance and check-in status.
type UserPoint struct {
	UserID      string `gorm:"primaryKey;type:char(26);unique;not null"`
	Balance     int    `gorm:"not null;default:0"`
	LastCheckIn time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

// PointTransactionType represents the type of points transaction.
type PointTransactionType string

const (
	PointTransactionTypeEarn      PointTransactionType = "earn"
	PointTransactionTypeSpend     PointTransactionType = "spend"
	PointTransactionTypePurchase  PointTransactionType = "purchase"
	PointTransactionTypeTransfer  PointTransactionType = "transfer"
	PointTransactionTypeRefund    PointTransactionType = "refund"
	PointTransactionTypeAdminAdj  PointTransactionType = "admin_adjustment"
)

// PointTransaction represents a points transaction record.
type PointTransaction struct {
	ID            string             `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID        string             `gorm:"type:char(26);not null;index"`
	Amount        int                `gorm:"not null"`
	Type          PointTransactionType `gorm:"type:varchar(50);not null"`
	Description   string             `gorm:"type:mediumtext"`
	RelatedUserID *string            `gorm:"type:char(26)"`
	RelatedGameID *string            `gorm:"type:char(26)"`
	CreatedAt     time.Time

	User        *User        `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	RelatedUser *User        `gorm:"foreignKey:RelatedUserID;references:ID;constraint:OnDelete:SET NULL;"`
	RelatedGame *DeveloperGame `gorm:"foreignKey:RelatedGameID;references:ID;constraint:OnDelete:SET NULL;"`
}

// PointPurchaseStatus represents the status of a point purchase.
type PointPurchaseStatus string

const (
	PointPurchaseStatusPending   PointPurchaseStatus = "pending"
	PointPurchaseStatusCompleted PointPurchaseStatus = "completed"
	PointPurchaseStatusFailed    PointPurchaseStatus = "failed"
	PointPurchaseStatusCancelled PointPurchaseStatus = "cancelled"
)

// PointPurchase represents a points purchase transaction.
type PointPurchase struct {
	ID            string             `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID        string             `gorm:"type:char(26);not null;index"`
	Amount        int                `gorm:"not null"`
	Price         float64            `gorm:"not null"`
	Currency      string             `gorm:"type:varchar(10);not null;default:'USD'"`
	Status        PointPurchaseStatus `gorm:"type:varchar(20);not null;default:'pending'"`
	PaymentURL    string             `gorm:"type:mediumtext"`
	PaymentMethod *string            `gorm:"type:varchar(50)"`
	TransactionID *string            `gorm:"type:varchar(255)"`
	CompletedAt   *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}
