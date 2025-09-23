package core

import "redisServer/internal/data_structure"

var dictStore *data_structure.Dict

func init() {
	dictStore = data_structure.CreateDict()
}
