package badger

import (
	"bytes"
	"strconv"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/xyctruth/profiler/pkg/storage"
)

var (
	Sequence          = []byte{0x80}
	PrefixProfiles    = []byte{0x81}
	PrefixProfileMeta = []byte{0x82}
	PrefixSampleType  = []byte{0x83}
	PrefixTarget      = []byte{0x84}
)

func buildProfileKey(id string) []byte {
	buf := bytes.NewBuffer(PrefixProfiles)
	buf.Write([]byte(id))
	return buf.Bytes()
}

func buildBaseProfileMetaKey(sampleType string, target string) []byte {
	buf := bytes.NewBuffer(PrefixProfileMeta)
	buf.Write([]byte(sampleType))
	buf.Write([]byte(target))
	return buf.Bytes()
}

func buildProfileMetaKey(sampleType string, target string, datetime time.Time) []byte {
	buf := bytes.NewBuffer(PrefixProfileMeta)
	buf.Write([]byte(sampleType))
	buf.Write([]byte(target))
	buf.Write(storage.BuildKey(datetime))
	return buf.Bytes()
}

func buildSampleTypeKey(sampleType string) []byte {
	buf := bytes.NewBuffer(PrefixSampleType)
	buf.Write([]byte(sampleType))
	return buf.Bytes()
}

func buildTargetKey(target string) []byte {
	buf := bytes.NewBuffer(PrefixTarget)
	buf.Write([]byte(target))
	return buf.Bytes()
}

func newProfileEntry(id uint64, val []byte, ttl time.Duration) *badger.Entry {
	entry := badger.NewEntry(buildProfileKey(strconv.FormatUint(id, 10)), val)
	if ttl > 0 {
		entry = entry.WithTTL(ttl)
	}
	return entry
}

func newProfileMetaEntry(meta *storage.ProfileMeta, ttl time.Duration) (*badger.Entry, error) {
	metaBytes, err := meta.Encode()
	if err != nil {
		return nil, err
	}
	entry := badger.NewEntry(buildProfileMetaKey(meta.SampleType, meta.TargetName, time.Now()), metaBytes)
	if ttl > 0 {
		entry = entry.WithTTL(ttl)
	}
	return entry, nil
}

func newSampleTypeEntry(sampleType string, profileType string, ttl time.Duration) *badger.Entry {
	entry := badger.NewEntry(buildSampleTypeKey(sampleType), []byte(profileType))
	if ttl > 0 {
		entry = entry.WithTTL(ttl)
	}
	return entry
}

func newTargetKeyEntry(target string, ttl time.Duration) *badger.Entry {
	entry := badger.NewEntry(buildTargetKey(target), []byte{})
	if ttl > 0 {
		entry = entry.WithTTL(ttl)
	}
	return entry
}

func deleteSampleTypeKey(sampleType []byte) string {
	return string(sampleType[1:])
}

func deleteTargetKey(target []byte) string {
	return string(target[1:])
}
