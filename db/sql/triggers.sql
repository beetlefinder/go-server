-- Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE FUNCTION user_changed() RETURNS TRIGGER AS $user_changed$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        NEW.deleted_at := NOW();
    END IF;
    NEW.updated_at := NOW();

    RETURN NEW;
END;
$user_changed$ LANGUAGE plpgsql;

CREATE TRIGGER user_changed BEFORE INSERT OR UPDATE OR DELETE ON public.user
    FOR EACH ROW EXECUTE PROCEDURE user_changed();
