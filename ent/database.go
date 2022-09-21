package ent

import (
	"context"
	"entgo.io/ent/entc/integration/ent"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"log"
	"testing"
	"time"
)

/*type Model struct {
	ID        uuid.UUID      `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}*/

var database *Client

func Database() *Client {
	return database
}

func InitDatabase(databaseUrl string /*, config *gorm.Config*/) *Client {
	var err error
	var db *Client
	retries := 0
	for retries <= 10 {
		db, err = Open("postgres", databaseUrl)
		retries++
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources %v", err)
	}

	return db
}

func MigrateTestDatabase(db *ent.Client) {
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func InitProductionDatabase(databaseUrl string /*, config *gorm.Config*/) *Client {
	database = InitDatabase(databaseUrl /*, config*/)
	return Database()
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

func InitTestDatabase(t *testing.T, port string) *Client {
	databaseURL, mockDBID := CreateTestDatabase(port)
	db := InitDatabase(databaseURL) /*, &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	}*/
	t.Cleanup(func() { DestroyTestDatabase(mockDBID) })
	return db
}

/*func GenerateUser() dto.CreateUserDTO {
	return dto.CreateUserDTO{
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Name:      gofakeit.Name(),
		Password:  gofakeit.Password(true, true, true, true, false, 32),
		BirthDate: gofakeit.Date(),
		Avatar:    gofakeit.Person().Image,
		Biography: gofakeit.LoremIpsumParagraph(3, 5, 12, "\n"),
		Gender:    gofakeit.Gender(),
	}
}*/
