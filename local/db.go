package local

import (
    "fmt"
    "github.com/boltdb/bolt"
    "time"
)

type Local struct {
    db *bolt.DB
    bucket map[string]*bolt.Bucket
}
func (l Local) New (name string)  {
    l.db, _ = bolt.Open(fmt.Sprintf(`eros_%s.db`, name), 0600, &bolt.Options{Timeout: 1 * time.Second})
}


func (l Local) Bucket (name string) *bolt.Bucket {
    _ = l.db.Update(func(tx *bolt.Tx) error {
        l.bucket[name], _ = tx.CreateBucket([]byte(name))
        return nil
    })
   return l.bucket[name]
}

func (l Local) Put(bucket, key, val string) error {

    return l.Bucket(bucket).Put([]byte(key), []byte(key))
}


func (l Local) Get(bucket, key string) string {

    return string(l.Bucket(bucket).Get([]byte(key)))
}
