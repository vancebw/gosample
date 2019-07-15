package handler

import (
	"context"
	"github.com/micro/go-log"
	"gosample/portal/models"
	pb "gosample/portal/proto"
)

type StudentHandler struct {
}

func (s *StudentHandler) CreateStudent(context.Context, *pb.Student, *pb.Response) error {
	return nil
}

func (s *StudentHandler) ListStudent(ctx context.Context, req *pb.Request, rep *pb.ListResponse) error {
	log.Log("Received List Students request")
	var students []*pb.Student
	err := models.GetAllStudent(students)
	if err != nil {
		return err
	}
	rep.Students = students
	return nil
}
