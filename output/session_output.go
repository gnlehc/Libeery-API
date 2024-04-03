package output

import "Libeery/model"

type SessionOutput struct {
	Data       []*model.MsSession
	BaseOutput BaseOutput
}
