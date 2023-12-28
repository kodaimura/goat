CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create trigger trg_users_upd BEFORE UPDATE ON users FOR EACH ROW
  	execute procedure set_update_time();