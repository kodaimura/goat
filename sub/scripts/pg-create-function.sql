create function set_update_time() returns trigger AS '
  	BEGIN
    	new.updated_at := ''now'';
    	return new;
  	END;
' language 'plpgsql';