create table LoanRequests(
	RID integer primary key NOT NULL AUTO_INCREMENT,
	UPIN varchar(64) not null,
	LoanType char(32) not null,
	Amount varchar(12) not null,
	AdditionalNotes varchar(256) default null,
	--
	RefObjectId integer not null,
	Reviewed boolean default False,
    --
	foreign key (UPIN) references Users(PIN),
	constraint chk_amount check
		(Amount in ('15,000,000 T', '45,000,000 T', '95,000,000 T'))
);
