package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets8d719ec162ec583ada54591e2074c0338d05401a = "create table if not exists users (\n  id uuid not null default gen_random_uuid(),\n  organization_id uuid not null,\n  unit_id uuid not null,\n  email string unique not null,\n  phone string unique not null,\n  password string not null,\n  verification_code string not null,\n  status int not null default 0,\n  created_at int not null default 0,\n  updated_at int not null default 0,\n  deleted_at int not null default 0\n);\n"

// Assets returns go-assets FileSystem
var vfs = assets.NewFileSystem(map[string][]string{"/": []string{"V20190219140301__users.sql"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1553710826, 1553710826631304405),
		Data:     nil,
	}, "/V20190219140301__users.sql": &assets.File{
		Path:     "/V20190219140301__users.sql",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1550668714, 1550668714000000000),
		Data:     []byte(_Assets8d719ec162ec583ada54591e2074c0338d05401a),
	}}, "")
