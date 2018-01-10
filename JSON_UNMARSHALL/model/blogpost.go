package model

//BlogPost que armazena dados de um post em um Blog
type BlogPost struct {
	//Essas tags JSOON devem refletir da mesma forma que está lá na API. (https://jsonplaceholder.typicode.com/)
	UsuarioID int    `json:"userid"`
	ID        int    `json:"id"`
	Titulo    string `json:"title"`
	Texto     string `json:"body"`
}
