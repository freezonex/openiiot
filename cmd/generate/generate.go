// Code generated by hertztool.
package main

import (
	"context"
	//"fmt"
	"os"
	"strings"

	"gorm.io/gen"
	//"gorm.io/gen/field"
	//"gorm.io/gorm/logger"

	//"code.byted.org/gorm/bytedgen"
	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
)

// https://github.com/go-gorm/gen#query

type JsonContains interface {
	// where("JSON_CONTAINS(@@tags, @tag)")
	JSONContains(tags string, tag string) (gen.T, error)
}

var db *mysql.MySQL

// generated code
func main() {

	//init mysql
	c, err := config.Init()
	if err != nil {
		panic(err)
	}

	db, err = mysql.Init(&c.DBConfig)
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]
	if strings.ToLower(args[0]) == "openiiot" {
		generateDbOpeniiotModels()
	}

}

func generateDbOpeniiotModels() {
	db := mysql.DBOpeniiot(context.Background())

	g := gen.NewGenerator(gen.Config{
		OutPath:           "biz/dal/query_openiiot",
		ModelPkgPath:      "biz/dal/model_openiiot",
		FieldWithIndexTag: true,
		FieldNullable:     true,
	})

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("tenant"),
		g.GenerateModel("user"),
		g.GenerateModel("edge"),
		g.GenerateModel("core"),
		g.GenerateModel("app"),
		g.GenerateModel("flow"),
		g.GenerateModel("flow_edge"),
		g.GenerateModel("flow_core"),
		g.GenerateModel("flow_app"),
		g.GenerateModel("global_config"),
		g.GenerateModel("wms_warehouse"),
		g.GenerateModel("wms_storage_location"),
		g.GenerateModel("wms_material"),
		g.GenerateModel("wms_inbound"),
		g.GenerateModel("wms_inbound_record"),
		g.GenerateModel("wms_outbound"),
		g.GenerateModel("wms_outbound_record"),
		g.GenerateModel("wms_stocktaking"),
	)
	g.Execute()
}
