CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	user_name TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create trigger trg_users_upd AFTER UPDATE ON users FOR EACH ROW
  	execute procedure set_update_time();