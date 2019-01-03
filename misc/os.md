operating system
---

## signals

| signal      | number   | Comment                                 | key board |
|:------------|----------|-----------------------------------------|-----------|
| HUP         | 1        | Hungup or death of terminal             |           |
| INT         | 2        | Interrupt from keyboard                 |  ctrl+c   |
| KILL        | 9        | Kill Signal.                            |           |
| SIGPIPE     | 13       |                                         |           |
| TERM        | 15       | Termination signal. Default signal.     |           |
| CONT        | 18       | Contninue process if stopped            |           |
| STOP        | 19       | Stop process.                           |  ctrl+z   |


## boot sequence

- BIOS
  - initialie and load MBR
  - where is the MBR usually? It usually first 512 bytes in the head of disk.
  - during the BIOS startup, F12 or F13 can change the boot sequence.
- Boot loader (GRUB)
  - Load Kernel into memory.
  - Run the kernel via GRUB.
- Kernel
  - Mount filesystem
  - Initialize devices/drivers
  - Execute init process
- Init
  - Run level
  - Login prompt
  - udev, procfs/sysfs (sysinit)
