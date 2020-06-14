CREATE TABLE IF NOT EXISTS users (
    id varying(1024) PRIMARY KEY,
    login character varying(1024) NOT NULL UNIQUE,
	first_name character varying(1024) NOT NULL,
	middle_name character varying(1024) NOT NULL,
	last_name character varying(1024) NOT NULL,
    sex character(1) NOT NULL DEFAULT 'u',
	email character varying(1024) NOT NULL,
	phone character varying(16) NOT NULL,
    weight smallint,
    height smallint,
    pwhash character varying(1024) NOT NULL
);

COMMENT ON COLUMN users.sex IS 'u=undefined, m=male, f=female';
