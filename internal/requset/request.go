package requset

type InsertReq struct {
	Link string `json:"link"`
}

type FindByHashReq struct {
	Hash string `json:"hash"`
}
