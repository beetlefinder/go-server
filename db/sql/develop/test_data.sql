-- Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

INSERT INTO public.user (nick) VALUES
('voledemar'),
('test_2'),
('test_3');

INSERT INTO public.auth (login, pass_hash) VALUES
('voledemar', '1234hashMD5loool'),
('test_2_login', 'ergsrthbr8muJMGae85g'),
('my_login', 'rwagw3t4tfr');
