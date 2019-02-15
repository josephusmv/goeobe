package eobedb

//QueryDefn Definition of the DB query
// IMPORTANT !! TODO!!! this structure should not be loaded once, it should be loaded during query!!!!
type QueryDefn struct {
	QueryActionName  string   `yaml:"QueryActionName"` //Key to find this Query Define
	TableName        string   `yaml:"TableName"`
	QueryType        string   `yaml:"QueryType"`        //SELECT, UPDATE, INSERT, DELETE
	WhereReadyStr    string   `yaml:"WhereReadyStr"`    // For WHERE Where ready string, if provided will not need for combine Where inside.
	ParamterColNames []string `yaml:"ParamterColNames"` // For WHERE parameters
	ParametrLogicOp  []string `yaml:"ParametrLogicOp"`  // For WHERE parameters, support only AND, OR
	ParamterOper     []string `yaml:"ParamterOper"`     // For WHERE parameters, support only eq, neq, gt, get, lt, let, like
	ExpectedColNames []string `yaml:"ExpectedColNames"` // For UPDATE, INSERT to set values, For SELECT indicates the query columns
}

func (qry *QueryDefn) paramterOperStr(index int) string {
	switch qry.ParamterOper[index] {
	case "neq":
		return "!="
	case "like":
		return " like "
	case "gt":
		return ">"
	case "lt":
		return "<"
	case "get":
		return ">="
	case "let":
		return "<="
	case "eq":
		return "="
	default:
		return "="

	}
}

//QueryData Carry the query data for query
type QueryData struct {
	QryActionDfn    QueryDefn
	ParameterValues []string //Carry Values for WHERE statements.
	ExpectedValues  []string //Carry Values for INSERT/UPDATE VALUES/SET statment.
}

//QueryData Carry the query data for query
type QueryResult struct {
	QueryActionName string
	QueryRows       [][]string //Rows get from DB, strictly same order of QueryColumns
	AffectedRows    int64      //Tottal selected row counts for Query, affected rows for exec
	LastIndex       int64
	QueryErr        error
}
