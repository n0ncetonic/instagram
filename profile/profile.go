package profile

// Data is a wrapper for profile data
type Data struct {
	SharedData
}

// GetConfig gets a Config from profile data
//
// Returns
//
// Config - Config struct
func (d Data) GetConfig() Config {
	return d.Config
}

// GetUser gets a User from profile data
//
// Returns
//
// User - User struct
func (d Data) GetUser() User {
	return d.EntryData.ProfilePage[0].GraphQL.User
}
