package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                   string                `gorm:"column:id;primaryKey"`
	Name                 *string               `gorm:"type:varchar(255);column:name"`
	Email                string                `gorm:"type:varchar(255);unique;not null;column:email"`
	Password             string                `gorm:"type:varchar(255);not null;column:password"`
	CreatedAt            time.Time             `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt            time.Time             `gorm:"autoUpdateTime;column:updated_at"`
	LikedWallpapers      []LikedWallpaper      `gorm:"foreignKey:UserID"`
	BookmarkedWallpapers []BookmarkedWallpaper `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "Users"
}

type Artists struct {
	gorm.Model

	ID         uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ClerkID    string      `gorm:"type:varchar(255);unique;not null"`
	Name       *string     `gorm:"type:varchar(255)"`
	Email      string      `gorm:"type:varchar(255);unique;not null"`
	MobileNo   *string     `gorm:"type:varchar(255);unique"`
	Gender     *Gender     `gorm:"type:varchar(50)"`
	Country    *string     `gorm:"type:varchar(255)"`
	Role       Role        `gorm:"type:varchar(50);default:ARTIST"`
	IsVerified bool        `gorm:"default:false"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
	Wallpapers []Wallpaper `gorm:"foreignKey:ArtistID"`
}

func (Artists) TableName() string {
	return "Artists"
}

type Wallpaper struct {
	gorm.Model

	ID                   uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title                *string               `gorm:"type:varchar(255)"`
	ImageURL             *string               `gorm:"type:varchar(255)"`
	Likes                int                   `gorm:"default:0"`
	Platform             *Platform             `gorm:"type:varchar(50)"`
	CategoriesID         uuid.UUID             `gorm:"type:uuid"`
	Categories           Category              `gorm:"foreignKey:CategoriesID;constraint:OnDelete:CASCADE"`
	ArtistID             uuid.UUID             `gorm:"type:uuid"`
	Artist               Artists               `gorm:"foreignKey:ArtistID;constraint:OnDelete:CASCADE"`
	CreatedAt            time.Time             `gorm:"autoCreateTime"`
	UpdatedAt            time.Time             `gorm:"autoUpdateTime"`
	LikedWallpapers      []LikedWallpaper      `gorm:"foreignKey:WallpaperID"`
	BookmarkedWallpapers []BookmarkedWallpaper `gorm:"foreignKey:WallpaperID"`
}

func (Wallpaper) TableName() string {
	return "Wallpapers"
}

type LikedWallpaper struct {
	gorm.Model

	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	WallpaperID uuid.UUID `gorm:"type:uuid"`
	Wallpaper   Wallpaper `gorm:"foreignKey:WallpaperID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (LikedWallpaper) TableName() string {
	return "LikedWallpapers"
}

type BookmarkedWallpaper struct {
	gorm.Model

	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	WallpaperID uuid.UUID `gorm:"type:uuid"`
	Wallpaper   Wallpaper `gorm:"foreignKey:WallpaperID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (BookmarkedWallpaper) TableName() string {
	return "BookmarkedWallpapers"
}

type Category struct {
	gorm.Model

	ID         uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       *string     `gorm:"type:varchar(255)"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
	Wallpapers []Wallpaper `gorm:"foreignKey:CategoriesID"`
}

func (Category) TableName() string {
	return "Categories"
}

type Platform string

const (
	Mobile  Platform = "MOBILE"
	Tablet  Platform = "TABLET"
	Desktop Platform = "DESKTOP"
)

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type Role string

const (
	Artist Role = "ARTIST"
	Admin  Role = "ADMIN"
)
