package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	// 1. ดึง ID จาก URL parameter (:id)
	id := c.Param("id")

	// 2. รับข้อมูล JSON จาก Request Body
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. ส่งข้อมูลไปให้ Service ทำงาน
	if err := h.Service.UpdateStudent(id, student); err != nil {
		// ถ้าไม่พบ ID ให้ส่ง 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// 4. ถ้าสำเร็จ ส่งข้อมูลที่อัปเดตแล้วกลับไปพร้อม Status 200 OK
	student.Id = id
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	// 1. รับ ID จาก URL Parameter
	id := c.Param("id")

	// 2. เรียก Service ให้ลบข้อมูล
	if err := h.Service.DeleteStudent(id); err != nil {
		// หากไม่พบ ID หรือลบไม่สำเร็จ ให้ส่ง 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// 3. หากลบสำเร็จ ให้ส่ง Status 204 No Content (ไม่มีเนื้อหาตอบกลับ)
	c.Status(http.StatusNoContent)
}
