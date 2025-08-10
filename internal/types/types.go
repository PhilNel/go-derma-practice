package types

type Special struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
	CurrentPrice     int    `json:"current_price"`
	OriginalPrice    int    `json:"original_price"`
	Currency         string `json:"currency"`
	Status           string `json:"status"`
	ActiveFrom       string `json:"active_from"`
	ActiveUntil      string `json:"active_until"`
	ImageKey         string `json:"image_key"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type SpecialsResponse struct {
	Specials []Special `json:"specials"`
}
