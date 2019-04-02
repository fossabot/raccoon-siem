package assets

import (
	"fmt"
	"github.com/jessevdk/go-assets"
)

// Читает файлы миграции для указанного сервиса и СУБД
func GetMigrationFiles() (files []*assets.File) {
	for _, fileName := range vfs.Dirs["/"] {
		filePath := fmt.Sprintf("/%s", fileName)
		files = append(files, vfs.Files[filePath])
	}
	return
}
