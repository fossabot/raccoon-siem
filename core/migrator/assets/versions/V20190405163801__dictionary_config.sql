create table dictionary_config (
  id uuid primary key not null DEFAULT gen_random_uuid(),
  name string not null,
  payload json not null
);