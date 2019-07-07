package service

import (
	"github.com/Sirupsen/logrus"
)

func (movieAPI *movieAPIServer) handleRedisErr() {
	for {
		select {
		case <-movieAPI.ctx.Done():
			return
		case redisWorkerVal := <-movieAPI.redisWorkerChan:
			err := redisWorkerVal.statusCMD.Err()
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (movieAPI *movieAPIServer) handleSQLErr() {
	for {
		select {
		case <-movieAPI.ctx.Done():
			return
		case sqlWorkerVal := <-movieAPI.sqlWorkerChan:
			if sqlWorkerVal.err != nil {
				logrus.Error(sqlWorkerVal.err)
			}
		}
	}
}
