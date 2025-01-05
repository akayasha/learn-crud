package main

import (
	"learn-crud/database"
	"learn-crud/handlers"
	"learn-crud/models"
	"learn-crud/repository"
	"learn-crud/routes"
	"learn-crud/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "learn-crud/docs"             // Import Swagger docs
)

// @title Learn CRUD API
// @version 1.0
// @description API for managing CRUD operations for students, teachers, subjects, and grades.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// Connect to database
	database.Connect()

	// Drop existing tables to ensure clean migration
	database.DB.Exec("SET FOREIGN_KEY_CHECKS=0")

	// Drop tables in reverse order of dependencies
	database.DB.Migrator().DropTable(&models.Nilai{})
	database.DB.Migrator().DropTable(&models.Subject{})
	database.DB.Migrator().DropTable(&models.Student{})
	database.DB.Migrator().DropTable(&models.Teacher{})

	// Create tables in order of dependencies
	err := database.DB.AutoMigrate(
		&models.Teacher{}, // No dependencies
		&models.Student{}, // No dependencies
		&models.Subject{}, // Depends on Teacher
		&models.Nilai{},   // Depends on Student and Subject
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Add the foreign key constraint manually
	database.DB.Exec(`
        ALTER TABLE nilais 
        ADD CONSTRAINT fk_nilais_student 
        FOREIGN KEY (nim) 
        REFERENCES students(nim) 
        ON DELETE CASCADE
    `)

	database.DB.Exec("SET FOREIGN_KEY_CHECKS=1")

	// Initialize repository, service, and handler
	studentRepo := repository.NewStudentRepository(database.DB)
	studentService := services.NewStudentService(studentRepo)
	studentHandler := handlers.NewStudentHandler(studentService)

	nilaiRepo := repository.NewNilaiRepository(database.DB)
	nilaiService := services.NewNilaiService(nilaiRepo)
	nilaiHandler := handlers.NewNilaiHandler(nilaiService)

	subjectsRepo := repository.NewSubjectRepository(database.DB)
	subjectService := services.NewSubjectService(subjectsRepo)
	subjectHandler := handlers.NewSubjectHandler(subjectService)

	teacherRepo := repository.NewTeacherRepository(database.DB)
	teacherService := services.NewTeacherService(teacherRepo)
	teacherHandler := handlers.NewTeacherHandler(teacherService)

	// Setup router
	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.StudentRoutes(router, studentHandler)
	routes.NilaiRoutes(router, nilaiHandler)
	routes.SubjectRoutes(router, subjectHandler)
	routes.TeacherRoutes(router, teacherHandler)

	// Run server
	router.Run(":8080")
}
