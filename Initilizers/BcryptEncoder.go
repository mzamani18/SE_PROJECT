package initilizers

import utils "TikOn/Utils"

var PasswordEncoder utils.PasswordEncoder

func NewBCryptPasswordEncoder() {
	PasswordEncoder = utils.NewDefaultBCryptPasswordEncoder()
}
