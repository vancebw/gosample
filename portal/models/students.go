package models

type Student struct {
	StuId   int    `json:"stuId"`
	StuName string `json:"stuName"`
	Age     int    `json:"age"`
	Sex     int    `json:"sex"`
}

func GetAllStudent(s *[]Student) (err error) {
	if err = GetSlave().Find(s).Error; err != nil {
		return err
	}
	return nil
}

func GetById(s *Student, id string) (err error) {
	if err = GetSlave().Where("stu_id = ?", id).First(s).Error; err != nil {
		return err
	}
	return nil
}

func Save(s *Student) (err error) {
	if err := GetMaster().Create(s).Error; err != nil {
		return err
	}
	return nil
}

func Delete(s *Student, id string) (err error) {
	if err := GetSlave().Where("stu_id = ?", id).Delete(s).Error; err != nil {
		return err
	}
	return nil
}
