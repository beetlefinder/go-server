-- Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

DROP DATABASE IF EXISTS beetlefinder;
CREATE DATABASE beetlefinder WITH OWNER = postgres;

\c -reuse-previous=on beetlefinder

\ir ./schema.sql
\ir ./triggers.sql
