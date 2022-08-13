create function set_update_time() returns opaque AS '
  	BEGIN
    	new.updated_at := ''now'';
    	return new;
  	END;
' language 'plpgsql';