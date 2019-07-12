package handler

import (
	pb "cinema/user/pb/user"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/broker"
)

const (
	userCreateTopic = "com.cinema.srv.user.created"
	userLoginTopic  = "com.cinema.srv.user.login"
)

type UserHandler struct {
}

func (u *UserHandler) ListStudent(ctx context.Context, req *pb.LoginRequest, user *pb.LoginResponse) error {
	doc := bsonx.Doc{}
	err := module.DB.Collection("user").FindOne(context.Background(),
		bson.D{{"username", req.Username}}).Decode(&doc)
	if err != nil {
		return err
	}
	validate := util.ValidatePassword(doc.Lookup("password").String(), req.Password)
	if !validate {
		return fmt.Errorf("密码错误")
	}
	userId := doc.Lookup("_id").ObjectID().Hex()
	user.UserId = userId
	user.Nickname = doc.Lookup("nickname").StringValue()
	user.Username = doc.Lookup("username").StringValue()
	user.Phone = doc.Lookup("phone").StringValue()
	i, _ := doc.Lookup("birthday").Int64OK()
	user.Birthday = i
	b, _ := json.Marshal(&user)
	msg := broker.Message{
		Header: map[string]string{"userId": userId},
		Body:   []byte(b),
	}
	err = broker.Publish(userLoginTopic, &msg)
	if err != nil {
		raven.CaptureError(err, nil)
	}
	return nil
}