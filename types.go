package twitterstream

type Hashtags struct {
    Indices                     []int
    Text                        string
}

type Urls struct {
    Url                         string
    Indices                     []int
    Display_url                 string
    Expanded_url                string
}

type User_mentions struct {
    Indices                     []int
    Screen_name                 string
    Id_str                      string
    Name                        string
    Id                          int64
}

type Media struct {
    Id                          int64
    Id_str                      string
    Media_url                   string
    Media_url_https             string
    Url                         string
    Display_url                 string
    Expanded_url                string
    Type                        string
    Indices                     []int
    Screen_name                 string
    Sizes                       []Sizes
}

type Sizes struct {
    large                       Dimensions
    medium                      Dimensions
    small                       Dimensions
    thumb                       Dimensions
}

type Dimensions struct {
    w                           int64
    resize                      string
    h                           int64
}

 type Entities struct {
    Hashtags                    []Hashtags
    Urls                        []Urls
    User_mentions               []User_mentions
    Media                       []Media
}

type User struct {
    Lang                               string
    Verified                           bool
    Followers_count                    int
    Location                           string
    Screen_name                        string
    Following                          bool
    Friends_count                      int
    Profile_background_color           string
    Favourites_count                   int
    Description                        string
    Notifications                      string
    Profile_text_color                 string
    Url                                string
    Time_zone                          string
    Statuses_count                     int
    Profile_link_color                 string
    Geo_enabled                        bool
    Profile_background_image_url       string
    Protected                          bool
    Contributors_enabled               bool
    Profile_sidebar_fill_color         string
    Name                               string
    Profile_background_tile            string
    Created_at                         string
    Profile_image_url                  string
    Id                                 int64
    Id_str                             string
    Utc_offset                         int
    Profile_sidebar_border_color       string
    Profile_image_url_https            string
    Profile_background_image_url_https string
}

type Tweet struct {
    Text                        string
    Truncated                   bool
    Geo                         string
    In_reply_to_screen_name     string
    Favorited                   bool
    Source                      string
    Contributors                string
    In_reply_to_status_id       string
    In_reply_to_user_id         int64
    Id                          int64
    Created_at                  string
    Entities                    Entities
    User                        User
}

type SiteStreamMessage struct {
    For_user                    int64
    Message                     Tweet
}

type Event struct {
    Target                      User
    Source                      User
    Created_at                  string
    Event                       string
}

type FriendList struct {
    Friends                     []int64
}
