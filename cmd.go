package commandjsonio

import (
	"context"
	"encoding/json"
	"io"
	"os/exec"
)

func setupExecCmd(ctx context.Context, cmdArgs, cmdEnvs []string, workDir string) (cmd *exec.Cmd) {
	cmd = &exec.Cmd{
		Path: cmdArgs[0],
		Args: cmdArgs,
		Dir:  workDir,
	}
	if len(cmdEnvs) > 0 {
		cmd.Env = cmdEnvs
	}
	return
}

// CmdJSONReaderWriter provide command with JSON I/O support methods.
type CmdJSONReaderWriter struct {
	Cmd *exec.Cmd

	stdinPipe  io.WriteCloser
	stdoutPipe io.ReadCloser
	stdinEnc   *json.Encoder
	stdoutDec  *json.Decoder
}

// NewCommandJSONReaderWriter create new instance of CmdJSONReaderWriter.
func NewCommandJSONReaderWriter(ctx context.Context, cmdArgs, cmdEnvs []string, workDir string) (cmd *CmdJSONReaderWriter) {
	cmdInst := setupExecCmd(ctx, cmdArgs, cmdEnvs, workDir)
	cmd = &CmdJSONReaderWriter{
		Cmd: cmdInst,
	}
	return
}

// Start starts the specified command and setup JSON reader and writer.
func (cmd *CmdJSONReaderWriter) Start() (err error) {
	stdinPipe, err := cmd.Cmd.StdinPipe()
	if nil != err {
		return
	}
	stdoutPipe, err := cmd.Cmd.StdoutPipe()
	if nil != err {
		stdinPipe.Close()
		return
	}
	if err = cmd.Cmd.Start(); nil != err {
		stdinPipe.Close()
		stdoutPipe.Close()
		return
	}
	cmd.stdinPipe = stdinPipe
	cmd.stdoutPipe = stdoutPipe
	cmd.stdinEnc = json.NewEncoder(stdinPipe)
	cmd.stdoutDec = json.NewDecoder(stdoutPipe)
	return
}

// Write convert given data v into JSON and send it into STDIN pipe.
func (cmd *CmdJSONReaderWriter) Write(v interface{}) (err error) {
	return cmd.stdinEnc.Encode(v)
}

// Read load JSON from STDOUT and decode it into given reference v.
func (cmd *CmdJSONReaderWriter) Read(v interface{}) (err error) {
	return cmd.stdoutDec.Decode(v)
}

// CloseStdin close stdin pipe of command.
func (cmd *CmdJSONReaderWriter) CloseStdin() (err error) {
	cmd.stdinEnc = nil
	err = cmd.stdinPipe.Close()
	cmd.stdinPipe = nil
	return
}

// Wait call .Wait() of running command.
func (cmd *CmdJSONReaderWriter) Wait() (err error) {
	err = cmd.Cmd.Wait()
	cmd.stdinEnc = nil
	cmd.stdoutDec = nil
	cmd.stdinPipe = nil
	cmd.stdoutPipe = nil
	return
}
