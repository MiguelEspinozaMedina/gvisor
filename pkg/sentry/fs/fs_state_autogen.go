// automatically generated by stateify.

package fs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *StableAttr) beforeSave() {}
func (x *StableAttr) save(m state.Map) {
	x.beforeSave()
	m.Save("Type", &x.Type)
	m.Save("DeviceID", &x.DeviceID)
	m.Save("InodeID", &x.InodeID)
	m.Save("BlockSize", &x.BlockSize)
	m.Save("DeviceFileMajor", &x.DeviceFileMajor)
	m.Save("DeviceFileMinor", &x.DeviceFileMinor)
}

func (x *StableAttr) afterLoad() {}
func (x *StableAttr) load(m state.Map) {
	m.Load("Type", &x.Type)
	m.Load("DeviceID", &x.DeviceID)
	m.Load("InodeID", &x.InodeID)
	m.Load("BlockSize", &x.BlockSize)
	m.Load("DeviceFileMajor", &x.DeviceFileMajor)
	m.Load("DeviceFileMinor", &x.DeviceFileMinor)
}

func (x *UnstableAttr) beforeSave() {}
func (x *UnstableAttr) save(m state.Map) {
	x.beforeSave()
	m.Save("Size", &x.Size)
	m.Save("Usage", &x.Usage)
	m.Save("Perms", &x.Perms)
	m.Save("Owner", &x.Owner)
	m.Save("AccessTime", &x.AccessTime)
	m.Save("ModificationTime", &x.ModificationTime)
	m.Save("StatusChangeTime", &x.StatusChangeTime)
	m.Save("Links", &x.Links)
}

func (x *UnstableAttr) afterLoad() {}
func (x *UnstableAttr) load(m state.Map) {
	m.Load("Size", &x.Size)
	m.Load("Usage", &x.Usage)
	m.Load("Perms", &x.Perms)
	m.Load("Owner", &x.Owner)
	m.Load("AccessTime", &x.AccessTime)
	m.Load("ModificationTime", &x.ModificationTime)
	m.Load("StatusChangeTime", &x.StatusChangeTime)
	m.Load("Links", &x.Links)
}

func (x *AttrMask) beforeSave() {}
func (x *AttrMask) save(m state.Map) {
	x.beforeSave()
	m.Save("Type", &x.Type)
	m.Save("DeviceID", &x.DeviceID)
	m.Save("InodeID", &x.InodeID)
	m.Save("BlockSize", &x.BlockSize)
	m.Save("Size", &x.Size)
	m.Save("Usage", &x.Usage)
	m.Save("Perms", &x.Perms)
	m.Save("UID", &x.UID)
	m.Save("GID", &x.GID)
	m.Save("AccessTime", &x.AccessTime)
	m.Save("ModificationTime", &x.ModificationTime)
	m.Save("StatusChangeTime", &x.StatusChangeTime)
	m.Save("Links", &x.Links)
}

func (x *AttrMask) afterLoad() {}
func (x *AttrMask) load(m state.Map) {
	m.Load("Type", &x.Type)
	m.Load("DeviceID", &x.DeviceID)
	m.Load("InodeID", &x.InodeID)
	m.Load("BlockSize", &x.BlockSize)
	m.Load("Size", &x.Size)
	m.Load("Usage", &x.Usage)
	m.Load("Perms", &x.Perms)
	m.Load("UID", &x.UID)
	m.Load("GID", &x.GID)
	m.Load("AccessTime", &x.AccessTime)
	m.Load("ModificationTime", &x.ModificationTime)
	m.Load("StatusChangeTime", &x.StatusChangeTime)
	m.Load("Links", &x.Links)
}

func (x *PermMask) beforeSave() {}
func (x *PermMask) save(m state.Map) {
	x.beforeSave()
	m.Save("Read", &x.Read)
	m.Save("Write", &x.Write)
	m.Save("Execute", &x.Execute)
}

func (x *PermMask) afterLoad() {}
func (x *PermMask) load(m state.Map) {
	m.Load("Read", &x.Read)
	m.Load("Write", &x.Write)
	m.Load("Execute", &x.Execute)
}

