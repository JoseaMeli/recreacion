package domain

type Usuario struct {
	Nombre		string
	Mail		string
	Nick		string
	Contraseña	string
}

func NewUsuario(nombre string, mail string , nick string, contraseña string) *Usuario{

	u := Usuario {nombre, mail, nick, contraseña}
	return &u
}
