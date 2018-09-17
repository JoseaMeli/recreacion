package service

import (
	"github.com/IT-Accelerator-Tweet/src/domain"
	"testing"
)

func TestCrearVariosUsuarios(t *testing.T) {

	// Initialization
	usuarioManager := NewUsuarioManager()

	nombre := "Jose"
	mail := "jose@meli.com"
	nick := "Hyr0z"
	contraseña := "123"

	usuario1 := domain.NewUsuario(nombre, mail, nick, contraseña)

	nombre2 := "Kei"
	mail2 := "kei@meli.com"
	nick2 := "keinaka"
	contraseña2 := "1234"

	usuario2 := domain.NewUsuario(nombre2, mail2, nick2, contraseña2)

	// Operation
	usuarioManager.AgregarUsuario(usuario1)
	usuarioManager.AgregarUsuario(usuario2)

	// Validation
	usuarios := usuarioManager.GetUsuarios()

	if len(usuarios) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(usuarios))
		return
	}
}

func TestSoloPuedenTweetearLosUsuariosRegistrados(t *testing.T) {

	// Initialization
	usuarioManager := NewUsuarioManager()

	nombre := "Jose"
	mail := "jose@meli.com"
	nick := "Hyr0z"
	contraseña := "123"

	usuario1 := domain.NewUsuario(nombre, mail, nick, contraseña)

	usuarioManager.AgregarUsuario(usuario1)

	usuarioEncontrado := usuarioManager.GetUsuarioByMail(mail)

	if usuarioEncontrado == nil {
		t.Errorf("Expected find user")
	}
}