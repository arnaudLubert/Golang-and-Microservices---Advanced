package conf

type Configuration struct {
    Host            string      `yaml:"host"`
    Port            string      `yaml:"port"`
    Storage         string      `yaml:"storage"`
    Credentials     Credentials `yaml:"credentials"`
    UsersService    Service     `yaml:"users_service"`
}

type Credentials struct {
    ApiKey          string      `yaml:"api_key"`
}

type Service struct {
    Url             string      `yaml:"url"`
    ApiKey          string      `yaml:"api_key"`
}