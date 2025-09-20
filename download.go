package govotcli

import (
	"os"
	"os/exec"
)

func Download(url string, path string, filename string) error {
	cmd := exec.Command("vot-cli", "--output="+path, "--output-file="+filename, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
