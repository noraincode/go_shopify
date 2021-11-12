package go_shopify

import (
	"bytes"
	"os"
	"testing"
)

func TestLeveledLogger(t *testing.T) {
	tests := []struct {
		level  int
		input  string
		stdout string
		stderr string
	}{
		{
			level:  LevelError,
			input:  "log",
			stderr: "[ERROR] error log\n",
			stdout: "",
		},
		{
			level:  LevelWarn,
			input:  "log",
			stderr: "[ERROR] error log\n[WARN] warn log\n",
			stdout: "",
		},
		{
			level:  LevelInfo,
			input:  "log",
			stderr: "[ERROR] error log\n[WARN] warn log\n",
			stdout: "[INFO] info log\n",
		},
		{
			level:  LevelDebug,
			input:  "log",
			stderr: "[ERROR] error log\n[WARN] warn log\n",
			stdout: "[INFO] info log\n[DEBUG] debug log\n",
		},
	}

	for _, test := range tests {
		err := &bytes.Buffer{}
		out := &bytes.Buffer{}
		log := &LeveledLogger{Level: test.level, stderrOverride: err, stdoutOverride: out}

		log.Errorf("error %s", test.input)
		log.Warnf("warn %s", test.input)
		log.Infof("info %s", test.input)
		log.Debugf("debug %s", test.input)

		stdout := out.String()
		stderr := err.String()

		if stdout != test.stdout {
			t.Errorf("leveled logger %d expected stdout \"%s\" received \"%s\"", test.level, test.stdout, stdout)
		}
		if stderr != test.stderr {
			t.Errorf("leveled logger %d expected stderr \"%s\" received \"%s\"", test.level, test.stderr, stderr)
		}
	}

	log := &LeveledLogger{Level: LevelDebug}
	if log.stderr() != os.Stderr {
		t.Errorf("leveled logger with no stderr override expects os.Stderr")
	}
	if log.stdout() != os.Stdout {
		t.Errorf("leveled logger with no stdout override expects os.Stdout")
	}

}
