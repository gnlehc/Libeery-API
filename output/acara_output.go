package output

import "Libeery/model"

type GetAcaraOutput struct {
	Data       []*model.MsAcara
	BaseOutput BaseOutput
}
