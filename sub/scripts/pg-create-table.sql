CREATE TABLE IF NOT EXISTS user (
	user_id SERIAL PRIMARY KEY,
	user_name TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create trigger users_update_trg AFTER UPDATE ON user FOR EACH ROW
  	execute procedure set_update_time();