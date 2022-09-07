package models

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type Model struct {
	ID        uuid.UUID      `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

var database *gorm.DB

func Database() *gorm.DB {
	return database
}

func InitDatabase(databaseUrl string, config *gorm.Config) *gorm.DB {
	var err error
	var db *gorm.DB
	retries := 0
	for retries <= 10 {
		db, err = gorm.Open(postgres.Open(databaseUrl), config)
		retries++
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitProductionDatabase(databaseUrl string, config *gorm.Config) *gorm.DB {
	database = InitDatabase(databaseUrl, config)
	return Database()
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTestDatabase(port string) (string, string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create a Docker client: %v", err)
	}

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: "cockroachdb/cockroach",
			Cmd:   []string{"start-single-node", "--insecure"},
			ExposedPorts: nat.PortSet{
				"26257": {},
			},
		},
		&container.HostConfig{
			Binds: []string{
				"/var/run/docker.sock:/var/run/docker.sock",
			},
			PortBindings: nat.PortMap{
				"26257": []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: port,
					},
				},
			},
		}, nil, nil, "")
	if err != nil {
		log.Fatalf("Failed to create container models: %v", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("Failed to start container models: %v", err)
	}

	return "postgresql://root@127.0.0.1:" + port + "/defaultdb?sslmode=disable", resp.ID
}

func DestroyTestDatabase(mockDBID string) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts()
	if err != nil {
		log.Fatalf("Failed to create a Docker client: %v", err)
	}

	if err := cli.ContainerStop(ctx, mockDBID, nil); err != nil {
		log.Fatalf("Failed to stop container models: %v", err)
	}

	if err := cli.ContainerRemove(ctx, mockDBID, types.ContainerRemoveOptions{}); err != nil {
		log.Fatalf("Failed to remove container models: %v", err)
	}
}

func InitTestDatabase(t *testing.T, port string) *gorm.DB {
	databaseURL, mockDBID := CreateTestDatabase(port)
	db := InitDatabase(databaseURL, &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	MigrateDatabase(db)
	t.Cleanup(func() { DestroyTestDatabase(mockDBID) })
	return db
}
