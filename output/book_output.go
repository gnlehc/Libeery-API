package output

import "Libeery/model"

type ListBookOutput struct {
	Data       []*model.MsBook
	BaseOutput BaseOutput
}

type GetBookDetailsOutput struct {
	Data       *model.MsBook
	BaseOutput BaseOutput
}
