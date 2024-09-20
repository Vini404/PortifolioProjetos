drop table if exists AccountHolder CASCADE;
drop table if exists Customer CASCADE;
drop table if exists Account CASCADE;
drop table if exists Balance CASCADE;
drop table if exists Transactions CASCADE;

create table Customer(
	ID SERIAL,
	FullName varchar(100) not null,
	Phone varchar(50) not null,
	Email varchar(255) not null,
	Birthday timestamp not null,
	CreatedTimeStamp timestamp not null,
	UpdatedTimeStamp timestamp,
	
	primary KEY(ID)
);

create table AccountHolder (
	ID SERIAL,
	IDCustomer int not null,
	IsActive bool not null,
	CreatedTimeStamp timestamp not null,
	UpdatedTimeStamp timestamp,
	
	primary KEY(ID),
	constraint fk_customer foreign KEY(IDCustomer) references Customer(ID)
);

create table Account (
	ID SERIAL,
	IDAccountHolder int not null,
	IsActive bool not null,
	CreatedTimeStamp timestamp not null,
	UpdatedTimeStamp timestamp,
	
	primary KEY(ID),
	constraint fk_accountholder foreign KEY(IDAccountHolder) references AccountHolder(ID)
);


create table Balance (
	ID SERIAL,
	IDAccount int not null,
	Ammount decimal not null,
	AmmountBlocked decimal not null,
	UpdatedTimeStamp timestamp,
	
	primary KEY(ID),
	constraint fk_account foreign KEY(IDAccount) references Account(ID)
);

create table Transactions (
	ID SERIAL,
	IDDebitAccount int not null,
	IDCreditAccount int not null,
	Ammount decimal not null,
	CreatedTimeStamp timestamp not null,
	TransactionType integer not null,
	primary KEY(ID),
	constraint fk_account_debit foreign KEY(IDDebitAccount) references Account(ID),
	constraint fk_account_credit foreign KEY(IDCreditAccount) references Account(ID)
);


