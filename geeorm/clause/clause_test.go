package clause

import (
	"geeorm/log"
	"testing"
)

func testSelect(t *testing.T) {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(SELECT, "user", []string{"*"})
	clause.Set(WHERE, "name = ? ", "sen")
	clause.Set(ORDERBY, "age asc")
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	log.Info(sql, vars)

}

func TestClause_build(t *testing.T) {
	testSelect(t)
}
