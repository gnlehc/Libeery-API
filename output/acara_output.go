package output

import "Libeery/model"

type GetAcaraOutput struct {
	Data       []*model.MsAcara
	BaseOutput BaseOutput
}

type GetAcaraDetailsOutput struct {
	Data       *model.MsAcara
	BaseOutput BaseOutput
}
