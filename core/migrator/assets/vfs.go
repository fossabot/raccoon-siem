package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _vfs3e975935882629283f1b7c4906411eea520d4a24 = "create table destination_config (\n  id uuid primary key not null DEFAULT gen_random_uuid(),\n  name string not null,\n  payload json not null\n);"

// vfs returns go-assets FileSystem
var vfs = assets.NewFileSystem(map[string][]string{"/": []string{"V20190403152501__destination_config.sql"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1554461513, 1554461513723966818),
		Data:     nil,
	}, "/V20190403152501__destination_config.sql": &assets.File{
		Path:     "/V20190403152501__destination_config.sql",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1554461513, 1554461513721749689),
		Data:     []byte(_vfs3e975935882629283f1b7c4906411eea520d4a24),
	}}, "")
