package go_mod_dependency_test

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/wazsmwazsm/mortar"
)

func Dependency() {
	fmt.Println("dependency : github.com/wazsmwazsm/mortar")
}
