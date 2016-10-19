package gormgen

import "text/template"

// Make sure that the template compiles during package initialization
func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
package {{.PkgName}}

{{if ne (.PkgName) "gormgen"}}
import "github.com/MohamedBassem/gormgen"
{{end}}
import "github.com/jinzhu/gorm"
import "fmt"

{{range .Structs}}

	func (t *{{.StructName}}) Save(db *gorm.DB) error {
		return db.Save(t).Error
	}

	func (t *{{.StructName}}) Delete(db *gorm.DB) error {
		return db.Delete(t).Error
	}

	type {{.QueryBuilderName}} struct {
		order []string
		where []struct{
			prefix string
			value interface{}
		}
		limit int
		offset int
	}

	func (qb *{{.QueryBuilderName}}) buildQuery(db *gorm.DB) *gorm.DB {
		ret := db
		for _, where := range qb.where {
			ret = ret.Where(where.prefix, where.value)
		}
		for _, order := range qb.order {
			ret = ret.Order(order)
		}
		ret = ret.Limit(qb.limit).Offset(qb.offset)
		return ret
	}

	func (qb *{{.QueryBuilderName}}) Count(db *gorm.DB) (int, error) {
		var c int
		res := qb.buildQuery(db).Model(&{{.StructName}}{}).Count(&c)
		if res.RecordNotFound() {
			c = 0
		}
		return c, res.Error
	}

	func (qb *{{.QueryBuilderName}}) First(db *gorm.DB) (*{{.StructName}}, error) {
		ret := &{{.StructName}}{}
		res := qb.buildQuery(db).First(ret)
		if res.RecordNotFound() {
			ret = nil
		}
		return ret, res.Error
	}

	func (qb *{{.QueryBuilderName}}) QueryOne(db *gorm.DB) (*{{.StructName}}, error) {
		qb.limit = 1
		ret, err := qb.QueryAll(db)
		if len(ret) > 0 {
			return &ret[0], err
		}else{
			return nil, err
		}
	}

	func (qb *{{.QueryBuilderName}}) QueryAll(db *gorm.DB) ([]{{.StructName}}, error) {
		ret := []{{.StructName}}{}
		err := qb.buildQuery(db).Find(&ret).Error
		return ret, err
	}

	func (qb *{{.QueryBuilderName}}) Limit(limit int) *{{.QueryBuilderName}} {
		qb.limit = limit
		return qb
	}

	func (qb *{{.QueryBuilderName}}) Offset(offset int) *{{.QueryBuilderName}} {
		qb.offset = offset
		return qb
	}

	{{$queryBuilderName := .QueryBuilderName}}
	{{range .Fields}}
		func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}(p {{if ne ($.PkgName) "gormgen"}}gormgen.{{end}}Predict, value {{.FieldType}}) *{{$queryBuilderName}} {
			 qb.where = append(qb.where, struct {
				prefix string
				value interface{}
			}{
				fmt.Sprintf("%v %v ?", "{{.ColumnName}}", p.String()),
				value,
			})
			return qb
		}

		func (qb *{{$queryBuilderName}}) OrderBy{{call $.Helpers.Titelize .FieldName}}(asc bool) *{{$queryBuilderName}} {
			order := "DESC"
			if asc {
				order = "ASC"
			}

			qb.order = append(qb.order, "{{.ColumnName}} " + order)
			return qb
		}
	{{end}}
{{end}}
`)
