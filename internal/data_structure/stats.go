package data_structure

type KeyspaceStat struct {	
	Key int64
	Expire int64
}

var HashKeySpaceStat KeyspaceStat = KeyspaceStat{
	Key: 0,
	Expire: 0,
}