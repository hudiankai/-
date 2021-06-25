package Imagehash

import (
	"Dhash/Imagehash/etcs"
	"Dhash/Imagehash/transforms"
	"errors"
	"github.com/nfnt/resize"
	"image"
)

// AverageHash fuction returns a hash computation of average hash.
// Implementation follows
// http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html
// 实际上是对图片颜色的比较
func AverageHash(img image.Image) (*ImageHash, error,int) {
	if img == nil {
		return nil, errors.New("Image object can not be nil"),1
	}

	// Create 64bits hash.
	ahash := NewImageHash(0, AHash)
	// 将照片缩减成8*8像素的缩略图，resize.Bilinear参数将会对所有可以影响输出像素的输入像素进行高质量的重采样滤波
	resized := resize.Resize(8, 8, img, resize.Bilinear)
	// 将缩略图转化为灰度矩阵
	pixels := transforms.Rgb2Gray(resized)
	// 将灰度矩阵（二维数组）转化为一维数组
	flattens := transforms.FlattenPixels(pixels, 8, 8)
	// 获取像素的平均值
	avg := etcs.MeanOfPixels(flattens)

	for idx, p := range flattens {
		if p > avg {
			// 若大约平均像素，则将该位设定为1，否则为0
			ahash.leftShiftSet(len(flattens) - idx - 1)
		}
	}

	return ahash, nil,1
}