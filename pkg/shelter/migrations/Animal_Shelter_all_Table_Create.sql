CREATE TABLE IF NOT EXISTS Animals(
    ID int PRIMARY KEY,
    Kind_Of_Animal varchar(255) NOT NULL,
    Kind_Of_Breed varchar(255) NOT NULL,
	Name varchar(255) NOT NULL,
    Age int NOT NULL,
	Description text NOT NULL,
    );

CREATE TABLE IF NOT EXISTS Shelters(
    Id          int PRIMARY KEY,
	Title       varchar(30),
	Description text,
	Address     varchar(50),
	Coordinates varchar(50)
    );

	
CREATE TABLE IF NOT EXISTS Admins(
	AdminID int PRIMARY KEY,
    Admin_Email text,
	Adminame varchar(30),
	Password text,
    Number_of_phone_Admin varchar(50),
    Profile_Picture_Admin text,
	Role varchar(50),
	foreign key(Role) references Roles(Role_name)
);
