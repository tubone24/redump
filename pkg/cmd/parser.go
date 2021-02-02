package cmd

var DocOptConf struct {
	Migrate     bool `docopt:"migrate"`
	List        bool `docopt:"list"`
	Dump        bool `docopt:"dump"`
	Restore     bool `docopt:"restore"`
	Help        bool `docopt:"-h,--help"`
	Version     bool `docopt:"-v,--version"`
	Concurrency bool `docopt:"-c,--concurrency"`
	Issue       bool `docopt:"-i,--issue"`
	Number      int  `docopt:"<number>"`
}
