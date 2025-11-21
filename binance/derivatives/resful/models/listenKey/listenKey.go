package listenkey

type Params struct{}

type PostResult struct {
	ListenKey string `json:"listenKey"`
}

type PutResult = PostResult

type DeleteResult struct{}
