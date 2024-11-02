package oscap

import (
	"log"
	"os/exec"

	"github.com/marcusburghardt/openscap-prototype/config"
)

func ValidateProfile(profile string) {

}

func constructScanCommand(openscapFiles map[string]string, profile string) ([]string, error) {
	profileName, err := config.SanitizeInput(profile)
	if err != nil {
		return nil, err
	}

	datastream := openscapFiles["datastream"]
	tailoringFile := openscapFiles["policy"]
	resultsFile := openscapFiles["results"]
	arfFile := openscapFiles["arf"]

	cmd := []string{
		"oscap",
		"xccdf",
		"eval",
		"--profile",
		profileName,
		"--results",
		resultsFile,
		"--results-arf",
		arfFile,
	}

	if tailoringFile != "" {
		cmd = append(cmd, "--tailoring-file", tailoringFile)
	}
	cmd = append(cmd, datastream)
	return cmd, nil
}

func OscapScan(openscapFiles map[string]string, profile string) ([]byte, error) {
	command, err := constructScanCommand(openscapFiles, profile)
	if err != nil {
		return nil, err
	}

	log.Printf("Executing the command: '%v'", command)
	cmd := exec.Command(command[0], command[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
