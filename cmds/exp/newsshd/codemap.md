# cmds/exp/newsshd/

## Responsibility
SSH server that accepts connections and spawns an interactive shell with PTY support.

## Integration Points
- github.com/gliderlabs/ssh: used for SSH server implementation
- github.com/creack/pty: used for PTY allocation for interactive shells
- pkg/uroot/unixflag: used for flag parsing
