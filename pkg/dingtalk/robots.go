package dingtalk

type Configs map[string]*Robot

type Robots struct {
	configs     Configs
	collections map[string]*Robot
}

var collections *Robots

func Init(configs Configs) {
	collections = New(configs)
}

func Get(conn string) *Robot {
	return collections.Get(conn)
}

func New(configs Configs) *Robots {
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
