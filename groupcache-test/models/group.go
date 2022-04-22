package models

type Group struct {
	Name string
}

func (g *Group) Get(key string) string {
	return g.Name
}
