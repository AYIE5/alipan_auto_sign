package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cook"`
	KKCookie       string `yaml:"_UP_A4A_11_=wb96318138284d96b0a685a9d36e9d35; _UP_30C_6A_=st96362019cura7iji0nwmtlrsn45ctg; _UP_TS_=sg1707f4e0faf1e2ac0db8509956c3e192e; _UP_E37_B7_=sg1707f4e0faf1e2ac0db8509956c3e192e; _UP_TG_=st96362019cura7iji0nwmtlrsn45ctg; _UP_335_2B_=1; tfstk=eK3XAu0o-tXXU-IcFKOrOlEVXea6TxTeG1NttfQV6rUAWAGgUSS2mlJsVAHzHq-03YwsTfba3jzv1AMKhlkssIVTBAH_Qm-y8jc0SPpUCe8Ei75JHOvUjV4JzP49LL8m8Iq4SvwEDTseJOdHpihfiZbVSr4mb_-upJg7GuKr38poXqFbcJhx_NQ99SZbpje54AQUde15f_ZhW7weNQsGjw7M15HG5Ir4D7VyWQO5YSZYZ7weNQsGjoFuacdWNMPf.; CwsSessionId=2c10e2a9-9214-4613-b70d-2815adb2c3bd; _UP_D_=pc; __pus=bdf30cfec20e8871ebc1a5e109d99bc1AATBWyridojx0tq53LK/cN1R51eWZz0PfTOx50tJShaNpd3elYyS024V5XuEtCsfGpqF5QhFdYu3EAO+z+plmOnF; __kp=d2bbc000-061e-11ef-8cdf-d9dc1d9fb11d; __kps=AAQajWWVMbmc8tBRBPTUwyCM; __ktd=vnRtNyw1u57+OF3Ki2tmFA==; __uid=AAQajWWVMbmc8tBRBPTUwyCM; __puus=1744cadfb94dd8c402bf3555614a0533AASzVgLvOPJvUSXwfPl98ip/3a2UTNzIRLrng3ZzWscO8WsIzAe5DMQplc4Q9HbEPyrTGt+HKTUZN1jJ4E8QsK2F08xyG4Gi80B15lxONf1n2LiEAJVnzcB3WIcMWpm+5Fwrhg5Hs1p5/IMxFGKwJFQjE5DPc0WUk33n9p52yE+TlVgt3X/r00YwC1xLt32dTDhLzK0g3rGskj3B9GtT/sE6"`
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
