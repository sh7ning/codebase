package dingtalk

import (
	"app/pkg/cfg"
	"app/pkg/utils/dingtalk_robot"
)

var dingTalkRobots = make(map[string]*dingtalk_robot.Robot)

func InitInstances() {
	for name, config := range cfg.AppConfig.DingTalk {
		dingTalkRobots[name] = dingtalk_robot.NewRobot(config.Token, config.Secret)
	}
}

func Instance(name string) *dingtalk_robot.Robot {
	if name == "" {
		name = "default"
	}

	return dingTalkRobots[name]
}
