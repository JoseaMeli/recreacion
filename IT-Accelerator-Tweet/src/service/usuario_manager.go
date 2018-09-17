package service

import "github.com/IT-Accelerator-Tweet/src/domain"

type UsuarioManager struct {
	Usuarios []*domain.Usuario
}

func NewUsuarioManager() *UsuarioManager{
	um := UsuarioManager{}
	um.Usuarios = make([]*domain.Usuario, 0)

	return &um
}

func (um *UsuarioManager) AgregarUsuario(usuario *domain.Usuario) {
	um.Usuarios = append(um.Usuarios, usuario)
}

func (um *UsuarioManager) GetUsuarios() []*domain.Usuario {
	return um.Usuarios
}

func (um *UsuarioManager)GetUsuarioByMail(mail string)  *domain.Usuario{
	for _,usuario := range um.Usuarios {
		if usuario.Mail == mail {
			return usuario
		}
	}
	return nil
}

func (um *UsuarioManager)GetUsuarioByNombre(nombre string)  *domain.Usuario{
	for _,usuario := range um.Usuarios {
		if usuario.Nombre == nombre {
			return usuario
		}
	}
	return nil
}


