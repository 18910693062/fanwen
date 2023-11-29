package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"go-edu/config"
	"go.uber.org/zap"
)

var jenkins *gojenkins.Jenkins

func JenkinsClient(host string, user string, pwd string) *gojenkins.Jenkins {
	ctx := context.Background()
	jenkins, err := gojenkins.CreateJenkins(nil, host, user, pwd).Init(ctx)
	if err != nil {
		zap.L().Error(err.Error())
		return nil
	}
	zap.L().Info("jenkins链接成功")
	return jenkins
}

//func GetJenkinsJobAll() any {
//	ctx := context.Background()
//	jobList, _ := jenkins.GetAllJobNames(ctx)
//	jobnameall := list.New()
//	for _, job := range jobList {
//		jobnameall.PushBack(job.Name)
//	}
//	return jobnameall
//}

func Init(cfg *config.JenkinsConfig) (err error) {
	jenkinsHost := cfg.Host
	password := cfg.Password
	user := cfg.User
	jenkins = JenkinsClient(jenkinsHost, user, password)
	return

}

func GetJobConfig(name string) string {
	ctx := context.Background()
	a, _ := jenkins.GetJob(ctx, name)
	getConfig, _ := a.GetConfig(ctx)
	return getConfig
}

func BuilJob(name string) int {
	ctx := context.Background()
	a, _ := jenkins.BuildJob(ctx, name, nil)
	return int(a)
}

func CreateJob(name string, jobconfig string) {
	ctx := context.Background()
	_, err := jenkins.CreateJob(ctx, jobconfig, name)
	if err != nil {
		return
	}
	return
}
