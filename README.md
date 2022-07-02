# Attendance Platform
This is a simple project I used as a service for employees to be able to log their attendance. There are services for employees such as check in and check out attendance. Employees will be able to create, edit, and delete some activities for attendance. Getting activity history by date and getting attendance history are some of the other features.

## Technologies
- Go 1.18
- PostgreSQL 14
- Gin Web Framework
- GORM
- Swagger 2.0
- JSON Web Token

## Usecases
1. Employee were be able to register and login.
2. Employee were be able to check in attendance.
    - Employee were be able to add activity.
    - Employee were be able to edit activity.
    - Employee were be able to delete activity.
3. Employee were be able to check out attendance.
4. Employee were be able to get activities history by date.
5. Employee were be able to get attendance history.
6. Employee were be able to logout.

## Code Structure
The design contains several layers and components and very much similar to onion or clean architecture attempt.

### Components
1. Controllers
2. Services
3. Repositories

#### Controllers
Controllers is where all the http handlers exist. This layer is responsible to hold all the http handlers and request validation.

#### Services
Services mediates communication between a controller and repository layer. The service layer contains business logic.

#### Repositories
Repositories is for accessing the database and helps to extend the CRUD operations on the database.

![alt text](https://github.com/muhammadarash1997/attendance-platform/blob/master/Flow Chart.png?raw=true)
