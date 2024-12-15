package benchmark

import (
	"go/go-backend-api/global"
	"go/go-backend-api/internal/initialize"
	"go/go-backend-api/internal/po"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

func TestMain(m *testing.M) {
	// Change working directory to project root
	projectRoot := filepath.Join("..", "..") // Adjust path according to your project structure
	if err := os.Chdir(projectRoot); err != nil {
		panic(err)
	}
	initialize.LoadConfig()
	initialize.InitLogger()
	initialize.InitMysql()
	// Run tests
	os.Exit(m.Run())
}

func inSertRecord(b *testing.B) {
	uuid, _ := uuid.NewUUID()
	user := po.User{
		UUID: uuid,
	}
	err := global.MyDB.Create(&user).Error
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkInsert(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			inSertRecord(b)
		}
	})
}
