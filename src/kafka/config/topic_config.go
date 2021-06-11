package config

import "reflect"

type TopicConfig struct {
	Name                      string
	DistributedMemoryBehavior string
	Prefix                    string
	StructName                reflect.Type
	Command                   chan bool
}
