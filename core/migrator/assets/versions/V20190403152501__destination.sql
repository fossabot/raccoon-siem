create table destination (
  id uuid primary key not null DEFAULT gen_random_uuid(),
  name string not null,
  kind string not null,
  url string not null default '',
  subject string not null default '',
  index_name string not null default '',
  static_index bool not null default false,
  batch_size int not null default 0
);