func (x *FilePermissions) beforeSave() {}
func (x *FilePermissions) save(m state.Map) {
	x.beforeSave()
	m.Save("User", &x.User)
	m.Save("Group", &x.Group)
	m.Save("Other", &x.Other)
	m.Save("Sticky", &x.Sticky)
	m.Save("SetUID", &x.SetUID)
	m.Save("SetGID", &x.SetGID)
}

func (x *FilePermissions) afterLoad() {}
func (x *FilePermissions) load(m state.Map) {
	m.Load("User", &x.User)
	m.Load("Group", &x.Group)
	m.Load("Other", &x.Other)
	m.Load("Sticky", &x.Sticky)
	m.Load("SetUID", &x.SetUID)
	m.Load("SetGID", &x.SetGID)
}

func (x *FileOwner) beforeSave() {}
func (x *FileOwner) save(m state.Map) {
	x.beforeSave()
	m.Save("UID", &x.UID)
	m.Save("GID", &x.GID)
}

func (x *FileOwner) afterLoad() {}
func (x *FileOwner) load(m state.Map) {
	m.Load("UID", &x.UID)
	m.Load("GID", &x.GID)
}

func (x *DentAttr) beforeSave() {}
func (x *DentAttr) save(m state.Map) {
	x.beforeSave()
	m.Save("Type", &x.Type)
	m.Save("InodeID", &x.InodeID)
}

func (x *DentAttr) afterLoad() {}
func (x *DentAttr) load(m state.Map) {
	m.Load("Type", &x.Type)
	m.Load("InodeID", &x.InodeID)
}

func (x *SortedDentryMap) beforeSave() {}
func (x *SortedDentryMap) save(m state.Map) {
	x.beforeSave()
	m.Save("names", &x.names)
	m.Save("entries", &x.entries)
}

func (x *SortedDentryMap) afterLoad() {}
func (x *SortedDentryMap) load(m state.Map) {
	m.Load("names", &x.names)
	m.Load("entries", &x.entries)
}

func (x *Dirent) save(m state.Map) {
	x.beforeSave()
	var children map[string]*Dirent = x.saveChildren()
	m.SaveValue("children", children)
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("userVisible", &x.userVisible)
	m.Save("Inode", &x.Inode)
	m.Save("name", &x.name)
	m.Save("parent", &x.parent)
	m.Save("deleted", &x.deleted)
	m.Save("frozen", &x.frozen)
	m.Save("mounted", &x.mounted)
}

func (x *Dirent) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("userVisible", &x.userVisible)
	m.Load("Inode", &x.Inode)
	m.Load("name", &x.name)
	m.Load("parent", &x.parent)
	m.Load("deleted", &x.deleted)
	m.Load("frozen", &x.frozen)
	m.Load("mounted", &x.mounted)
	m.LoadValue("children", new(map[string]*Dirent), func(y interface{}) { x.loadChildren(y.(map[string]*Dirent)) })
	m.AfterLoad(x.afterLoad)
}

func (x *DirentCache) beforeSave() {}
func (x *DirentCache) save(m state.Map) {
	x.beforeSave()
	if !state.IsZeroValue(x.currentSize) { m.Failf("currentSize is %v, expected zero", x.currentSize) }
	if !state.IsZeroValue(x.list) { m.Failf("list is %v, expected zero", x.list) }
	m.Save("maxSize", &x.maxSize)
	m.Save("limit", &x.limit)
}

func (x *DirentCache) afterLoad() {}
func (x *DirentCache) load(m state.Map) {
	m.Load("maxSize", &x.maxSize)
	m.Load("limit", &x.limit)
}

func (x *DirentCacheLimiter) beforeSave() {}
func (x *DirentCacheLimiter) save(m state.Map) {
	x.beforeSave()
	if !state.IsZeroValue(x.count) { m.Failf("count is %v, expected zero", x.count) }
	m.Save("max", &x.max)
}

func (x *DirentCacheLimiter) afterLoad() {}
func (x *DirentCacheLimiter) load(m state.Map) {
	m.Load("max", &x.max)
}

