package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var binaryPath string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "img2char-test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	binaryPath = filepath.Join(dir, "img2char")
	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic("failed to build binary: " + err.Error())
	}

	os.Exit(m.Run())
}

func TestCLI_NoArgsExitsWithUsage(t *testing.T) {
	cmd := exec.Command(binaryPath)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected non-zero exit code with no arguments")
	}
	if !strings.Contains(string(out), "Usage:") {
		t.Errorf("output = %q, want to contain 'Usage:'", string(out))
	}
}

func TestCLI_640x200OutputMatchesExpected(t *testing.T) {
	expected, err := os.ReadFile("testimages/sample640x200.txt")
	if err != nil {
		t.Fatalf("read expected output: %v", err)
	}

	cmd := exec.Command(binaryPath, "testimages/sample640x200.png")
	got, err := cmd.Output()
	if err != nil {
		t.Fatalf("run img2char: %v", err)
	}

	if string(got) != string(expected) {
		t.Errorf("640x200 output mismatch:\n--- got ---\n%s\n--- want ---\n%s", string(got), string(expected))
	}
}

func TestCLI_320x200OutputMatchesExpected(t *testing.T) {
	expected, err := os.ReadFile("testimages/sample320x200.txt")
	if err != nil {
		t.Fatalf("read expected output: %v", err)
	}

	cmd := exec.Command(binaryPath, "testimages/sample320x200.png")
	got, err := cmd.Output()
	if err != nil {
		t.Fatalf("run img2char: %v", err)
	}

	if string(got) != string(expected) {
		t.Errorf("320x200 output mismatch:\n--- got ---\n%s\n--- want ---\n%s", string(got), string(expected))
	}
}

func TestCLI_NonexistentFileExitsWithError(t *testing.T) {
	cmd := exec.Command(binaryPath, "nonexistent.png")
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected non-zero exit code for nonexistent file")
	}
}

func TestCLI_VerboseFlagPrintsBlocksToStderr(t *testing.T) {
	cmd := exec.Command(binaryPath, "-v", "testimages/sample640x200.png")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("stderr pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start: %v", err)
	}

	buf := make([]byte, 4096)
	n, _ := stderr.Read(buf)
	stderrOutput := string(buf[:n])

	if err := cmd.Wait(); err != nil {
		t.Fatalf("wait: %v", err)
	}

	if !strings.Contains(stderrOutput, "blocks:") {
		t.Errorf("stderr = %q, want to contain 'blocks:'", stderrOutput)
	}
}
