package commandjsonio

import (
	"bytes"
	"context"
	"encoding/json"
	"time"
)

// DefaultRunCommandJSONTimeout define default timeout duration for RunCommandJSON.
const DefaultRunCommandJSONTimeout = time.Second * 10

// RunCommandJSON run given command and send JSON encoded inputRef into STDIN of command.
// Output of STDOUT will be decode as JSON and put into outputRef.
func RunCommandJSON(ctx context.Context, cmdArgs, cmdEnvs []string, workDir string, timeoutDuration time.Duration, inputRef, outputRef interface{}) (err error) {
	b, err := json.Marshal(inputRef)
	if nil != err {
		err = &ErrEncodeInput{
			err: err,
		}
		return
	}
	stdinReader := bytes.NewReader(b)
	if ctx == nil {
		ctx = context.Background()
	}
	if timeoutDuration <= 0 {
		timeoutDuration = DefaultRunCommandJSONTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, timeoutDuration)
	defer cancel()
	cmdInst := setupExecCmd(ctx, cmdArgs, cmdEnvs, workDir)
	cmdInst.Stdin = stdinReader
	stdoutPipe, err := cmdInst.StdoutPipe()
	if err != nil {
		return
	}
	if err = cmdInst.Start(); err != nil {
		return
	}
	var err1, err2 error
	err1 = json.NewDecoder(stdoutPipe).Decode(outputRef)
	err2 = cmdInst.Wait()
	if nil != err1 {
		err = &ErrDecodeOutput{
			err: err1,
		}
		return
	}
	if nil != err2 {
		err = err2
	}
	return
}