func (x *direntList) beforeSave() {}
func (x *direntList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *direntList) afterLoad() {}
func (x *direntList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *direntEntry) beforeSave() {}
func (x *direntEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *direntEntry) afterLoad() {}
func (x *direntEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *eventList) beforeSave() {}
func (x *eventList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *eventList) afterLoad() {}
func (x *eventList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *eventEntry) beforeSave() {}
func (x *eventEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *eventEntry) afterLoad() {}
func (x *eventEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *File) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("UniqueID", &x.UniqueID)
	m.Save("Dirent", &x.Dirent)
	m.Save("flags", &x.flags)
	m.Save("async", &x.async)
	m.Save("FileOperations", &x.FileOperations)
	m.Save("offset", &x.offset)
}

func (x *File) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("UniqueID", &x.UniqueID)
	m.Load("Dirent", &x.Dirent)
	m.Load("flags", &x.flags)
	m.Load("async", &x.async)
	m.LoadWait("FileOperations", &x.FileOperations)
	m.Load("offset", &x.offset)
	m.AfterLoad(x.afterLoad)
}

func (x *overlayFileOperations) beforeSave() {}
func (x *overlayFileOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("upper", &x.upper)
	m.Save("lower", &x.lower)
	m.Save("dirCursor", &x.dirCursor)
	m.Save("dirCache", &x.dirCache)
}

func (x *overlayFileOperations) afterLoad() {}
func (x *overlayFileOperations) load(m state.Map) {
	m.Load("upper", &x.upper)
	m.Load("lower", &x.lower)
	m.Load("dirCursor", &x.dirCursor)
	m.Load("dirCache", &x.dirCache)
}

func (x *overlayMappingIdentity) beforeSave() {}
func (x *overlayMappingIdentity) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("id", &x.id)
	m.Save("overlayFile", &x.overlayFile)
}

func (x *overlayMappingIdentity) afterLoad() {}
func (x *overlayMappingIdentity) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("id", &x.id)
	m.Load("overlayFile", &x.overlayFile)
}

func (x *MountSourceFlags) beforeSave() {}
func (x *MountSourceFlags) save(m state.Map) {
	x.beforeSave()
	m.Save("ReadOnly", &x.ReadOnly)
	m.Save("NoAtime", &x.NoAtime)
	m.Save("ForcePageCache", &x.ForcePageCache)
	m.Save("NoExec", &x.NoExec)
}

func (x *MountSourceFlags) afterLoad() {}
func (x *MountSourceFlags) load(m state.Map) {
	m.Load("ReadOnly", &x.ReadOnly)
	m.Load("NoAtime", &x.NoAtime)
	m.Load("ForcePageCache", &x.ForcePageCache)
	m.Load("NoExec", &x.NoExec)
}

func (x *FileFlags) beforeSave() {}
func (x *FileFlags) save(m state.Map) {
	x.beforeSave()
	m.Save("Direct", &x.Direct)
	m.Save("NonBlocking", &x.NonBlocking)
	m.Save("Sync", &x.Sync)
	m.Save("Append", &x.Append)
	m.Save("Read", &x.Read)
	m.Save("Write", &x.Write)
	m.Save("Pread", &x.Pread)
	m.Save("Pwrite", &x.Pwrite)
	m.Save("Directory", &x.Directory)
	m.Save("Async", &x.Async)
	m.Save("LargeFile", &x.LargeFile)
	m.Save("NonSeekable", &x.NonSeekable)
}

func (x *FileFlags) afterLoad() {}
func (x *FileFlags) load(m state.Map) {
	m.Load("Direct", &x.Direct)
	m.Load("NonBlocking", &x.NonBlocking)
	m.Load("Sync", &x.Sync)
	m.Load("Append", &x.Append)
	m.Load("Read", &x.Read)
	m.Load("Write", &x.Write)
	m.Load("Pread", &x.Pread)
	m.Load("Pwrite", &x.Pwrite)
	m.Load("Directory", &x.Directory)
	m.Load("Async", &x.Async)
	m.Load("LargeFile", &x.LargeFile)
	m.Load("NonSeekable", &x.NonSeekable)
}

