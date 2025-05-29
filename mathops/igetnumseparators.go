package mathops

type IGetNumSeparators interface {
	GetOutputSeparators() (*NumericSeparatorDto, error)

	GetInputSeparators() (*NumericSeparatorDto, error)
}
