package oldtemp

const cStrRowCountValueLenNotMatch = "Row count and actually count mismatch, expected: %d, actual: %d"

//GetFormatter return nil on error
func GetFormatter(expectFormat string) FormatterInf {
	switch expectFormat {
	case cStrFMTJSON:
		return jsonFmt{}
	default:
		return nil
	}
}

type FormatterInf interface {
	DoAction(rouwCount, newIndx int64, srcRows [][]string) ([]byte, error)
}