func (x *Inode) beforeSave() {}
func (x *Inode) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("InodeOperations", &x.InodeOperations)
	m.Save("StableAttr", &x.StableAttr)
	m.Save("LockCtx", &x.LockCtx)
	m.Save("Watches", &x.Watches)
	m.Save("MountSource", &x.MountSource)
	m.Save("overlay", &x.overlay)
}

func (x *Inode) afterLoad() {}
func (x *Inode) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("InodeOperations", &x.InodeOperations)
	m.Load("StableAttr", &x.StableAttr)
	m.Load("LockCtx", &x.LockCtx)
	m.Load("Watches", &x.Watches)
	m.Load("MountSource", &x.MountSource)
	m.Load("overlay", &x.overlay)
}

func (x *LockCtx) beforeSave() {}
func (x *LockCtx) save(m state.Map) {
	x.beforeSave()
	m.Save("Posix", &x.Posix)
	m.Save("BSD", &x.BSD)
}

func (x *LockCtx) afterLoad() {}
func (x *LockCtx) load(m state.Map) {
	m.Load("Posix", &x.Posix)
	m.Load("BSD", &x.BSD)
}

func (x *Watches) beforeSave() {}
func (x *Watches) save(m state.Map) {
	x.beforeSave()
	m.Save("ws", &x.ws)
	m.Save("unlinked", &x.unlinked)
}

func (x *Watches) afterLoad() {}
func (x *Watches) load(m state.Map) {
	m.Load("ws", &x.ws)
	m.Load("unlinked", &x.unlinked)
}

func (x *Inotify) beforeSave() {}
func (x *Inotify) save(m state.Map) {
	x.beforeSave()
	m.Save("id", &x.id)
	m.Save("events", &x.events)
	m.Save("scratch", &x.scratch)
	m.Save("nextWatch", &x.nextWatch)
	m.Save("watches", &x.watches)
}

func (x *Inotify) afterLoad() {}
func (x *Inotify) load(m state.Map) {
	m.Load("id", &x.id)
	m.Load("events", &x.events)
	m.Load("scratch", &x.scratch)
	m.Load("nextWatch", &x.nextWatch)
	m.Load("watches", &x.watches)
}

func (x *Event) beforeSave() {}
func (x *Event) save(m state.Map) {
	x.beforeSave()
	m.Save("eventEntry", &x.eventEntry)
	m.Save("wd", &x.wd)
	m.Save("mask", &x.mask)
	m.Save("cookie", &x.cookie)
	m.Save("len", &x.len)
	m.Save("name", &x.name)
}

func (x *Event) afterLoad() {}
func (x *Event) load(m state.Map) {
	m.Load("eventEntry", &x.eventEntry)
	m.Load("wd", &x.wd)
	m.Load("mask", &x.mask)
	m.Load("cookie", &x.cookie)
	m.Load("len", &x.len)
	m.Load("name", &x.name)
}

func (x *Watch) beforeSave() {}
func (x *Watch) save(m state.Map) {
	x.beforeSave()
	m.Save("owner", &x.owner)
	m.Save("wd", &x.wd)
	m.Save("target", &x.target)
	m.Save("unpinned", &x.unpinned)
	m.Save("mask", &x.mask)
	m.Save("pins", &x.pins)
}

func (x *Watch) afterLoad() {}
func (x *Watch) load(m state.Map) {
	m.Load("owner", &x.owner)
	m.Load("wd", &x.wd)
	m.Load("target", &x.target)
	m.Load("unpinned", &x.unpinned)
	m.Load("mask", &x.mask)
	m.Load("pins", &x.pins)
}

func (x *MountSource) beforeSave() {}
func (x *MountSource) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("MountSourceOperations", &x.MountSourceOperations)
	m.Save("FilesystemType", &x.FilesystemType)
	m.Save("Flags", &x.Flags)
	m.Save("fscache", &x.fscache)
	m.Save("direntRefs", &x.direntRefs)
}

