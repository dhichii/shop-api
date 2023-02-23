package response

type Toko struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}

type MyToko struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	IdUser   int    `json:"id_user"`
}
