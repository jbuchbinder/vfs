vfs for golang [![Build Status](https://drone.io/github.com/jbuchbinder/vfs/status.png)](https://drone.io/github.com/jbuchbinder/vfs/latest) [![Build Status](https://secure.travis-ci.org/jbuchbinder/vfs.png)](http://travis-ci.org/jbuchbinder/vfs) [![GoDoc](https://godoc.org/github.com/jbuchbinder/vfs?status.png)](https://godoc.org/github.com/jbuchbinder/vfs) [![Join the chat at https://gitter.im/jbuchbinder/vfs](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/jbuchbinder/vfs?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
======

vfs is library to support virtual filesystems. It provides basic abstractions of filesystems and implementations, like `OS` accessing the file system of the underlying OS and `memfs` a full filesystem in-memory.

This is a fork of [blang's original tree](https://github.com/blang/vfs), enhanced with:

 * golang.org/x/net/webdav FileSystem interface support
 * Database-backed filesystem support (in progress)

Usage
-----
```bash
$ go get github.com/jbuchbinder/vfs
```
Note: Always vendor your dependencies or fix on a specific version tag.

```go
import github.com/jbuchbinder/vfs
```

```go
// Create a vfs accessing the filesystem of the underlying OS
var osfs vfs.Filesystem = vfs.OS()
osfs.Mkdir("/tmp", 0777)

// Make the filesystem read-only:
osfs = vfs.ReadOnly(osfs) // Simply wrap filesystems to change its behaviour

// os.O_CREATE will fail and return vfs.ErrReadOnly
// os.O_RDWR is supported but Write(..) on the file is disabled
f, _ := osfs.OpenFile("/tmp/example.txt", os.O_RDWR, 0)

// Return vfs.ErrReadOnly
_, err := f.Write([]byte("Write on readonly fs?"))
if err != nil {
    fmt.Errorf("Filesystem is read only!\n")
}

// Create a fully writable filesystem in memory
mfs := memfs.Create()
mfs.Mkdir("/root", 0777)

// Create a vfs supporting mounts
// The root fs is accessing the filesystem of the underlying OS
fs := mountfs.Create(osfs)

// Mount a memfs inside /memfs
// /memfs may not exist
fs.Mount(mfs, "/memfs")

// This will create /testdir inside the memfs
fs.Mkdir("/memfs/testdir", 0777)

// This would create /tmp/testdir inside your OS fs
// But the rootfs `osfs` is read-only
fs.Mkdir("/tmp/testdir", 0777)
```

Check detailed examples below. Also check the [GoDocs](http://godoc.org/github.com/jbuchbinder/vfs).

Why should I use this lib?
-----

- Only Stdlib
- (Nearly) Fully tested (Coverage >90%)
- Easy to create your own filesystem
- Mock a full filesystem for testing (or use included `memfs`)
- Compose/Wrap Filesystems `ReadOnly(OS())` and write simple Wrappers
- Many features, see [GoDocs](http://godoc.org/github.com/jbuchbinder/vfs) and examples below

Features and Examples
-----

- [OS Filesystem support](http://godoc.org/github.com/jbuchbinder/vfs#example-OsFS)
- [ReadOnly Wrapper](http://godoc.org/github.com/jbuchbinder/vfs#example-RoFS)
- [DummyFS for quick mocking](http://godoc.org/github.com/jbuchbinder/vfs#example-DummyFS)
- [MemFS - full in-memory filesystem](http://godoc.org/github.com/jbuchbinder/vfs/memfs#example-MemFS)
- [MountFS - support mounts across filesystems](http://godoc.org/github.com/jbuchbinder/vfs/mountfs#example-MountFS)

Current state: ALPHA
-----

While the functionality is quite stable and heavily tested, interfaces are subject to change. 

    You need more/less abstraction? Let me know by creating a Issue, thank you.

Motivation
-----

I simply couldn't find any lib supporting this wide range of variation and adaptability.

Contribution
-----

Feel free to make a pull request. For bigger changes create a issue first to discuss about it.

License
-----

See [LICENSE](LICENSE) file.
