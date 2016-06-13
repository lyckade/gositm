package conf

//Conf are all configuration properties, which should be able to be
//stored inside a separate configuration file
type Conf struct {
	Recursive    bool //Recursive walk for the backup
	BackupFolder string
	Ignore       []string //Ignore Patterns
}

//Default contains alle the default parameters
var Default = Conf{
	Recursive:    true,
	BackupFolder: "000_Backup",
	Ignore:       []string{".git", "000_Backup"},
}

//Properties are the configuration parameter for the other modules
var Properties Conf

func init() {
	Properties = Default
}
