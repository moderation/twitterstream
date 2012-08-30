package twitterstream

type Hashtags struct {
    Indices [2]uint32
	Text    string
}

type Urls struct {
	Url          string
	Indices      [2]uint32
	Display_url  string
	Expanded_url string
}

type User_mentions struct {
	Indices     [2]uint32
	Screen_name string
	Id_str      string
	Name        string
	Id          uint64
}

type Media struct {
	Id              uint64
	Id_str          string
	Media_url       string
	Media_url_https string
	Url             string
	Display_url     string
	Expanded_url    string
	Type            string
	Indices         [2]uint32
	Screen_name     string
	Sizes           []Sizes
}

type Sizes struct {
	large  map[string]Dimensions
	medium map[string]Dimensions
	small  map[string]Dimensions
	thumb  map[string]Dimensions
}

type Dimensions struct {
	w      uint32
	resize string
	h      uint32
}

type Entities struct {
	Hashtags      []Hashtags
	Urls          []Urls
	User_mentions []User_mentions
	Media         []Media
}

type Place struct {
	Name         string
	Url          string
	Country_code string
	Id           string
	Display_url  string
	Place_type   string
	Country      string
	Full_name    string
	Attributes   map[string]string
	Bounding_box Box
}

type Geo struct {
	Type        string
	Coordinates []float32
}

type Box struct {
	Type        string
	Coordinates [][][]float32
}

type User struct {
	Contributors_enabled               bool
	Created_at                         string
	Default_profile_image              bool
	Default_profile                    bool
	Description                        string
	Favourites_count                   uint32
	Followers_count                    uint32
	Follow_request_sent                *bool
	Following                          *bool
	Friends_count                      uint32
	Geo_enabled                        bool
	Id                                 uint64
	Id_str                             string
	Is_translator                      bool
	Lang                               string
	Listed_count                       uint32
	Location                           string
	Name                               string
	Notifications                      *bool
	Profile_background_color           string
	Profile_background_image_url       string
	Profile_background_image_url_https string
	Profile_background_tile            bool
	Profile_image_url                  string
	Profile_image_url_https            string
	Profile_link_color                 string
	Profile_sidebar_border_color       string
	Profile_sidebar_fill_color         string
	Profile_text_color                 string
	Profile_use_background_image       bool
	Protected                          bool
	Screen_name                        string
	Show_all_inline_media              bool
	Statuses_count                     uint32
	Time_zone                          string
	Url                                string
	Utc_offset                         int16
	Verified                           bool
}

type Tweet struct {
	Contributors              *string
	Created_at                string
	Favorited                 bool
	Id                        uint64
	Id_str                    string
	In_reply_to_screen_name   *string
	In_reply_to_status_id     *uint64
	In_reply_to_status_id_str *string
	In_reply_to_user_id       *uint64
	In_reply_to_user_id_str   *string
	Possibly_sensitive        bool
	Retweeted                 bool
	Retweeted_count           uint32
	Retweeted_status          Retweet
	Source                    string
	Text                      string
	Truncated                 bool
	Coordinates               Geo
	Entities                  Entities
	User                      User
	Place                     Place
	Geo                       Geo
}

type Metadata struct {
	Result_type string
}

type Retweet struct {
	Retweeted_count uint32
	Created_at      string
	Retweeted       bool
	Truncated       bool
	Coordinates     Geo
	Geo             Geo
	User            User
}

type Results struct {
	Created_at        		string
	From_user         		string
	From_user_id      		uint64
	From_user_id_name 		string
	From_user_id_str  		string
	Id                		uint64
	Id_str            		string
	Iso_language_code 		string
	Profile_image_url 		string
	Profile_image_url_https string
	Source            		string
	Text              		string
	To_user           		*string
	To_user_id        		uint64
	To_user_id_s      		*string
	To_user_name      		*string
	Metadata          		Metadata
	Entities          		Entities
}

type Query struct {
	Completed_in     float32
	Max_id           uint64
	Max_id_str       string
	Next_page        string
	Page             uint8
	Previous_page    string
	Query            string
	Refresh_url      string
	Results_per_page uint8
	Since_id         uint8
	Since_id_str     string
	Results          []Results
}
