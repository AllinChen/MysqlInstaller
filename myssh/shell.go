package myssh

//Mv 移动命令
func (c Cli) Mv(OldPath, NewPath string) (res string, err error) {
	res, err = c.Run("mv -f " + OldPath + " " + NewPath)
	if err == nil {
		return res, nil
	}
	return "", err

}

//Cp 复制命令
func (c Cli) Cp(OldPath, NewPath string) (res string, err error) {
	res, err = c.Run("cp -rf " + OldPath + " " + NewPath)
	if err == nil {
		return res, nil
	}
	return "", err

}

//Rm 删除命令
func (c Cli) Rm(Path string) (res string, err error) {
	res, err = c.Run("Rm -rf " + Path)
	if err == nil {
		return res, nil
	}
	return "", err
}

//Mkdir 建立文件夹命令
func (c Cli) Mkdir(NewPath string) (res string, err error) {
	res, err = c.Run("Mkdir -p " + NewPath)
	if err == nil {
		return res, nil
	}
	return "", err
}
