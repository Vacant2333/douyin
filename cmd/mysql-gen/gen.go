package main

import (
	"douyin/pkg/sql/dal"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	DB := dal.ConnectDB()
	generateModel(DB)
}

func generateModel(DB *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "pkg/sql/dal/query",
		//  Default global struct code | Interface type query code
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		// If field is null , use ptr
		FieldNullable: false,
		//
		FieldCoverable: false,
		//
		FieldSignable: false,
		//
		FieldWithIndexTag: false,
		//
		FieldWithTypeTag: true,
	})

	g.UseDB(DB)

	//Map different integer type
	// FIXME (temporay like this)
	dataMap := map[string]func(detailType string) (dataType string){
		"tinyint":   func(detailType string) (dataType string) { return "uint" },
		"smallint":  func(detailType string) (dataType string) { return "uint" },
		"mediumint": func(detailType string) (dataType string) { return "uint" },
		"bigint":    func(detailType string) (dataType string) { return "uint" },
		"int":       func(detailType string) (dataType string) { return "uint" },
	}
	g.WithDataTypeMap(dataMap)

	// Json tag customize
	//make some tag  to stirng (long number)
	/*
	 jsonFields := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	 	toStringField := `balance, `
	 	if strings.Contains(toStringField, columnName) {
	 		return columnName + ",string"
	 	}
	 	return columnName
	 })
	*/

	// Rename defalt field (We don't need)
	// autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	//softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")

	// 模型自定义选项组
	// fieldOpts := []gen.ModelOpt{ /*jsonFields, autoCreateTimeField, autoUpdateTimeField, softDeleteField*/ }

	// Get the handle by create the model for relation operations.
	// Will be override by coming operations;
	UserModel := g.GenerateModel("user")
	ChatModel := g.GenerateModel("chat")
	FollowModel := g.GenerateModel("follow")
	CommentModel := g.GenerateModel("comment")
	VideoModel := g.GenerateModel("video")
	FavoriteModel := g.GenerateModel("favorite")

	// // 创建全部模型文件, 并覆盖前面创建的同名模型
	allModel := g.GenerateAllTable()

	// TODO: add relations
	// 创建有关联关系的模型文件
	// 可以用于指定外键
	// chatmodel := g.GenerateModel("chat",
	//	//  user 一对多 address 关联, 外键`uid`在 * 表中
	// 	gen.FieldNewTag("sender", "foreigney:sender"),
	// 	gen.FieldRelate(field.HasOne, "Sender", User, &field.RelateConfig{ GORMTag: "foreignKey:sender"}),
	// 	gen.FieldRelate(field.HasOne, "ID", User, &field.RelateConfig{GORMTag: "foreignKey:receiver"}),
	// )

	g.ApplyBasic(UserModel, ChatModel, FollowModel, CommentModel, VideoModel, FavoriteModel)
	g.ApplyBasic(allModel...)
	//g.ApplyBasic(chatmodel)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyInterface(func(dal.DefaultMethod) {}, UserModel, ChatModel, FollowModel, CommentModel, VideoModel, FavoriteModel)
	g.ApplyInterface(func(dal.UserMethod) {}, UserModel)
	g.ApplyInterface(func(dal.ChatMethod) {}, ChatModel)
	g.ApplyInterface(func(dal.FollowMethod) {}, FollowModel)
	g.ApplyInterface(func(dal.CommentMethod) {}, CommentModel)
	g.ApplyInterface(func(dal.VideoMethod) {}, VideoModel)
	g.ApplyInterface(func(dal.FavoriteMethod) {}, FavoriteModel)

	g.Execute()
}

// Config generator's basic configuration
/*
type Config struct {
	OutPath      string // query code path
	OutFile      string // query code file name, default: gen.go
	ModelPkgPath string // generated model code's package name
	WithUnitTest bool   // generate unit test for query code

	// generate model global configuration
	FieldNullable     bool // generate pointer when field is nullable
	FieldCoverable    bool // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
	FieldSignable     bool // detect integer field's unsigned type, adjust generated data type
	FieldWithIndexTag bool // generate with gorm index tag
	FieldWithTypeTag  bool // generate with gorm column type tag

	Mode GenerateMode // generate mode
	// contains filtered or unexported fields
}
*/
