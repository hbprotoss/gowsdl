package example

type HelloService interface {
	HelloMulti(message string, entity *Entity, ids []int32, list []int32) (err error)
	HelloList(messages []string) (entity []*Entity, err error)
	Hello(message string) (entity *Entity, err error)
	
}