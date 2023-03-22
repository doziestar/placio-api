package models

import (
	"github.com/gofiber/fiber/v2"
	"placio-app/database"
	"time"

	"github.com/google/uuid"
)

var DB = database.DB

type Login struct {
	ID      string    `gorm:"primaryKey"`
	UserID  string    `gorm:"not null"`
	IP      string    `gorm:"not null"`
	Time    time.Time `gorm:"not null"`
	Browser string
	Device  string
}

func (l Login) Create(id string, c *fiber.Ctx) (*Login, error) {
	m := &LoginModel{}
	login, err := m.Create(id, c.IP(), c.Get("User-Agent"), c.Get("Device"))
	if err != nil {
		return nil, err
	}
	return login, nil
}

// LoginModel is the model for login schema
type LoginModel struct {
	//DB *gorm.DB
}

func (m *LoginModel) Create(userID string, ip string, browser string, device string) (*Login, error) {
	now := time.Now()

	login := &Login{
		ID:      uuid.New().String(),
		UserID:  userID,
		IP:      ip,
		Time:    now,
		Browser: browser,
		Device:  device,
	}

	if err := DB.Create(login).Error; err != nil {
		return nil, err
	}

	return login, nil
}

type LoginFlag struct {
	IP      string
	Device  string
	Browser string
}

type LoginVerification struct {
	Flag       *LoginFlag
	Level      int
	Time       string
	Suspicious bool
}

func (m *LoginModel) Verify(userID string, current *Login) (*LoginVerification, error) {
	riskLevel := 0

	flag := &LoginFlag{
		IP:      current.IP,
		Device:  current.Device,
		Browser: current.Browser,
	}

	var history []Login
	if err := DB.Where("user_id = ? AND id != ?", userID, current.ID).
		Limit(500).
		Find(&history).Error; err != nil {
		return nil, err
	}

	// if this isn't the user's first log in: perform security checks
	if len(history) > 0 {
		// has the user logged in from this IP address before?
		if findIndex(history, func(i int) bool { return history[i].IP == current.IP }) < 0 {
			riskLevel++
		}

		// has the user logged in with this browser before?
		if findIndex(history, func(i int) bool { return history[i].Browser == current.Browser }) < 0 {
			riskLevel++
		}

		// if this is a third device, set maximum risk level
		devices := len(filter(history, func(i int) bool { return history[i].Device != current.Device }))
		if devices > 1 {
			riskLevel++
		}
	}

	timeStr := current.Time.Format("2006-01-02 15:04:05")

	return &LoginVerification{
		Flag:       flag,
		Level:      riskLevel,
		Time:       timeStr,
		Suspicious: riskLevel > 0,
	}, nil
}

func findIndex(a []Login, f func(int) bool) int {
	for i, _ := range a {
		if f(i) {
			return i
		}
	}
	return -1
}

func filter(a []Login, f func(int) bool) []Login {
	filtered := make([]Login, 0)
	for i, v := range a {
		if f(i) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
