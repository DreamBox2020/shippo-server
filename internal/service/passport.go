package service

import "shippo-server/utils/box"

func (s *Service) PassportCreate(c *box.Context, passport string) (interface{}, error) {
	p, err := s.dao.GetPassport(passport)

	if err != nil {
		return nil, err
	}

	// 如果不存在，就创建一个新的通行证
	if p.Token == "" {

	}

	data := make(map[string]interface{}, 2)
	data["passport"] = p.Token
	data["uid"] = p.UserId

	return data, err
}
