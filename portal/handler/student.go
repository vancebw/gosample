package handler

import (
	"context"
	"gosample/portal/models"
	pb "gosample/portal/proto"
)

const (
	studentsTopic = "gosample.srv.student.all"
)

type UserHandler struct {
}

func (u *UserHandler) ListStudent(ctx context.Context, req *pb.Request, rep *pb.StudentResponse) error {

	var students []models.Student
	err := models.GetAllStudent(&students)
	if err != nil {
		return err
	}
	return nil
}
