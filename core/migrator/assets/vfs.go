package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _vfsa2a26de090ddd15a2a0a701ec8958c4f8e0ecbb5 = "create table destination (\n  id uuid primary key not null DEFAULT gen_random_uuid(),\n  name string not null,\n  payload json not null\n);"

// vfs returns go-assets FileSystem
var vfs = assets.NewFileSystem(map[string][]string{"/": []string{"V20190403152501__destination.sql"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1554391736, 1554391736765320601),
		Data:     nil,
	}, "/V20190403152501__destination.sql": &assets.File{
		Path:     "/V20190403152501__destination.sql",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1554391736, 1554391736764473067),
		Data:     []byte(_vfsa2a26de090ddd15a2a0a701ec8958c4f8e0ecbb5),
	}}, "")
