package joincode

type CreateJoinCodeRequest struct {
	JoinCommand string `json:"joinCommand"`
}

type CreateJoinCodeResponse struct {
	JoinCode string `json:"joinCode"`
}

type ResolveJoinCodeResponse struct {
	JoinCode    string `json:"joinCode"`
	JoinCommand string `json:"joinCommand"`
}
