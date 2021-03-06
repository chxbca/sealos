package install

import (
	"github.com/wonderivan/logger"
	"golang.org/x/crypto/ssh"
	"os"
)

//CheckValid is
func (s *SealosInstaller) CheckValid() {
	hosts := append(Masters, Nodes...)
	if len(Masters) == 0{
		s.Print("Fail")
		logger.Error("master not allow empty")
		os.Exit(1)
	}
	if User == ""{
		s.Print("Fail")
		logger.Error("user not allow empty")
		os.Exit(1)
	}
	var session *ssh.Session
	var errors []error
	for _, h := range hosts {
		session, err := Connect(User, Passwd, PrivateKeyFile, h)
		if err != nil {
			logger.Error("[%s] ------------ check error", h)
			logger.Error("[%s] ------------ error[%s]", h, err)
			errors = append(errors, err)
		} else {
			logger.Crit("[%s]  ------------ check ok", h)
			logger.Crit("[%s]  ------------ session[%p]", h, session)
		}
	}
	defer func() {
		if session != nil {
			session.Close()
		}
	}()
	if len(errors) > 0 {
		logger.Error("has some linux server is connection ssh is failed")
		os.Exit(1)
	}
}
