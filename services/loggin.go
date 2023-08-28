package services

import (
	"context"
	"fmt"
	"time"
)

type logginService struct {
	next Service
}

func NewLogginService(next Service) Service {
	return &logginService{

		next: next,
	}
}

func (cfs *logginService) GetCatFact(ctx context.Context)(fact *CatFact,err error){
    
     defer func (start time.Time)  {
		  fmt.Printf("fact=%v err=%s took=%v ",fact,err,time.Since(start))
	 }(time.Now())

     return cfs.next.GetCatFact(ctx)
}