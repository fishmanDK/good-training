package service

import "time"



type Service struct{
	
} 

func NewService() *Service{
	return &Service{}
}

func (s *Service) TimeToday() string{
	today := time.Now()
	day := today.Format("2000-03-02")

	return day
}