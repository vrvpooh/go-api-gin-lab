package services

import (
	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) error {
	// เรียกใช้ฟังก์ชัน Update ที่เราเพิ่งสร้างใน Repo
	return s.Repo.Update(id, student)
}

func (s *StudentService) DeleteStudent(id string) error {
	// เรียกใช้ฟังก์ชัน Delete ที่เราเพิ่งสร้างใน Repo
	return s.Repo.Delete(id)
}
