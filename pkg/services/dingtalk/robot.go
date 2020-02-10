package dingtalk

import (
	"app/pkg/cfg"
	"app/pkg/utils/dingtalk_robot"
)

var dingTalkRobots = make(map[string]*dingtalk_robot.Robot)

func Instance(name string) *dingtalk_robot.Robot {
	if name == "" {
		name = "default"
	}

	if _, ok := dingTalkRobots[name]; !ok {
		if config, ok := cfg.AppConfig.DingTalk[name]; ok {
			dingTalkRobots[name] = dingtalk_robot.NewRobot(config.Token, config.Secret)
		} else {
			return nil
		}
	}

	return dingTalkRobots[name]
}
