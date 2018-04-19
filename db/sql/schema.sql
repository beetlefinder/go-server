-- Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE public.user (
    id          SERIAL      PRIMARY KEY CHECK(id > 0),
    nick        TEXT        NOT NULL,
    is_deleted  BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE TABLE public.auth (
    id          INTEGER     REFERENCES public.user(id),
    login       TEXT        NOT NULL,
    pass_hash   TEXT        NOT NULL,
    is_deleted  BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE TABLE public.alert (
    id          SERIAL      PRIMARY KEY CHECK(id > 0),
    id_user     INTEGER     NOT NULL REFERENCES public.user(id),
    message     TEXT,   
    is_deleted  BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);