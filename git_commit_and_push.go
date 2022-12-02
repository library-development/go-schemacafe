package schemacafe

import "os/exec"

func gitCommitAndPush(dir string, event *Event) error {
	gitpath, err := exec.LookPath("git")
	if err != nil {
		return err
	}
	cmd := exec.Command(gitpath, "add", ".")
	cmd.Dir = dir
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command(gitpath, "commit", "-m", event.CommitMessage())
	cmd.Dir = dir
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command(gitpath, "push")
	cmd.Dir = dir
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
