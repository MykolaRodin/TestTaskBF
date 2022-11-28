## Q4

### Description
Write a user management project which will include:
  - A master view which will list all users in a data grid. This
screen will assist users with all CRUD operations. User will
be able to press 3 buttons (New,Edit,Delete). Edit and Delete
operations will require row selection from the data grid.
  - A detailed view which will show the fields as form. Form will
have 2 buttons (Action, Back). Text of the action button will
change according to the operation opened the detail view. For
example if the “New” operation is selected from the master,
the detailed view action button text will be “Create”. Please
see the mappings below.
    - New: Create
    - Edit: Save
    - Delete: Delete
  - A REST service to support functions below. Please note that
API paths and HTTP methods and HTTP Statuses are important
for us.
    - Returns all users
    - Return the user with the desired “id”
    - Save the given user.
    - Update data of the user with the desired “id”
    - Delete the user with the desired “id”
  - Backend must be written with Go. You are free to choose any
database you desire. Remember all operations must be
persistent.
  - Frontend must be written with JS/TS using React.

### PostgreSQL setup
1. Open new terminal
2. Run `docker run -d --rm --name postgres -e POSTGRES_PASSWORD=secretpassword -p 5432:5432 postgres:14` to run PostgreSQL image in a container
3. Run `docker exec -ti -u postgres postgres psql` to start working in the container
4. Run `CREATE DATABASE gousers;` to create `gousers` database
5. Run `\c gousers` to connect to the database
6. Run
```
CREATE TABLE users (
  id INTEGER PRIMARY KEY UNIQUE GENERATED ALWAYS AS IDENTITY,
  first_name VARCHAR(200) NOT NULL,
  last_name VARCHAR(200) NOT NULL,
  email VARCHAR(200) DEFAULT NULL
);
```
to create `users` table inside `gousers` database.

### Backend setup
1. Open new terminal with project directory
2. Run `cd backend` to move to backend directory
3. Run `go run .` to run the backend Go application

### Frontend setup
1. Open new terminal with project directory
2. Run `cd frontend` to move to backend directory
3. Run `npm install` to install all dependencies
4. Run `npm start` to start the frontend React application
