// +build remote

package config

// isDirectory tests whether the given path exists and is a directory. It
// follows symlinks.
func isDirectory(path string) error {
	return nil
}

func (c *EngineConfig) validatePaths() error {
	return nil
}

func (c *ContainersConfig) validateDevices() error {
	return nil
}

func (c *ContainersConfig) validateUlimits() error {
	return nil
}
