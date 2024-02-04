create function set_update_time() returns trigger AS '
  	BEGIN
    	new.updated_at := NOW();
    	return new;
  	END;
' language 'plpgsql';