func (x *MountSource) afterLoad() {}
func (x *MountSource) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("MountSourceOperations", &x.MountSourceOperations)
	m.Load("FilesystemType", &x.FilesystemType)
	m.Load("Flags", &x.Flags)
	m.Load("fscache", &x.fscache)
	m.Load("direntRefs", &x.direntRefs)
}

func (x *SimpleMountSourceOperations) beforeSave() {}
func (x *SimpleMountSourceOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("keep", &x.keep)
	m.Save("revalidate", &x.revalidate)
}

func (x *SimpleMountSourceOperations) afterLoad() {}
func (x *SimpleMountSourceOperations) load(m state.Map) {
	m.Load("keep", &x.keep)
	m.Load("revalidate", &x.revalidate)
}

func (x *overlayMountSourceOperations) beforeSave() {}
func (x *overlayMountSourceOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("upper", &x.upper)
	m.Save("lower", &x.lower)
}

func (x *overlayMountSourceOperations) afterLoad() {}
func (x *overlayMountSourceOperations) load(m state.Map) {
	m.Load("upper", &x.upper)
	m.Load("lower", &x.lower)
}

func (x *overlayFilesystem) beforeSave() {}
func (x *overlayFilesystem) save(m state.Map) {
	x.beforeSave()
}

func (x *overlayFilesystem) afterLoad() {}
func (x *overlayFilesystem) load(m state.Map) {
}

func (x *Mount) beforeSave() {}
func (x *Mount) save(m state.Map) {
	x.beforeSave()
	m.Save("ID", &x.ID)
	m.Save("ParentID", &x.ParentID)
	m.Save("root", &x.root)
	m.Save("previous", &x.previous)
}

func (x *Mount) afterLoad() {}
func (x *Mount) load(m state.Map) {
	m.Load("ID", &x.ID)
	m.Load("ParentID", &x.ParentID)
	m.Load("root", &x.root)
	m.Load("previous", &x.previous)
}

func (x *MountNamespace) beforeSave() {}
func (x *MountNamespace) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("userns", &x.userns)
	m.Save("root", &x.root)
	m.Save("mounts", &x.mounts)
	m.Save("mountID", &x.mountID)
}

func (x *MountNamespace) afterLoad() {}
func (x *MountNamespace) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("userns", &x.userns)
	m.Load("root", &x.root)
	m.Load("mounts", &x.mounts)
	m.Load("mountID", &x.mountID)
}

func (x *overlayEntry) beforeSave() {}
func (x *overlayEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("lowerExists", &x.lowerExists)
	m.Save("lower", &x.lower)
	m.Save("mappings", &x.mappings)
	m.Save("upper", &x.upper)
}

func (x *overlayEntry) afterLoad() {}
func (x *overlayEntry) load(m state.Map) {
	m.Load("lowerExists", &x.lowerExists)
	m.Load("lower", &x.lower)
	m.Load("mappings", &x.mappings)
	m.Load("upper", &x.upper)
}

