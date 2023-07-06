package users

func validatePassword(password string) error {
	return nil

}

func (u User) validateForCreateNewInstance() error {
	if err := validatePassword(u.Password); err != nil {
		return err
	}
	return nil
}
