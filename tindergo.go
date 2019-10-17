package tindergo

type TinderGo struct {
	auth      TAuthResponseData
	requester *TRequest
}

func New() *TinderGo {
	return &TinderGo{}
}

func (t *TinderGo) Authenticate(token string) error {
	a := NewTAuth(token)
	auth, err := a.Send()
	if err != nil {
		return err
	}
	t.auth = auth

	r := NewTRequest(t.auth.APIToken)
	t.requester = r

	return nil
}

func (t *TinderGo) SetAPIToken(token string) {
	t.auth.APIToken = token

	r := NewTRequest(t.auth.APIToken)
	t.requester = r
}

func (t *TinderGo) APIToken() string {
	return t.auth.APIToken
}
