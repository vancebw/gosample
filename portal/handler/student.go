package handler

import (
	pb "gosample/portal/pb/student"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/broker"
	"gosample/portal/models"
)

const (
	studentsTopic  = "gosample.srv.student.all"
)

type UserHandler struct {
}

func (u *UserHandler) ListStudent(ctx context.Context, student *pb.LoginResponse) error {

	var students []models.Student
	err := models.GetAllStudent(&students)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(&user)
	msg := broker.Message{
		Header: map[string]string{"stuId": userId},
		Body:   []byte(b),
	}
	err = broker.Publish(studentsTopic, &msg)
	if err != nil {
		raven.CaptureError(err, nil)
	}
	return nil
}