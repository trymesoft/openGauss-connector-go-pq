package pq

import (
	"os/user"
	"strings"
)

func defaultSettings() map[string]string {
	settings := make(map[string]string)

	settings["host"] = defaultHost()
	settings["port"] = "5432"

	// Default to the OS user name. Purposely ignoring err getting user name from
	// OS. The client application will simply have to specify the user in that
	// case (which they typically will be doing anyway).
	user, err := user.Current()
	if err == nil {
		// Windows gives us the username here as `DOMAIN\user` or `LOCALPCNAME\user`,
		// but the libpq default is just the `user` portion, so we strip off the first part.
		username := user.Username
		if strings.Contains(username, "\\") {
			username = username[strings.LastIndex(username, "\\")+1:]
		}

		settings["user"] = username
	}

	settings["min_read_buffer_size"] = "8192"
	settings["application_name"] = "go-driver"
	return settings
}

// defaultHost attempts to mimic libpq's default host. libpq uses the default unix socket location on *nix and localhost
// on Windows. The default socket location is compiled into libpq. Since conn does not have access to that default it
// checks the existence of common locations.
func defaultHost() string {
	return "localhost"
}
