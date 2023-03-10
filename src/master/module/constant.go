package module

import "os"

var (
	ModuleDirectory = "./modules/"
)

func init() {
	if moduleDir, ok := os.LookupEnv("MODULE_DIR"); ok {
		ModuleDirectory = moduleDir
	}
}
