package Imagehash

import (
	"Dhash/Imagehash/transforms"
	"errors"
	"github.com/nfnt/resize"
	"image"
)

// DifferenceHash function returns a hash computation of difference hash.
// Implementation follows
// http://www.hackerfactor.com/blog/?/archives/529-Kind-of-Like-That.html
// 实质上是基于渐变的感知哈希算法，是对比相邻像素的差异
func DifferenceHash(img image.Image) (*ImageHash, error,int) {
	if img == nil {
		return nil, errors.New("Image object can not be nil"),3
	}

	dhash := NewImageHash(0, DHash)
	// 对图像进行宽度和高度的调整，在此缩放为9*8像素
	// 每个像素均保留一个RGB值（红绿蓝色彩模式），故像素过高，信息量较大，处理较为麻烦
	// resize.Bilinear参数将会对所有可以影响输出像素的输入像素进行高质量的重采样滤波
	resized := resize.Resize(9, 8, img, resize.Bilinear)
	// 将RGB值转换为灰度矩阵，
	pixels := transforms.Rgb2Gray(resized)
	idx := 0
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[i])-1; j++ {
			if pixels[i][j] < pixels[i][j+1] {
				//计算灰度矩阵相邻像素的差异值，通过进位的方式得到只含有0、1的64位的哈希值
				dhash.leftShiftSet(64 - idx - 1)
			}
			idx++
		}
	}

	return dhash, nil,3
}