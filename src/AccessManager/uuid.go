package accessmanager

import (
	"sync"
)

var (
	accessLevelMp = map[string]byte{}
	accessLevelMt sync.Mutex
	uuid_Mp_pin   = map[string]string{}
	uuid_Mt_pin   sync.Mutex
)

func SetUuid_pin(uuid, pin string) {
	uuid_Mt_pin.Lock()
	uuid_Mp_pin[uuid] = pin
	uuid_Mt_pin.Unlock()
}

func GetUuid_pin(uuid string) string {
	uuid_Mt_pin.Lock()
	pin := uuid_Mp_pin[uuid]
	uuid_Mt_pin.Unlock()
	return pin
}

func SetUuid_accessLevel(uuid string, aLevel byte) {
	accessLevelMt.Lock()
	accessLevelMp[uuid] = aLevel
	accessLevelMt.Unlock()
}

func GetUser_accessLevel(uuid string) byte {
	accessLevelMt.Lock()
	aLevel, ok := accessLevelMp[uuid]
	if !ok {
		aLevel = 255
	}
	accessLevelMt.Unlock()
	//
	return aLevel
}
