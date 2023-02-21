package nciutil

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflakeId() string {
	snowflake.Epoch = 1672527600000
	node, _ := snowflake.NewNode(1)
	id := node.Generate()
	return id.String()
}
