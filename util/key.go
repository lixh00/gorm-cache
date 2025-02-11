package util

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// SetInstanceId
// @description: 设置实例ID
// @param id
func SetInstanceId(id string) {
	InstanceId = id
}

// GetInstanceId
// @description: 获取实例ID
// @return string
func GetInstanceId() string {
	if InstanceId != "" {
		return InstanceId
	}
	charList := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	length := 5
	str := make([]byte, 0)
	for i := 0; i < length; i++ {
		str = append(str, charList[rand.Intn(len(charList))])
	}
	InstanceId = string(str)
	return InstanceId
}

func GenPrimaryCacheKey(instanceId string, tableName string, primaryKey string) string {
	return fmt.Sprintf("%s:%s:p:%s:%s", GormCachePrefix, instanceId, tableName, primaryKey)
}

func GenPrimaryCachePrefix(instanceId string, tableName string) string {
	return GormCachePrefix + ":" + instanceId + ":p:" + tableName
}

func GenSearchCacheKey(instanceId string, tableName string, sql string, vars ...interface{}) string {
	buf := strings.Builder{}
	buf.WriteString(sql)
	for _, v := range vars {
		pv := reflect.ValueOf(v)
		if pv.Kind() == reflect.Ptr {
			buf.WriteString(fmt.Sprintf(":%v", pv.Elem()))
		} else {
			buf.WriteString(fmt.Sprintf(":%v", v))
		}
	}
	return fmt.Sprintf("%s:%s:s:%s:%s", GormCachePrefix, instanceId, tableName, buf.String())
}

func GenSearchCachePrefix(instanceId string, tableName string) string {
	return GormCachePrefix + ":" + instanceId + ":s:" + tableName
}

func GenSingleFlightKey(tableName string, sql string, vars ...interface{}) string {
	buf := strings.Builder{}
	buf.WriteString(sql)
	for _, v := range vars {
		pv := reflect.ValueOf(v)
		if pv.Kind() == reflect.Ptr {
			buf.WriteString(fmt.Sprintf(":%v", pv.Elem()))
		} else {
			buf.WriteString(fmt.Sprintf(":%v", v))
		}
	}
	return fmt.Sprintf("%s:%s", tableName, buf.String())
}
