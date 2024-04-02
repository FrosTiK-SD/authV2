package interfaces

type StudentsRollNoReq struct {
	RollNos []int `json:"rollNos" bson:"rollNos"`
}