package api

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	db, env, err := callDb()

	if err != nil {
		return nil, err
	}

	db, err = checkDbConn(db)
	if err != nil {
		return nil, err
	}

	db, err = migrateDb(db)
	if err != nil {
		return nil, errorDbConn(err)
	}

	if env != "PROD" {
		seedDb(db)
	}

	return db, nil
}

func errorDbConn(err error) error {
	return fmt.Errorf("failed to connect database: %w", err)
}

func callDb() (*gorm.DB, string, error) {
	var db *gorm.DB
	var err error
	env := os.Getenv("ENVIRONMENT")

	if env == "PROD" {
		dbUrl := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	}

	if env == "STAGING" {
		dbUrl := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		db.Exec(fmt.Sprintf("DROP TABLE %v;", "product"))
		db.Exec(fmt.Sprintf("DROP TABLE %v;", "transaction"))
		db.Exec(fmt.Sprintf("DROP TABLE %v;", "room"))
		db.Exec(fmt.Sprintf("DROP TABLE %v;", "user"))
	}

	if env == "TEST" {
		db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			return nil, env, errorDbConn(err)
		}

		err = db.Exec("PRAGMA foreign_keys = ON", nil).Error
	}

	if err != nil {
		return nil, env, errorDbConn(err)
	}

	if db != nil {
		log.Println("Call DB success")
		return db, env, nil
	}

	db, err = callDbDev()

	if err != nil {
		return nil, env, errorDbConn(err)
	}

	log.Println("Call DB success")
	return db, env, nil
}

func callDbDev() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Open DB Root only for creating the intended DB
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SUPER_USER"), os.Getenv("DB_SUPER_PASSWORD"), os.Getenv("DB_ROOT"))
	dbRoot, errRoot := gorm.Open(postgres.Open(config), &gorm.Config{})

	if errRoot != nil {
		return nil, errorDbConn(errRoot)
	}

	// Implicitly silences error in case the intended DB already exists
	dbRoot.Exec(fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("DB_NAME")))

	// Close DB Root
	sqlDbRoot, errRoot := dbRoot.DB()
	if errRoot != nil {
		return nil, errRoot
	}
	errRoot = sqlDbRoot.Close()
	if errRoot != nil {
		return nil, errRoot
	}

	// Open the intended DB
	config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, errorDbConn(err)
	}

	log.Println("Call DB Dev success")
	return db, nil
}

func checkDbConn(db *gorm.DB) (*gorm.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errorDbConn(err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, errorDbConn(err)
	}

	log.Println("Check DB connection success")
	return db, nil
}

func migrateDb(db *gorm.DB) (*gorm.DB, error) {
	if err := db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{}); err != nil {
		return nil, errorDbConn(err)
	}

	log.Println("Migrate DB success")
	return db, nil
}

func seedDb(db *gorm.DB) {
	pb, _ := bcrypt.GenerateFromPassword([]byte("12345678"), 8)
	newUsers := []models.Users{
		{Nama: "Penjual", Role: "consumer", NoHp: "+6285775066878", Email: "penjual@custom.com", Password: string(pb), NoRek: "1234567890"},
		{Nama: "Pembeli", Role: "consumer", NoHp: "+6281586173213", Email: "pembeli@custom.com", Password: string(pb), NoRek: "0987654321"},
	}
	seedTable(db, &models.Users{}, &newUsers)

	penjualID := uint(1)
	pembeliID := uint(2)
	roomCodeLength := 10
	newRooms := []models.Rooms{
		{
			PenjualID: penjualID,
			PembeliID: &pembeliID,
			RoomCode:  rooms.GenerateRoomCode(roomCodeLength, 1),
			Status:    "mulai transaksi",
			Product: &models.Products{
				Nama:      "Durian",
				Deskripsi: "Durian Monthong yang sudah matang",
				Harga:     10000,
				Kuantitas: 10,
			},
		},
		{
			PenjualID: penjualID,
			PembeliID: &pembeliID,
			RoomCode:  rooms.GenerateRoomCode(roomCodeLength, 2),
			Status:    "mulai transaksi",
			Product: &models.Products{
				Nama:      "Gundam",
				Deskripsi: "Gundam PG",
				Harga:     800000,
				Kuantitas: 1,
			},
		},
		{
			PenjualID: pembeliID,
			PembeliID: &penjualID,
			RoomCode:  rooms.GenerateRoomCode(roomCodeLength, 3),
			Status:    "mulai transaksi",
			Product: &models.Products{
				Nama:      "Buah Naga",
				Deskripsi: "Buah Naga yang sudah matang",
				Harga:     5000,
				Kuantitas: 10,
			},
		},
		{
			PenjualID: pembeliID,
			PembeliID: &penjualID,
			RoomCode:  rooms.GenerateRoomCode(roomCodeLength, 4),
			Status:    "mulai transaksi",
			Product: &models.Products{
				Nama:      "Nendoroid",
				Deskripsi: "One Piece Nendoroid",
				Harga:     500000,
				Kuantitas: 1,
			},
		},
	}
	seedTable(db, &models.Rooms{}, &newRooms)
}

func seedTable(db *gorm.DB, table any, newRecords any) {
	if !db.Migrator().HasTable(table) {
		return
	}

	if err := db.First(table).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		res := db.Create(newRecords)
		if res.Error != nil {
			log.Println(res.Error.Error())
		}
	}
}
