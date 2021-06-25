package Imagehash

import (
	"errors"
	"math/bits"
)

// Kind describes the kinds of hash.
type Kind int

// ImageHash is a struct of hash computation.
type ImageHash struct {
	hash uint64
	kind Kind
}

// ExtImageHash is a struct of big hash computation.
type ExtImageHash struct {
	hash []uint64
	kind Kind
	bits int
}

const (
	// Unknown is a enum value of the unknown hash.
	Unknown Kind = iota
	// AHash is a enum value of the average hash.
	AHash
	//PHash is a enum value of the perceptual hash.
	PHash
	// DHash is a enum value of the difference hash.
	DHash
	// WHash is a enum value of the wavelet hash.
	WHash
)

// 创建一个新的图片Hash
func NewImageHash(hash uint64, kind Kind) *ImageHash {
	return &ImageHash{hash: hash, kind: kind}
}

func (h *ImageHash) leftShiftSet(idx int) {
	h.hash |= 1 << uint(idx)
	//与下问相同
	//var hhash uint64
	//hhash := 1 << uint(idx)
	//h.hash = h.hash | hhash
}

// GetKind method returns a kind of image hash.
func (h *ImageHash) GetKind() Kind {
	return h.kind
}

// GetHash method returns a 64bits hash value.
func (h *ImageHash) GetHash() uint64 {
	return h.hash
}

func popcnt(x uint64) int { return bits.OnesCount64(x) }

// 计算两个hash的海明距离
func (h *ImageHash) Distance(other *ImageHash) (int, error) {
	if h.GetKind() != other.GetKind() {
		return -1, errors.New("Image hashes's kind should be identical")
	}

	lhash := h.GetHash()
	rhash := other.GetHash()

	hamming := lhash ^ rhash
	return popcnt(hamming), nil
}
