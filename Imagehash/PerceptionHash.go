package Imagehash

import (
	"Dhash/Imagehash/etcs"
	"Dhash/Imagehash/transforms"
	"errors"
	"github.com/nfnt/resize"
	"image"
)

// PerceptionHash function returns a hash computation of phash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
// 实质上是对频率的比较
func PerceptionHash(img image.Image) (*ImageHash, error,int) {
	if img == nil {
		return nil, errors.New("Image object can not be nil"),2
	}

	phash := NewImageHash(0, PHash)
	// 将照片压缩为64*64像素的缩略图，
	resized := resize.Resize(64, 64, img, resize.Bilinear)
	// 将缩略图转化为灰度矩阵
	pixels := transforms.Rgb2Gray(resized)
	// 利用离散余弦变换（DCT）降低频率，缩小矩阵
	dct := transforms.DCT2D(pixels, 64, 64)
	// 将二维数组转化为一维数组
	flattens := transforms.FlattenPixels(dct, 8, 8)
	// 获取像素的平均值
	median := etcs.MedianOfPixels(flattens)

	for idx, p := range flattens {
		if p > median {
			phash.leftShiftSet(len(flattens) - idx - 1)
		}
	}
	return phash, nil,2
}
