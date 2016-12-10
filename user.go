package docbase

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
}
