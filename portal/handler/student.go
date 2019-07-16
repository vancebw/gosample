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
	log.Log("Received List Students request, Params:" + req.String())
	var students []models.Student
	err := models.GetAllStudent(&students)
	var result []*pb.Student
	for _, s := range students {
		var student pb.Student
		student.StuId = int32(s.StuId)
		student.Age = int32(s.Age)
		student.StuName = s.StuName
		student.Sex = int32(s.Sex)
		result = append(result, &student)
	}
	rep.Students = result

	if err != nil {
		return err
	}

	return nil
}
