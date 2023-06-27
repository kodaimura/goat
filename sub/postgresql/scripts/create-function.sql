create function set_update_time() returns trigger AS '
  	BEGIN
    	new.update_at := NOW();
    	return new;
  	END;
' language 'plpgsql';