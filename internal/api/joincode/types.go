package joincode

import "time"

type CreateJoinCodeRequest struct {
	WorkerJoinCommand       string `json:"workerJoinCommand"`
	ControlPlaneJoinCommand string `json:"controlPlaneJoinCommand,omitempty"`
}

type CreateJoinCodeResponse struct {
	JoinCode  string    `json:"joinCode"`
	CreatedAt time.Time `json:"createdAt"`
}

type ResolveJoinCodeResponse struct {
	JoinCode    string    `json:"joinCode"`
	NodeRole    string    `json:"nodeRole"`
	JoinCommand string    `json:"joinCommand"`
	CreatedAt   time.Time `json:"createdAt"`
}
