package helpers

import (
	"go-agreenery/constants"
	"go-agreenery/models"
	"log"
	"math/rand"
	"strings"
	"time"

	"gorm.io/gorm"
)

func SendWateringScheduleNotifications(db *gorm.DB) {
	messages := []string{
		"Jangan lupa siram tanaman!",
		"Tanaman haus!",
		"Waktu menyiram tanaman!",
		"Akar tanaman pasti senang jika disiram.",
		"Daun-daun terlihat layu, butuh air!",
		"Yuk, berikan tanaman kesayanganmu seteguk air!",
		"Jangan sampai tanaman kesayanganmu mati kehausan!",
		"Media tanam terlihat kering, saatnya menyiram.",
		"Siram tanaman secara merata agar semua bagian tercukupi.",
		"Dengan menyiram teratur, tanaman akan tumbuh subur.",
		"Tanaman berbunga akan lebih cantik jika terawat dengan baik.",
	}

	now := time.Now()
	today := now.Weekday().String()

	schedules := models.ListWateringSchedule{}
	if err := db.Where("start_date <= ? AND end_date >= ? AND turn_on_notif = true", now, now).Find(&schedules).Error; err != nil {
		log.Printf("error fetching watering schedules %v", err)
		return
	}

	var days []string
	for _, schedule := range schedules {
		days = strings.Split(schedule.RepeatEvery, ",")
		if containsDay(days, today) {
			notification := models.UserNotification{
				UserID:    schedule.UserID,
				Title:     randomMessage(messages),
				Subtitle:  constants.GeneralSubtitle,
				ActionURL: schedule.ID,
				Icon:      "https://storage.googleapis.com/agreenery/uploads/agreenery-logo.png",
			}

			if err := db.Create(&notification).Error; err != nil {
				log.Printf("error creating notification %v", err)
				return
			}
		}
	}
}

func SendAdminNotifications(db *gorm.DB) {
	notifications := models.ListNotification{}
	if err := db.Where("is_sent = false AND DATE(send_at) = CURDATE()").Find(&notifications).Error; err != nil {
		log.Printf("error fetching watering schedules %v", err)
		return
	}

	users := models.ListUser{}
	if err := db.Model(&users).
		Joins("INNER JOIN credentials ON users.credential_id = credentials.id").Where("credentials.role = 'user'").
		Find(&users).Error; err != nil {
		log.Printf("error fetching users %v", err)
		return
	}

	for _, notif := range notifications {
		for _, user := range users {
			notification := models.UserNotification{
				UserID:    user.ID,
				Title:     notif.Title,
				Subtitle:  notif.Subtitle,
				ActionURL: notif.ActionURL,
				Icon:      "https://storage.googleapis.com/agreenery/uploads/agreenery-logo.png",
			}

			if err := db.Create(&notification).Error; err != nil {
				log.Printf("error creating notification %v", err)
				return
			}
		}

		if err := db.Model(&notif).Where("id = ?", notif.ID).Update("is_sent", true).Error; err != nil {
			log.Printf("error updating notification %v", err)
			return
		}
	}
}

func containsDay(days []string, day string) bool {
	for _, d := range days {
		if strings.Trim(d, " ") == day {
			return true
		}
	}

	return false
}

func randomMessage(messages []string) string {
	index := rand.Intn(len(messages))
	return messages[index]
}
