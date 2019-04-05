package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _vfs3e975935882629283f1b7c4906411eea520d4a24 = "create table destination_config (\n  id uuid primary key not null DEFAULT gen_random_uuid(),\n  name string not null,\n  payload json not null\n);"
var _vfsedf378c94580c9c38d33ad256195be84c8ef99dc = "create table dictionary_config (\n  id uuid primary key not null DEFAULT gen_random_uuid(),\n  name string not null,\n  payload json not null\n);"

// vfs returns go-assets FileSystem
var vfs = assets.NewFileSystem(map[string][]string{"/": []string{"V20190403152501__destination_config.sql", "V20190405163801__dictionary_config.sql"}}, map[string]*assets.File{
	"/V20190405163801__dictionary_config.sql": &assets.File{
		Path:     "/V20190405163801__dictionary_config.sql",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1554473208, 1554473208554605095),
		Data:     []byte(_vfsedf378c94580c9c38d33ad256195be84c8ef99dc),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1554473216, 1554473216955007226),
		Data:     nil,
	}, "/V20190403152501__destination_config.sql": &assets.File{
		Path:     "/V20190403152501__destination_config.sql",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1554461513, 1554461513721749689),
		Data:     []byte(_vfs3e975935882629283f1b7c4906411eea520d4a24),
	}}, "")
