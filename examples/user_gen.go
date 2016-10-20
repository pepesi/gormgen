package example

///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

import (
	"fmt"
	"time"

	"github.com/MohamedBassem/gormgen"
	"github.com/jinzhu/gorm"
)

func (t *User) Save(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *User) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

type UserQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *UserQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *UserQueryBuilder) Count(db *gorm.DB) (int, error) {
	var c int
	res := qb.buildQuery(db).Model(&User{}).Count(&c)
	if res.RecordNotFound() {
		c = 0
	}
	return c, res.Error
}

func (qb *UserQueryBuilder) First(db *gorm.DB) (*User, error) {
	ret := &User{}
	res := qb.buildQuery(db).First(ret)
	if res.RecordNotFound() {
		ret = nil
	}
	return ret, res.Error
}

func (qb *UserQueryBuilder) QueryOne(db *gorm.DB) (*User, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return &ret[0], err
	} else {
		return nil, err
	}
}

func (qb *UserQueryBuilder) QueryAll(db *gorm.DB) ([]User, error) {
	ret := []User{}
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *UserQueryBuilder) Limit(limit int) *UserQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *UserQueryBuilder) Offset(offset int) *UserQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *UserQueryBuilder) WhereID(p gormgen.Predicate, value uint) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByID(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *UserQueryBuilder) WhereCreatedAt(p gormgen.Predicate, value time.Time) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByCreatedAt(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *UserQueryBuilder) WhereUpdatedAt(p gormgen.Predicate, value time.Time) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByUpdatedAt(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *UserQueryBuilder) WhereDeletedAt(p gormgen.Predicate, value *time.Time) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByDeletedAt(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "deleted_at "+order)
	return qb
}

func (qb *UserQueryBuilder) WhereName(p gormgen.Predicate, value string) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByName(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *UserQueryBuilder) WhereAge(p gormgen.Predicate, value int) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "age", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByAge(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "age "+order)
	return qb
}

func (qb *UserQueryBuilder) WhereEmail(p gormgen.Predicate, value string) *UserQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "email", p.String()),
		value,
	})
	return qb
}

func (qb *UserQueryBuilder) OrderByEmail(asc bool) *UserQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "email "+order)
	return qb
}