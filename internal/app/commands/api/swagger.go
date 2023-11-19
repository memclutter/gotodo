//go:generate swag fmt .
//go:generate swag init  --parseDependency --parseDepth 1 --instanceName api -g swagger.go --outputTypes go,yaml
package api

//	@title				gotodo
//	@version			1.0
//	@description		API for gotodo
//
//	@host				localhost
//	@BasePath			/
//
//	@Tag.name			Tasks
//	@Tag.description	Endpoint tasks
