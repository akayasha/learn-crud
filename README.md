# Learn CRUD Go Project

This project is a CRUD (Create, Read, Update, Delete) API for managing students, subjects, teachers, and grades. The project uses the Gin web framework to handle HTTP requests and is designed to be easily extendable and maintainable.

## Table of Contents

- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Routes](#routes)
  - [Student Routes](#student-routes)
  - [Teacher Routes](#teacher-routes)
  - [Subject Routes](#subject-routes)
  - [Grade Routes](#grade-routes)
- [Handler Functions](#handler-functions)
  - [NilaiHandler](#nilaihandler)
  - [StudentHandler](#studenthandler)
- [License](#license)

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM for Go (optional depending on the database used)
- [Swagger](https://swagger.io/) - API documentation

## Setup

To set up this project locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/learn-crud.git
    cd learn-crud
    ```

2. Install dependencies:

    ```bash
    go get github.com/gin-gonic/gin
    go get github.com/jinzhu/gorm
    go get github.com/swaggo/swag
    ```

3. Set up your database and configure it in the project (optional, depending on how the project is structured).

4. Run the project:

    ```bash
    go run main.go
    ```

    The server should now be running on `http://localhost:8080`.

## Routes

### Student Routes

- **GET /students**: Get all students
- **GET /students/nim**: Get student by NIM
- **GET /students/name**: Get students by name
- **GET /students/kelas**: Get students by class (kelas)
- **GET /students/status**: Get students by status
- **POST /students**: Create a new student
- **PUT /students/:nim**: Update a student's information
- **DELETE /students/:nim**: Delete a student

### Teacher Routes

- **GET /teachers**: Get all teachers
- **GET /teachers/:nip**: Get teacher by NIP
- **POST /teachers**: Create a new teacher
- **PUT /teachers/:nip**: Update a teacher's information
- **DELETE /teachers/:nip**: Delete a teacher
- **GET /teachers/status**: Get teachers by status

### Subject Routes

- **GET /subjects**: Get all subjects
- **GET /subjects/:id**: Get subject by ID
- **POST /subjects**: Create a new subject
- **PUT /subjects/:id**: Update a subject
- **DELETE /subjects/:id**: Delete a subject

### Grade Routes (Nilai)

- **GET /nilai**: Get grades by NIM
- **POST /nilai**: Create a new grade
- **PUT /nilai/:id**: Update a grade
- **DELETE /nilai/:id**: Delete a grade

## Handler Functions

### NilaiHandler

`NilaiHandler` provides the logic for handling routes related to grades (`/nilai`). The handler includes the following functions:

- `GetNilaiByNIM`: Get grades by a student's NIM.
- `CreateNilai`: Create a new grade.
- `UpdateNilai`: Update an existing grade.
- `DeleteNilai`: Delete a grade.

### StudentHandler

`StudentHandler` provides the logic for handling routes related to students (`/students`). The handler includes the following functions:

- `GetStudents`: Get all students.
- `GetStudentByNIM`: Get a student by NIM.
- `GetStudentByName`: Get students by name.
- `GetStudentByKelas`: Get students by class.
- `GetStudentByStatus`: Get students by status.
- `CreateStudent`: Create a new student.
- `UpdateStudent`: Update a student's information.
- `DeleteStudent`: Delete a student.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
