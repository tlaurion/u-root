# pkg/fb/

## Responsibility

Provides Linux framebuffer initialization and image drawing. Supports
rendering `image.Image` onto a framebuffer byte buffer, and initializing
the framebuffer device (`/dev/fb0`) via ioctls.

## Design Patterns

- **Direct pixel manipulation**: draws RGBA image data into BGR framebuffer format
- **ioctl-based initialization**: queries screen info from the framebuffer device

## Integration Points

- `cmds/fb`: CLI framebuffer tools

## Key Types/Functions

- `DrawOnBufAt(buf []byte, img image.Image, posx, posy, stride, bpp int)` — Draw image onto framebuffer at position
- `FbInit() (width, height, bpp int, fb []byte, err error)` — Initialize framebuffer device
