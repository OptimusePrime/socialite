package models

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

var instance *gorm.DB

var mockDBID string

func Instance() *gorm.DB {
	return instance
}

func InitDatabase(databaseUrl string) *gorm.DB {
	var err error
	retries := 0
	for retries <= 10 {
		instance, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
		retries++
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		log.Fatal(err)
	}

	return Instance()
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTestDatabase(port string) string {
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
	mockDBID = resp.ID

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("Failed to start container models: %v", err)
	}

	return "postgresql://root@127.0.0.1:" + port + "/defaultdb?sslmode=disable"
}

func DestroyTestDatabase() {
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

func InitTestDatabase(t *testing.T, port string) {
	MigrateDatabase(InitDatabase(CreateTestDatabase(port)))
	t.Cleanup(DestroyTestDatabase)
}
