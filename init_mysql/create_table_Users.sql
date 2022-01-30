create table Users(
	FullName    varchar(64),
	PhoneNumber varchar(32),
	PIN         varchar(32) primary key,
	UserPhoto   BLOB default null,
    --
	Credits     Integer,
	Pass        varchar(64),
	AccessLevel smallint
);