package domain

import (
	"testing"
)

func TestCrearUnNuevoUsuarioDeberiaDevolverElUsuarioCreado(t *testing.T) {

	var usuario *Usuario

	nombre := "Jose"
	mail := "jose@meli.com"
	nick := "Hyr0z"
	contraseña := "123"

	usuario = NewUsuario(nombre, mail, nick, contraseña)

	isValidUsuario(t, usuario, nombre, mail, nick, contraseña)
}



func isValidUsuario(t *testing.T, usuario *Usuario, nombre string, mail string, nick string, contraseña string) bool {

	if usuario.Nombre != nombre {
		t.Errorf("Expected id is %v but was %v", nombre, usuario.Nombre)
	}

	if usuario.Mail != mail {
		t.Errorf("Expected id is %v but was %v", mail, usuario.Mail)
	}

	if usuario.Nick != nick {
		t.Errorf("Expected id is %v but was %v", nick, usuario.Nick)
	}

	if usuario.Contraseña != contraseña {
		t.Errorf("Expected id is %v but was %v", contraseña, usuario.Contraseña)
	}

	return true
}