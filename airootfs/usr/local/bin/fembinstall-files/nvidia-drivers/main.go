package main

import (
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("lspci").Output()
	if err != nil {
		return
	}

	outputStr := string(out)
	var cmd *exec.Cmd

	switch {
	case strings.Contains(outputStr, "GeForce RTX"):
		cmd = exec.Command("pacman", "-S", "--noconfirm", "nvidia", "nvidia-utils")

	case strings.Contains(outputStr, "AMD"):
		cmd = exec.Command("pacman", "-S", "--noconfirm", "mesa", "lib32-mesa", "xf86-video-amdgpu", "vulkan-radeon", "lib32-vulkan-radeon")

	case strings.Contains(outputStr, "Intel"):
		cmd = exec.Command("pacman", "-S", "--noconfirm", "mesa", "vulkan-intel", "intel-media-driver")
	
	default:
		return
	}

	_ = cmd.Run()
}