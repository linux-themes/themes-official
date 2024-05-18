package themes

type Package struct {
	name      string
	link      string
	pack_type string
	selected  bool
}

var packages_offical = map[int]Package{
	1: {"mint", "https://", "icons", false},
	2: {"marble-shell", "https://github.com/linux-themes/themes/raw/main/themes/marble-shell.tar.gz", "themes", false},
}
