package director

import (
	"fmt"
	"octopus/utils"
)

//DirectorStart
type DirectorStart struct {
}

func (d *DirectorStart) Start() {
	//profilesCfgPath + "/" + profilesCfgPrefix + "." + profilesCfgType
	bananerPath := ReadCfg("octopus", "director", "bananer", "url").(string)
	if bananerPath == "" {
		bananerPath = ProfilesCfgPath + "/" + ProfilesBananerFileName
	} else {
		bananerPath = ProfilesCfgPath + "/" + bananerPath
	}
	strs := utils.ReadFileLine(bananerPath)
	for _, str := range strs {
		fmt.Println(str)
	}
	//fmt.Println("DirectorStart start")
}
