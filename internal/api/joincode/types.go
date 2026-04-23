package joincode

type CreateJoinCodeRequest struct {
	JoinCommand string `json:"joinCommand"`
	NodeRole    string `json:"nodeRole"`
}

type CreateJoinCodeResponse struct {
	JoinCode string `json:"joinCode"`
	NodeRole string `json:"nodeRole"`
}

type ResolveJoinCodeResponse struct {
	JoinCode    string `json:"joinCode"`
	JoinCommand string `json:"joinCommand"`
	NodeRole    string `json:"nodeRole"`
}
