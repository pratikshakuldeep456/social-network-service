package sns

type User struct {
	ID              int
	Name            string
	Email           string
	Password        string
	Profile         *Profile
	PendingRequests map[int]*User
	FriendList      map[int]*User
}

func NewUser(id int, name, email, pass string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: pass,
	}
}

func (u *User) UpdateProfile(ProfilePic,
	Bio string,
	Interests []string) {
	u.Profile = &Profile{
		ProfilePic: ProfilePic,
		Bio:        Bio,
		Interests:  Interests,
	}

}

func (u *User) SendRequest(id int) {

}
