/*
	providers 提供全局依赖
*/

package providers

import "emall/components"

var DBconnector *components.DBConnector
var RedisConnector *components.RedisConnector

// 初始化各个providers
func Init() {

}
