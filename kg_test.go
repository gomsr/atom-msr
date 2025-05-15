package atommsr

import (
	"fmt"
	"github.com/gomsr/atom-msr/kg"
	"testing"
)

// pk 值是零值就是灾难(如id=0)
func TestInit(t *testing.T) {
	Init("config_test.yaml")
	fmt.Printf("viper: %v", kg.V)
	fmt.Printf("zap: %v", kg.L)
	fmt.Printf("Config: %v", kg.C)
	fmt.Printf("Db: %v", kg.DB)

	//dal.SetDefault(kg.DB)
	//find, err := dal.Q.ArchivedUp.Preload(dal.Q.ArchivedUp.ArchivedUpsTag).Find()
	//if err != nil || find == nil {
	//	return
	//}
}
