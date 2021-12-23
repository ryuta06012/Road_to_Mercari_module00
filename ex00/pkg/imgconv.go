/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   imgconv.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/23 07:53:30 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/23 12:34:02 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func removeSrc(src string) error {
	err := os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}

func Convert(src, dst string, rmsrc bool) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	switch filepath.Ext(dst) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case ".gif":
		err = gif.Encode(df, img, nil)
	}
	if err != nil {
		return err
	}
	if rmsrc {
		if err = removeSrc(src); err != nil {
			return err
		}
	}
	return nil
}
