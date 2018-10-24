# 持久化说明
## repo接口
```
type Repo interface {
 // Config returns the ipfs configuration file from the repo. Changes made
 // to the returned config are not automatically persisted.
 Config() (*config.Config, error)

 // BackupConfig creates a backup of the current configuration file using
 // the given prefix for naming.
 BackupConfig(prefix string) (string, error)

 // SetConfig persists the given configuration struct to storage.
 SetConfig(*config.Config) error

 // SetConfigKey sets the given key-value pair within the config and persists it to storage.
 SetConfigKey(key string, value interface{}) error

 // GetConfigKey reads the value for the given key from the configuration in storage.
 GetConfigKey(key string) (interface{}, error)

 // Datastore returns a reference to the configured data storage backend.
 Datastore() Datastore

 // GetStorageUsage returns the number of bytes stored.
 GetStorageUsage() (uint64, error)

 // Keystore returns a reference to the key management interface.
 Keystore() keystore.Keystore

 // FileManager returns a reference to the filestore file manager.
 FileManager() *filestore.FileManager

 // SetAPIAddr sets the API address in the repo.
 SetAPIAddr(addr ma.Multiaddr) error

 // SwarmKey returns the configured shared symmetric key for the private networks feature.
 SwarmKey() ([]byte, error)

 io.Closer
}
```
```
type Filestore struct {
 fm *FileManager
 bs blockstore.Blockstore
}
```
```
type Blockstore interface {
 DeleteBlock(*cid.Cid) error
 Has(*cid.Cid) (bool, error)
 Get(*cid.Cid) (blocks.Block, error)

 // GetSize returns the CIDs mapped BlockSize
 GetSize(*cid.Cid) (int, error)

 // Put puts a given block to the underlying datastore
 Put(blocks.Block) error

 // PutMany puts a slice of blocks at the same time using batching
 // capabilities of the underlying datastore whenever possible.
 PutMany([]blocks.Block) error

 // AllKeysChan returns a channel from which
 // the CIDs in the Blockstore can be read. It should respect
 // the given context, closing the channel if it becomes Done.
 AllKeysChan(ctx context.Context) (<-chan *cid.Cid, error)

 // HashOnRead specifies if every read block should be
 // rehashed to make sure it matches its CID.
 HashOnRead(enabled bool)
}
```
## blocks目录
- 具体的IPLD内容

## leveldb
- [x] 内容dht表记录
- [x] ipns记录
- [ ] 节点key，/F5UXA3TTF4JCBMC4HJRBNOMOVBZRBQSWFONFZ3KEXTW2UIMVANT5L4JCBWLWIAIN
- [ ] pki，key:/F5YGWLYSECYFYOTCC24Y5KDTCDBFMK42LTWUJPHNVIQZKA3H2XYSEDMXMQAQ2
- [ ] key:/local/filesroot
- [ ] key:/local/pins

## 为便于分析写的小工具
- [x] [leveldb_read](/datastores/leveldb/leveldb_read.go)
- [x] [cid-dskey](/datastores/cid-dskey/cid-dskey.go)

## [示例](/datastores/example.md)

