package vim

import (
	"github.com/neovim/go-client/nvim"
)

var vim *nvim.Nvim

func Get() *nvim.Nvim {
	return vim
}

func SetVim(v *nvim.Nvim) {
	vim = v
}
