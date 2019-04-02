create table if not exists users (
  id uuid not null default gen_random_uuid(),
  organization_id uuid not null,
  unit_id uuid not null,
  email string unique not null,
  phone string unique not null,
  password string not null,
  verification_code string not null,
  status int not null default 0,
  created_at int not null default 0,
  updated_at int not null default 0,
  deleted_at int not null default 0
);
