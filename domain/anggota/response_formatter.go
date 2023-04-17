package anggota

type loginResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func LoginResponseFormatter(data Anggota, accToken string) loginResponse {
	var res loginResponse

	res.Username = data.Username
	res.AccessToken = accToken

	return res
}
