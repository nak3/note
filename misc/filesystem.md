note related to storage/filesystem
---

## soft link and hard link

### Hardlink
- Hardlink shares same inode number.
- Hardlink is reachable after you deleted the original file.
- Hardlink exists even after renaming it.


### Softlink
- Softlink points to the place of original file.
- Softlink can be created for directory.
- Softlink can beyond particion.

## inode

inode is a data, which has:

- owner
- access mode
- size
- atime, mtime, ctime (*ctime is not creation time, but change time)
- acl
- block list of where the data is
