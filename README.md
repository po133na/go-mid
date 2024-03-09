# Project Description

Our team has made the project dedicated to Animal Shelters. In this project we have implemented Backend (Go, PostgreSQL).

 This is the API where people can look for pets they can adopt (full data) and also give up a pet.

There are 4 instances: Shelter, User, Role and Animal. 

Our team have made a database for pets, containing description data and functional API.   

## Team Members
 
Polina Stelmakh 22B030588

Anel Tulepbergen 22B030602

Alina Amreyeva 22B031240

Dossym Ibray  22B030545


## Build

Run go run main.go in the integrated terminal. 

## API Structure

Animal Shelter REST API
```
POST /animals
GET /animals/:id
PUT /animals/:id
DELETE /animals/:id
```

## Instances' Structure
```
![Image alt](https://github.com/po133na/go-mid/mip.png)

Users
• UserID (AUTO_INCREMENT PRIMARY KEY)
• User_Email(TEXT)
• Username (VARCHAR(30))
• Password (encrypted password)
• Number_of_phone_user(VARCHAR(50))
• Role(Default “User”)
Animal
• ID(AUTO_INCREMENT PRIMARY KEY)
• Kind_Of_Animal(VARCHAR(255))
• Breed_Of_Animal(VARCHAR(255))
• Name(VARCHAR(255))
• Age(INTEGER)
• Description(TEXT)
Admins
• AdminID(AUTO_INCREMENT PRIMARY KEY)
• Admin_Email(TEXT)
• Adminame (VARCHAR(30))
• Password (encrypted password)
• Number_of_phone_Admin(VARCHAR(50))
• Role(VARCHAR(255), Type:Back_end, Front_end, G_admin, etc)
Role
• role_name(TEXT)
• Permissions(TEXT)
```
