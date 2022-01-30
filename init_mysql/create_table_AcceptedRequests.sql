create table AcceptedRequests(
	RID integer not null,
    --
	foreign key (RID) references LoanRequests(RID)
);

