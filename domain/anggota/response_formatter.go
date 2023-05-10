package anggota

type RegisterFormatter struct {
	ID          int    `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NoHp        string `json:"no_hp"`
	Status      string `json:"status"`
	Token       string `json:"token"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

func FormatRegister(anggota Anggota, token string) RegisterFormatter {
	formatter := RegisterFormatter{
		ID:          anggota.ID,
		NamaLengkap: anggota.NamaLengkap,
		Username:    anggota.Username,
		Password:    anggota.Password,
		NoHp:        anggota.NoHp,
		Status:      anggota.Status,
		Token:       token,
	}
	return formatter
}

func LoginResponseFormatter(data Anggota, accToken string) LoginResponse {
	res := LoginResponse{
		Username:    data.Username,
		Password:    data.Password,
		AccessToken: accToken,
	}
	return res
}
