package define

type InitSetting struct {
	Python         string
	WorkIP         string
	ProgramName    string
	ProgramVersion string
}

var Setting InitSetting

func LoadINI() {
	/*
		file, err := ini.Load("./config.ini")
		Setting = InitSetting{
			Python: file.Section("tools").Key("python").String(),
		}

		if err != nil {
			log.Info("配置文件读取错误，请检查文件路径:", err)
		}
	*/

}
