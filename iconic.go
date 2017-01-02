package main

import (
    "image"
    "image/png"
    "log"
    "os"

    "github.com/nfnt/resize"
)

// IconOut represents a output parameters of the final png set
type IconOut struct {
    filename  string
    imageData image.Image
}

func main() {
    file, err := os.Open("default.png") // Open default 1024 by 1024 icon
    if err != nil {
        log.Fatal(err)
    }

    img, err := png.Decode(file) // Decode png to file
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    img512 := resize.Resize(512, 0, img, resize.Lanczos3)
    img256 := resize.Resize(256, 0, img, resize.Lanczos3)
    img128 := resize.Resize(128, 0, img, resize.Lanczos3)
    img64 := resize.Resize(64, 0, img, resize.Lanczos3)
    img32 := resize.Resize(32, 0, img, resize.Lanczos3)
    img16 := resize.Resize(16, 0, img, resize.Lanczos3)

    out16 := IconOut{"icon_16x16.png", img16}
    out16x2 := IconOut{"icon_16x16@2x.png", img32}
    out32 := IconOut{"icon_32x32.png", img32}
    out32x2 := IconOut{"icon_32x32@2x.png", img64}
    out64 := IconOut{"icon_64x64.png", img64}
    out64x2 := IconOut{"icon_64x64@2x.png", img128}
    out128 := IconOut{"icon_128x128.png", img128}
    out128x2 := IconOut{"icon_128x128@2x.png", img256}
    out256 := IconOut{"icon_256x256.png", img256}
    out256x2 := IconOut{"icon_256x256@2x.png", img512}
    out512 := IconOut{"icon_512x512.png", img512}
    out512x2 := IconOut{"icon_512x512@2x.png", img}

    filenames := []IconOut{out16, out16x2, out32, out32x2, out64, out64x2, out128, out128x2, out256, out256x2, out512, out512x2}

    for _, i := range filenames {
        out, err := os.Create(i.filename)
        if err != nil {
            log.Fatal(err)
        }
        png.Encode(out, i.imageData)
        out.Close()
    }
}
