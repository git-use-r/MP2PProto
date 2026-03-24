package operations

type Send struct {
	BaseOperation
}

func jdif() {
	var op Operation = Send{}

	println(op.Execute())
}
