package dingtalk

type Configs map[string]*Robot

type Robots struct {
	configs     Configs
	collections map[string]*Robot
}

func Init(configs Configs) *Robots {
	collections := make(map[string]*Robot)

	for name, config := range configs {
		collections[name] = &Robot{
			Token:  config.Token,
			Secret: config.Secret,
		}
	}
	return &Robots{
		configs:     configs,
		collections: collections,
	}
}

func (r *Robots) Get(conn string) *Robot {
	if conn == "" {
		conn = "default"
	}

	if obj, ok := r.collections[conn]; ok {
		return obj
	}

	return nil
}
