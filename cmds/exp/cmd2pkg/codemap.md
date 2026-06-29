# cmds/exp/cmd2pkg/

## Responsibility
Converts one or more u-root commands into standalone Go packages (for use with gobusybox).

## Integration Points
- github.com/u-root/gobusybox/src/pkg/golang: used for Go environment configuration
- pkg/ulog: used for logging
