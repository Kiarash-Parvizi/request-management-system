create table DeclinedRequests(
	RID integer not null,
    --
	foreign key (RID) references LoanRequests(RID)
);

