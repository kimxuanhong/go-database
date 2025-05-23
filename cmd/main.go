package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	db2 "github.com/kimxuanhong/go-database/db"
	"github.com/kimxuanhong/go-database/repo"
	"gorm.io/gorm"
	"time"
)

// --- Ví dụ sử dụng:

type UserModel struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid"`
	PartnerId string    `gorm:"column:partner_id"`
	Total     int       `gorm:"column:total"`
	UserName  string    `gorm:"column:user_name"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UserModel) TableName() string {
	return "user_tbl"
}

func (u *UserModel) BeforeCreate(ctx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()
	return
}

func (u *UserModel) BeforeUpdate(ctx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *UserModel) GetTotal() int {
	return u.Total
}

type UserRepository struct {
	*repo.Repository[UserModel, string]
	FindByUserName                     func(ctx context.Context, username string) (*UserModel, error)                                 `repo:"@Query"`
	FindByUserNameAndEmailOrPartnerId  func(ctx context.Context, username string, email string, partnerId string) (*UserModel, error) `repo:"@Query"`
	FindAllByEmailOrderByIDDescLimit10 func(ctx context.Context, email string) ([]UserModel, error)                                   `repo:"@Query"`
}

func main() {
	// Giả lập DB (bạn thay thành *gorm.DB thật)

	var datab, _ = db2.Open(&db2.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "keycloak",
		Password: "password",
		DBName:   "keycloak",
		Schema:   "public",
		Debug:    true,
		SSLMode:  "disable",
		Driver:   "postgres",
	})

	repository := repo.NewRepository[UserModel, string](datab)
	r := &UserRepository{
		Repository: repository,
	}

	err := r.Repository.FillFuncFields(r)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Bây giờ bạn có thể gọi
	user, err := r.FindByUserName(ctx, "123")
	fmt.Println(user, err)
	user1, err := r.FindByUserNameAndEmailOrPartnerId(ctx, "123", "test@example.com", "asd")
	fmt.Println(user1, err)
	users, err := r.FindAllByEmailOrderByIDDescLimit10(ctx, "test@example.com")
	fmt.Println(users, err)

	// (GormDB chưa khởi tạo nên ví dụ này chỉ minh họa)
	fmt.Println("Repository methods injected successfully.")
}
