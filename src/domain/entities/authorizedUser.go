package entities


type AuthorizedUser struct {
	Username string
	Roleid int
}






func NewAuthorizedUser(username string, roleid int) *AuthorizedUser {
	return &AuthorizedUser{
		Username: username,
		Roleid: roleid,
	}
}

func (user *AuthorizedUser) CreateClient(createClientRequest CreateClientRequest) (*Client, error) {

	err := createClientRequest.Validate()
	if err != nil {
		return nil, err
	}
	return &Client{
		Name: createClientRequest.Name,
		Birthdate: createClientRequest.Birthdate,
		Email: createClientRequest.Email,
		Mobilephone: createClientRequest.Mobilephone,
		Address: createClientRequest.Address,
		Diseases: createClientRequest.Diseases,
		OtherInfo: createClientRequest.OtherInfo,
	} , nil
}