func init() {
	state.Register("fs.StableAttr", (*StableAttr)(nil), state.Fns{Save: (*StableAttr).save, Load: (*StableAttr).load})
	state.Register("fs.UnstableAttr", (*UnstableAttr)(nil), state.Fns{Save: (*UnstableAttr).save, Load: (*UnstableAttr).load})
	state.Register("fs.AttrMask", (*AttrMask)(nil), state.Fns{Save: (*AttrMask).save, Load: (*AttrMask).load})
	state.Register("fs.PermMask", (*PermMask)(nil), state.Fns{Save: (*PermMask).save, Load: (*PermMask).load})
	state.Register("fs.FilePermissions", (*FilePermissions)(nil), state.Fns{Save: (*FilePermissions).save, Load: (*FilePermissions).load})
	state.Register("fs.FileOwner", (*FileOwner)(nil), state.Fns{Save: (*FileOwner).save, Load: (*FileOwner).load})
	state.Register("fs.DentAttr", (*DentAttr)(nil), state.Fns{Save: (*DentAttr).save, Load: (*DentAttr).load})
	state.Register("fs.SortedDentryMap", (*SortedDentryMap)(nil), state.Fns{Save: (*SortedDentryMap).save, Load: (*SortedDentryMap).load})
	state.Register("fs.Dirent", (*Dirent)(nil), state.Fns{Save: (*Dirent).save, Load: (*Dirent).load})
	state.Register("fs.DirentCache", (*DirentCache)(nil), state.Fns{Save: (*DirentCache).save, Load: (*DirentCache).load})
	state.Register("fs.DirentCacheLimiter", (*DirentCacheLimiter)(nil), state.Fns{Save: (*DirentCacheLimiter).save, Load: (*DirentCacheLimiter).load})
	state.Register("fs.direntList", (*direntList)(nil), state.Fns{Save: (*direntList).save, Load: (*direntList).load})
	state.Register("fs.direntEntry", (*direntEntry)(nil), state.Fns{Save: (*direntEntry).save, Load: (*direntEntry).load})
	state.Register("fs.eventList", (*eventList)(nil), state.Fns{Save: (*eventList).save, Load: (*eventList).load})
	state.Register("fs.eventEntry", (*eventEntry)(nil), state.Fns{Save: (*eventEntry).save, Load: (*eventEntry).load})
	state.Register("fs.File", (*File)(nil), state.Fns{Save: (*File).save, Load: (*File).load})
	state.Register("fs.overlayFileOperations", (*overlayFileOperations)(nil), state.Fns{Save: (*overlayFileOperations).save, Load: (*overlayFileOperations).load})
	state.Register("fs.overlayMappingIdentity", (*overlayMappingIdentity)(nil), state.Fns{Save: (*overlayMappingIdentity).save, Load: (*overlayMappingIdentity).load})
	state.Register("fs.MountSourceFlags", (*MountSourceFlags)(nil), state.Fns{Save: (*MountSourceFlags).save, Load: (*MountSourceFlags).load})
	state.Register("fs.FileFlags", (*FileFlags)(nil), state.Fns{Save: (*FileFlags).save, Load: (*FileFlags).load})
	state.Register("fs.Inode", (*Inode)(nil), state.Fns{Save: (*Inode).save, Load: (*Inode).load})
	state.Register("fs.LockCtx", (*LockCtx)(nil), state.Fns{Save: (*LockCtx).save, Load: (*LockCtx).load})
	state.Register("fs.Watches", (*Watches)(nil), state.Fns{Save: (*Watches).save, Load: (*Watches).load})
	state.Register("fs.Inotify", (*Inotify)(nil), state.Fns{Save: (*Inotify).save, Load: (*Inotify).load})
	state.Register("fs.Event", (*Event)(nil), state.Fns{Save: (*Event).save, Load: (*Event).load})
	state.Register("fs.Watch", (*Watch)(nil), state.Fns{Save: (*Watch).save, Load: (*Watch).load})
	state.Register("fs.MountSource", (*MountSource)(nil), state.Fns{Save: (*MountSource).save, Load: (*MountSource).load})
	state.Register("fs.SimpleMountSourceOperations", (*SimpleMountSourceOperations)(nil), state.Fns{Save: (*SimpleMountSourceOperations).save, Load: (*SimpleMountSourceOperations).load})
	state.Register("fs.overlayMountSourceOperations", (*overlayMountSourceOperations)(nil), state.Fns{Save: (*overlayMountSourceOperations).save, Load: (*overlayMountSourceOperations).load})
	state.Register("fs.overlayFilesystem", (*overlayFilesystem)(nil), state.Fns{Save: (*overlayFilesystem).save, Load: (*overlayFilesystem).load})
	state.Register("fs.Mount", (*Mount)(nil), state.Fns{Save: (*Mount).save, Load: (*Mount).load})
	state.Register("fs.MountNamespace", (*MountNamespace)(nil), state.Fns{Save: (*MountNamespace).save, Load: (*MountNamespace).load})
	state.Register("fs.overlayEntry", (*overlayEntry)(nil), state.Fns{Save: (*overlayEntry).save, Load: (*overlayEntry).load})
}
