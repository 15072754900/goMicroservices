package jobtype

import "look-cp/app/order/model"

// DeferCloseHomestayOrderPayload defer close homestay order
type DeferCloseHomestayOrderPayload struct {
	Sn string
}

// PaySuccessNotifyUserPayLoad pay success notify user
type PaySuccessNotifyUserPayLoad struct {
	Order *model.HomestayOrder
}
