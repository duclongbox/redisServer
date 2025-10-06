package core

import "redisServer/internal/data_structure"

var dictStore *data_structure.Dict
var setStore map[string]*data_structure.SimpleSet
var cmsStore map[string]*data_structure.CMS
func init() {
	dictStore = data_structure.CreateDict()
	setStore = make(map[string]*data_structure.SimpleSet)
	cmsStore = make(map[string]*data_structure.CMS)
}
