package expressions

type ExpressionStack struct {
	stack []interface{}
}

func (es *ExpressionStack) Push(element interface{}) {
	es.stack = append(es.stack, element)
}

func (es *ExpressionStack) Pop() interface{} {
	length := len(es.stack)
	if length > 0 {
		el := es.stack[length-1]
		es.stack = es.stack[:length-1]
		return el
	}
	return nil
}
